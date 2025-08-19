package goscord

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Murilinho145SG/gouter/log"
)

func (b *Bot) SendRawMessage(channel string, payload *CreateMessagePayload) (io.ReadCloser, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return b.sendPost("/channels/"+channel+"/messages", data)
}

func (b *Bot) ToFormat(format string, msg string) string {
	return "```" + format + "\n" + msg + "\n" + "```"
}

type AllowedMentions struct {
	Parse       []string `json:"parse,omitempty"`        // "roles", "users", "everyone"
	Roles       []string `json:"roles,omitempty"`        // IDs de roles permitidos
	Users       []string `json:"users,omitempty"`        // IDs de usuários permitidos
	RepliedUser *bool    `json:"replied_user,omitempty"` // se true, vai pingar o autor da msg reply
}

type CreateMessagePayload struct {
	Content          string            `json:"content,omitempty"`           // mensagem de texto
	TTS              bool              `json:"tts,omitempty"`               // true = mensagem TTS
	Nonce            string            `json:"nonce,omitempty"`             // para evitar duplicação
	Embeds           []Embed           `json:"embeds,omitempty"`            // embeds ricos
	AllowedMentions  *AllowedMentions  `json:"allowed_mentions,omitempty"`  // quem pode ser pingado
	MessageReference *MessageReference `json:"message_reference,omitempty"` // reply
	Components       []Component       `json:"components,omitempty"`        // botões e menus (opcional)
	StickerIDs       []string          `json:"sticker_ids,omitempty"`       // até 3 stickers
	Flags            int               `json:"flags,omitempty"`             // SUPPRESS_EMBEDS, etc
	Poll             any               `json:"poll,omitempty"`              // objeto de enquete
}

func (b *Bot) SendMessage(channel, content string) error {
	payload := CreateMessagePayload{
		Content: content,
	}
	_, err := b.SendRawMessage(channel, &payload)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Reply(channel, msgId, content string) error {
	payload := CreateMessagePayload{
		Content: content,
		MessageReference: &MessageReference{
			MessageId: msgId,
		},
	}
	_, err := b.SendRawMessage(channel, &payload)
	if err != nil {
		return err
	}

	return nil
}

type MessageBuilder struct {
	bot     *Bot
	message CreateMessagePayload
	channel string
}

func (b *Bot) MessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		bot: b,
	}
}

func (mb *MessageBuilder) Content(content string) *MessageBuilder {
	mb.message.Content = content
	return mb
}

func (mb *MessageBuilder) Embed(embed Embed) *MessageBuilder {
	mb.message.Embeds = append(mb.message.Embeds, embed)
	return mb
}

func (mb *MessageBuilder) Component(components Component) *MessageBuilder {
	assertType(components)
	mb.message.Components = append(mb.message.Components, components)
	return mb
}

func assertType(components Component) {
	switch components.(type) {
	case *ActionRow:
		ar := components.(*ActionRow)
		ar.Type = ActionsRowType
		for i := range ar.Components {
			assertType(ar.Components[i])
		}
	case *Button:
		btn := components.(*Button)
		btn.Type = ButtonType
	case *StringSelect:
		ss := components.(*StringSelect)
		ss.Type = StringSelectType
	case *TextInput:
		ti := components.(*TextInput)
		ti.Type = TextInputType
	case *UserSelect:
		us := components.(*UserSelect)
		us.Type = UserSelectType
	case *RoleSelect:
		rs := components.(*RoleSelect)
		rs.Type = RoleSelectType
	case *MentionableSelect:
		ms := components.(*MentionableSelect)
		ms.Type = MentionableSelectType
	case *ChannelSelect:
		cs := components.(*ChannelSelect)
		cs.Type = ChannelSelectType
	case *Section:
		s := components.(*Section)
		s.Type = SectionType
		for i := range s.Accessory {
			assertType(s.Accessory[i])
		}
	case *TextDisplay:
		td := components.(*TextDisplay)
		td.Type = TextDisplayType
	case *Thumbnail:
		t := components.(*Thumbnail)
		t.Type = ThumbnailType
	case *MediaGallery:
		mg := components.(*MediaGallery)
		mg.Type = MediaGalleryType
	case *File:
		f := components.(*File)
		f.Type = FileType
	case *Separator:
		s := components.(*Separator)
		s.Type = SeparatorType
	case *Container:
		c := components.(*Container)
		c.Type = ContainerType
		for i := range c.Components {
			assertType(c.Components[i])
		}
	}
}

func (mb *MessageBuilder) Ephemeral(value bool) *MessageBuilder {
	mb.message.Flags = 64
	return mb
}

func (mb *MessageBuilder) SetChannel(channel string) *MessageBuilder {
	mb.channel = channel
	return mb
}

func (mb *MessageBuilder) Send() {
	log.Info(mb.channel, mb.message)
	data, err := json.MarshalIndent(mb.message, "", " ")
	if err != nil {
		log.Error(err)
		return
	}
	mb.bot.SendMessage(mb.channel, "```json\n"+string(data)+"\n```")

	rc, err := mb.bot.SendRawMessage(mb.channel, &mb.message)
	if err != nil {
		log.Error(err)
		return
	}
	b, err := io.ReadAll(rc)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(string(b))
}

type InteractionCallback struct {
	Id                       string `json:"id"`
	Type                     int    `json:"type"`
	ActivityInstanceId       string `json:"activity_instance_id,omitempty"`
	ResponseMessageId        string `json:"response_message_id,omitempty"`
	ResponseMessageLoading   bool   `json:"response_message_loading,omitempty"`
	ResponseMessageEphemeral bool   `json:"response_message_ephemeral,omitempty"`
}

func (b *Bot) RawInteractionResponse(interaction Interaction) (io.ReadCloser, error) {
	// ic := InteractionCallback{
	// 	Id:   interaction.Id,
	// 	Type: int(interaction.Type),
		
	// }
	return b.sendPost(fmt.Sprintf("/interactions/%s/%s/callback", interaction.Id, interaction.Token), nil)
}
