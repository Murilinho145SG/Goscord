package goscord

type Application struct {
	Id                                string                                   `json:"id"`
	Name                              string                                   `json:"name"`
	Icon                              string                                   `json:"icon"`
	Description                       string                                   `json:"description"`
	RPCOrigins                        []string                                 `json:"rpc_origins,omitempty"`
	BotPublic                         bool                                     `json:"bot_public"`
	BotRequireCodeGrant               bool                                     `json:"bot_require_code_grant"`
	Bot                               *User                                    `json:"bot,omitempty"`
	TermsOfServiceUrl                 string                                   `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyUrl                  string                                   `json:"privacy_policy_url,omitempty"`
	Owner                             *User                                    `json:"owner,omitempty"`
	VerifyKey                         string                                   `json:"verify_key"`
	Team                              *Team                                    `json:"team"`
	GuildId                           string                                   `json:"guild_id,omitempty"`
	Guild                             *Guild                                   `json:"guild,omitempty"`
	PrimarySkyId                      string                                   `json:"primary_sku_id,omitempty"`
	Slug                              string                                   `json:"slug,omitempty"`
	CoverImage                        string                                   `json:"cover_image,omitempty"`
	Flags                             int                                      `json:"flags,omitempty"`
	ApproximateGuildCount             int                                      `json:"approximate_guild_count,omitempty"`
	ApproximateUserInstallCount       int                                      `json:"approximate_user_install_count,omitempty"`
	ApproximateUserAuthorizationCount int                                      `json:"approximate_user_authorization_count,omitempty"`
	RedirectUris                      []string                                 `json:"redirect_uris,omitempty"`
	InteractionsEndpointUrl           string                                   `json:"interactions_endpoint_url,omitempty"`
	RoleConnectionsVerificationUrl    string                                   `json:"role_connections_verification_url,omitempty"`
	EventWebhooksUrl                  string                                   `json:"event_webhooks_url,omitempty"`
	EventWebhookStatus                int                                      `json:"event_webhooks_status,omitempty"`
	EventWebhooksTypes                []string                                 `json:"event_webhooks_types,omitempty"`
	Tags                              []string                                 `json:"tags,omitempty"`
	InstallParams                     *InstallParams                           `json:"install_params,omitempty"`
	IntegrationTypesConfig            *ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	CustomInstallUrl                  string                                   `json:"custom_install_url,omitempty"`
}

type Guild struct {
	Id                          string         `json:"id"`
	Name                        string         `json:"name"`
	Icon                        string         `json:"icon"`
	IconHash                    string         `json:"icon_hash,omitempty"`
	Splash                      string         `json:"splash"`
	DiscoverySplash             string         `json:"discovery_splash"`
	Owner                       bool           `json:"owner,omitempty"`
	OwnerId                     string         `json:"owner_id"`
	Permissions                 string         `json:"permissions,omitempty"`
	Region                      string         `json:"region"`
	AfkChannelId                string         `json:"afk_channel_id"`
	AfkTimeout                  int            `json:"afk_timeout"`
	WidgetEnabled               bool           `json:"widget_enabled,omitempty"`
	WidgetChannelId             string         `json:"widget_channel_id,omitempty"`
	VerificationLevel           int            `json:"verification_level"`
	DefaultMessageNotifications int            `json:"default_message_notifications"`
	ExplicitContentFilter       int            `json:"explicit_content_filter"`
	Roles                       []Role         `json:"roles"`
	Emojis                      []Emoji        `json:"emojis"`
	Features                    []string       `json:"features"`
	MFALevel                    int            `json:"mfa_level"`
	ApplicationId               string         `json:"application_id"`
	SystemChannelFlags          int            `json:"system_channel_flags"`
	RulesChannelId              string         `json:"rules_channel_id"`
	MaxPresences                int            `json:"max_presences,omitempty"`
	MaxMembers                  int            `json:"max_members,omitempty"`
	VanityUrlCode               string         `json:"vanity_url_code"`
	Description                 string         `json:"description"`
	Banner                      string         `json:"banner"`
	PremiumTier                 int            `json:"premium_tier"`
	PremiumSubscriptionCount    int            `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale"`
	PublicUpdatesChannelId      string         `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int            `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   int            `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      int            `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int            `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NSFWLevel                   int            `json:"nsfw_level"`
	Stickers                    []Sticker      `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelId       string         `json:"safety_alerts_channel_id"`
	IncidentsData               *IncidentsData `json:"incidents_data"`
}

type IncidentsData struct {
	InvitesDisableUntil string `json:"invites_disable_until"`
	DmsDisabledUntil    string `json:"dms_disabled_until"`
	DmSpamDetectedAt    string `json:"dm_spam_detected_at,omitempty"`
	RaidDetectedAt      string `json:"raid_detected_at,omitempty"`
}

