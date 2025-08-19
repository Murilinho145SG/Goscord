package goscord

import (
	"encoding/json"
)

type Message struct {
	Id                   string                      `json:"id"`
	ChannelId            string                      `json:"channel_id"`
	Author               *User                       `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            string                      `json:"timestamp"`
	EditedTimestamp      string                      `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []User                      `json:"mentions"`
	MentionRoles         []Role                      `json:"mention_roles"`
	MentionChannels      []ChannelMention            `json:"mention_channels,omitempty"`
	Attachments          []Attachments               `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                string                      `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookId            string                      `json:"webhook_id,omitempty"`
	Type                 int                         `json:"type"`
	Activity             *MessageActivity            `json:"activity,omitempty"`
	Application          *Application                `json:"application,omitempty"`
	ApplicationId        string                      `json:"application_id,omitempty"`
	Flags                int                         `json:"flags,omitempty"`
	MessageReference     *MessageReference           `json:"message_reference,omitempty"`
	MessageSnapshots     *MessageSnapshot            `json:"message_snapshots,omitempty"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	InteractionMetadata  *MessageInteractionMetadata `json:"interaction_metadata,omitempty"`
	Interaction          *MessageInteraction         `json:"interaction,omitempty"`
	Thread               *Channel                    `json:"thread,omitempty"`
	Components           []RawComponent              `json:"components,omitempty"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Position             int                         `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData       `json:"role_subscription_data,omitempty"`
	Resolved             *ResolvedData               `json:"resolved,omitempty"`
	Poll                 *Poll                       `json:"poll,omitempty"`
	Call                 *MessageCall                `json:"call,omitempty"`
	Bot                  *Bot
}

type MessageCall struct {
	Participants   []string `json:"participants"`
	EndedTimestamp string   `json:"ended_timestamp"`
}

type MessageCreate struct {
	*Message
}

func NewMessage(bot *Bot, data map[string]interface{}) *Message {
	var m Message
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	json.Unmarshal(b, &m)
	m.Bot = bot
	return &m
}

func (msg *Message) SendMessage(channel, content string) error {
	return msg.Bot.SendMessage(channel, content)
}

func (msg *Message) Reply(content string) error {
	return msg.Bot.Reply(msg.ChannelId, msg.Id, content)
}

func NewMessageCreate(bot *Bot, data map[string]interface{}) *MessageCreate {
	m := NewMessage(bot, data)
	return &MessageCreate{
		m,
	}
}
