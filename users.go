package goscord

import "fmt"

type User struct {
	Id                   string                `json:"id"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	GlobalName           string                `json:"global_name"`
	Avatar               string                `json:"avatar"`
	Bot                  bool                  `json:"bot,omitempty"`
	System               bool                  `json:"system,omitempty"`
	MFAEnabled           bool                  `json:"mfa_enabled,omitempty"`
	Banner               string                `json:"banner,omitempty"`
	AccentColor          int                   `json:"accent_color,omitempty"`
	Locale               string                `json:"locale,omitempty"`
	Verified             bool                  `json:"verified,omitempty"`
	Email                string                `json:"email,omitempty"`
	Flags                int                   `json:"flags,omitempty"`
	PremiumType          int                   `json:"premium_type,omitempty"`
	PublicFlags          int                   `json:"public_flags,omitempty"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles         *Collectible          `json:"collectible,omitempty"`
	PrimaryGuild         *UserPrimaryGuild     `json:"primary_guild,omitempty"`
}

func (user *User) GetAvatar() string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", user.Id, user.Avatar)
}