type Sticker struct {
	Id          string `json:"id"`
	PackId      string `json:"pack_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Type        int    `json:"type"`
	FormatType  int    `json:"format_type"`
	Available   bool   `json:"available,omitempty"`
	GuildId     string `json:"guild_id,omitempty"`
	User        *User  `json:"user,omitempty"`
	SortValue   int    `json:"sort_value,omitempty"`
}

type WelcomeScreen struct {
	Description     string               `json:"description"`
	WelcomeChannels WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelId   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiId     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

type Team struct {
	Icon        string      `json:"icon"`
	Id          string      `json:"id"`
	Members     *TeamMember `json:"members"`
	Name        string      `json:"name"`
	OwnerUserId string      `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState int    `json:"membership_state"`
	TeamId          string `json:"team_id"`
	User            *User  `json:"user"`
	Role            string `json:"role"`
}

type ApplicationIntegrationTypeConfiguration struct {
	Oauth2InstallParams *InstallParams `json:"oauth2_install_params,omitempty"`
}

type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}

type MessageActivity struct {
	Type    int    `json:"type"`
	PartyId string `json:"party_id,omitempty"`
}

type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	ME           bool                 `json:"me"`
	MEBurst      bool                 `json:"me_burst"`
	Emoji        Emoji                `json:"emoji"`
	BurstColors  []string             `json:"burst_colors"`
}

type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type Emoji struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Roles         []Role `json:"roles,omitempty"`
	User          User   `json:"user,omitempty"`
	RequireColons bool   `json:"require_colons,omitempty"`
	Managed       bool   `json:"managed,omitempty"`
	Animated      bool   `json:"animated,omitempty"`
	Available     bool   `json:"available,omitempty"`
}

type Embed struct {
	Title       string
	Type        string
	Description string
	Url         string
	Timestamp   string
	Color       int
	Footer      *EmbedFooter
	Image       *EmbedImage
	Thumbnail   *EmbedThumbnail
	Video       *EmbedVideo
	Provider    *EmbedProvider
	Author      *EmbedAuthor
	Fields      []EmbedField
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	Url          string `json:"url,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Attachments struct {
	Id           string  `json:"id"`
	Filename     string  `json:"filename"`
	Title        string  `json:"title,omitempty"`
	Description  string  `json:"description,omitempty"`
	ContentType  string  `json:"content_type,omitempty"`
	Size         int     `json:"size"`
	Url          string  `json:"url"`
	Proxy_Url    string  `json:"proxy_url"`
	Height       int     `json:"height,omitempty"`
	Width        int     `json:"width,omitempty"`
	Ephemeral    bool    `json:"ephemeral,omitempty"`
	DurationSecs float64 `json:"duration_secs,omitempty"`
	Waveform     string  `json:"waveform,omitempty"`
	Flags        int     `json:"flags,omitempty"`
}

type ChannelMention struct {
	Id      string `json:"id"`
	GuildId string `json:"guild_id"`
	Type    int    `json:"type"`
	Name    string `json:"name"`
}

type Role struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Color        int        `json:"color"`
	Colors       *RoleColor `json:"colors"`
	Hoist        bool       `json:"hoist"`
	Icon         string     `json:"icon,omitempty"`
	UnicodeEmoji string     `json:"unicode_emoji,omitempty"`
	Position     int        `json:"position"`
	Permissions  string     `json:"permissions"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
	Tags         *RoleTags  `json:"tags,omitempty"`
	Flags        int        `json:"flags"`
}

type RoleTags struct {
	BotId                 string `json:"bot_id,omitempty"`
	IntegrationId         string `json:"integration_id,omitempty"`
	PremiumSubscriber     *bool  `json:"premium_subscriber,omitempty"`
	SubscriptionListingId string `json:"subscription_listing_id,omitempty"`
	AvailableForPurchase  *bool  `json:"available_for_purchase,omitempty"`
	GuildConnections      *bool  `json:"guild_connections,omitempty"`
}

type RoleColor struct {
	PrimaryColor   int  `json:"primary_color"`
	SecondaryColor *int `json:"secondary_color"`
	TertiaryColor  *int `json:"tertiary_color"`
}

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

type UserPrimaryGuild struct {
	IdentityGuildId string `json:"identity_guild_id"`
	IdentityEnabled bool   `json:"identity_enabled"`
	Tag             string `json:"tag"`
	Badge           string `json:"badge"`
}

type AvatarDecorationData struct {
	Asset string `json:"asset"`
	SkuId string `json:"sku_id"`
}

type Collectible struct {
	Nameplate *Nameplate `json:"nameplate"`
}

type Nameplate struct {
	SkuId   string `json:"sku_id"`
	Asset   string `json:"asset"`
	Label   string `json:"label"`
	Palette string `json:"pallete"`
}

