package goscord

import (
	"encoding/json"

	"github.com/Murilinho145SG/gouter/log"
)

type InteractionType int
type InteractionCtxType int

const ()

type Interaction struct {
	Id                           string            `json:"id"`
	ApplicationId                string            `json:"application_id"`
	Type                         InteractionType   `json:"type"`
	Data                         *InteractionData  `json:"data,omitempty"`
	Guild                        *Guild            `json:"guild,omitempty"`
	GuildId                      string            `json:"guild_id,omitempty"`
	Channel                      *Channel          `json:"channel,omitempty"`
	ChannelId                    string            `json:"channel_id,omitempty"`
	Member                       *Member           `json:"member,omitempty"`
	User                         *User             `json:"user,omitempty"`
	Token                        string            `json:"token"`
	Version                      int               `json:"version"`
	Message                      *Message          `json:"message,omitempty"`
	AppPermissions               string            `json:"app_permissions"`
	Locale                       string            `json:"locale,omitempty"`
	Entitlements                 []Entitlement     `json:"entitlements"`
	AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
	Context                      int               `json:"context,omitempty"`
	AttachmentSizeLimit          int               `json:"attachment_size_limit"`
	Bot                          *Bot
}

type InteractionCreate struct {
	*Interaction
}

func NewInteraction(bot *Bot, data map[string]interface{}) *Interaction {
	var interaction Interaction
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return nil
	}

	if err := bot.SendMessage("1241131066560090185", string(b)); err != nil {
		log.Error(err)
	}

	if err := json.Unmarshal(b, &interaction); err != nil {
		log.Error(err)
		return nil
	}

	interaction.Bot = bot
	return &interaction
}

func NewInteractionCreate(bot *Bot, data map[string]interface{}) *InteractionCreate {
	i := NewInteraction(bot, data)

	return &InteractionCreate{
		i,
	}
}

type InteractionData struct {
	Id       int                     `json:"id"`
	Name     string                  `json:"name"`
	Type     int                     `json:"type"`
	Resolved *ResolvedData           `json:"resolved,omitempty"`
	Options  []InteractionDataOption `json:"options,omitempty"`
	GuildId  string                  `json:"guild_id,omitempty"`
	TargetId string                  `json:"target_id"`
}

type InteractionDataOption struct {
	Name    string                  `json:"name"`
	Type    int                     `json:"type"`
	Value   interface{}             `json:"value,omitempty"`
	Options []InteractionDataOption `json:"options,omitempty"`
	Focused bool                    `json:"focused,omitempty"`
}