type MessageReference struct {
	Type            int    `json:"type,omitempty"`
	MessageId       string `json:"message_id,omitempty"`
	ChannelId       string `json:"channel_id,omitempty"`
	GuildId         string `json:"guild_id,omitempty"`
	FailIfNotExists bool   `json:"fail_if_not_exists,omitempty"`
}

type MessageSnapshot struct {
	Message Message
}

type MessageInteractionMetadata struct {
	Id                            string                      `json:"id"`
	Type                          int                         `json:"type"`
	User                          *User                       `json:"user"`
	AuthorizingIntegrationOwners  []int                       `json:"authorizing_integration_owners"`
	OriginalResponseMessageId     string                      `json:"original_response_message_id,omitempty"`
	TargetUser                    *User                       `json:"target_user,omitempty"`
	TargetMessageId               string                      `json:"target_message_id,omitempty"`
	InteractedMessageId           string                      `json:"interacted_message_id,omitempty"`
	TriggeringInteractionMetadata *MessageInteractionMetadata `json:"triggering_interaction_metadata,omitempty"`
}

type MessageInteraction struct {
	Id     string  `json:"id"`
	Type   int     `json:"type"`
	Name   string  `json:"name"`
	User   *User   `json:"user"`
	Member *Member `json:"member"`
}

type Member struct {
	User                       *User                 `json:"user,omitempty"`
	Nick                       string                `json:"nick,omitempty"`
	Avatar                     string                `json:"avatar,omitempty"`
	Banner                     string                `json:"banner,omitempty"`
	Roles                      []string              `json:"roles"`
	JoinedAt                   string                `json:"joined_at"`
	PremiumSince               string                `json:"premium_since,omitempty"`
	Deaf                       bool                  `json:"deaf"`
	Mute                       bool                  `json:"mute"`
	Flags                      int                   `json:"flags"`
	Pending                    bool                  `json:"pending,omitempty"`
	Permissions                string                `json:"permissions,omitempty"`
	CommunicationDisabledUntil string                `json:"communication_disabled_until,omitempty"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
}

type Channel struct {
	Id                            string           `json:"id"`
	Type                          int              `json:"type"`
	GuildId                       string           `json:"guild_id,omitempty"`
	Position                      int              `json:"position,omitempty"`
	PermissionOverwrites          *Overwrite       `json:"permission_overwrites,omitempty"`
	Name                          string           `json:"name,omitempty"`
	Topic                         string           `json:"topic,omitempty"`
	NSFW                          bool             `json:"nsfw,omitempty"`
	LastMessageId                 string           `json:"last_message_id,omitempty"`
	Bitrate                       int              `json:"bitrate,omitempty"`
	UserLimit                     int              `json:"user_limit,omitempty"`
	RateLimitPerUser              int              `json:"rate_limit_per_user"`
	Recipients                    []User           `json:"recipients,omitempty"`
	Icon                          string           `json:"icon,omitempty"`
	OwnerId                       string           `json:"owner_id,omitempty"`
	ApplicationId                 string           `json:"application_id,omitempty"`
	Managed                       bool             `json:"managed,omitempty"`
	ParentId                      string           `json:"parent_id,omitempty"`
	LastPinTimestamp              string           `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     string           `json:"rtc_region,omitempty"`
	VideoQualityMode              int              `json:"video_quality_mode,omitempty"`
	MessageCount                  int              `json:"message_count,omitempty"`
	MemberCount                   int              `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata  `json:"thread_metadata,omitempty"`
	Member                        *ThreadMember    `json:"member,omitempty"`
	DefaultAutoArchiveDuration    int              `json:"default_auto_archive_duration,omitempty"`
	Permissions                   string           `json:"permissions,omitempty"`
	Flags                         int              `json:"flags,omitempty"`
	TotalMessageSent              int              `json:"total_message_sent,omitempty"`
	AvailableTags                 []Tag            `json:"available_tags,omitempty"`
	AppliedTags                   []string         `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser int              `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              int              `json:"default_sort_order,omitempty"`
	DefaultForumLayout            int              `json:"default_forum_layout,omitempty"`
}

type DefaultReaction struct {
	EmojiId   string `json:"emoji_id"`
	EmojiName string `json:"emoji_name"`
}

type Tag struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Moderated bool   `json:"moderated"`
	EmojiId   string `json:"emoji_id"`
	EmojiName string `json:"emoji_name"`
}

type ThreadMetadata struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked,omitempty"`
	Invitable           bool   `json:"invitable,omitempty"`
	CreateTimestamp     string `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	Id            string  `json:"id,omitempty"`
	UserId        string  `json:"user_id,omitempty"`
	JoinTimestamp string  `json:"join_timestamp"`
	Flags         int     `json:"flags"`
	Member        *Member `json:"member,omitempty"`
}

type Overwrite struct {
	Id    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}
