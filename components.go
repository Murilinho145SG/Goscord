package goscord

import (
	"encoding/json"
	"errors"
)

type ComponentsType uint

const (
	ActionsRowType        ComponentsType = 1
	ButtonType            ComponentsType = 2
	StringSelectType      ComponentsType = 3
	TextInputType         ComponentsType = 4
	UserSelectType        ComponentsType = 5
	RoleSelectType        ComponentsType = 6
	MentionableSelectType ComponentsType = 7
	ChannelSelectType     ComponentsType = 8
	SectionType           ComponentsType = 9
	TextDisplayType       ComponentsType = 10
	ThumbnailType         ComponentsType = 11
	MediaGalleryType      ComponentsType = 12
	FileType              ComponentsType = 13
	SeparatorType         ComponentsType = 14
	ContainerType         ComponentsType = 17
)

type Component interface {
	GetType() ComponentsType
}

type RawComponent struct {
	Type ComponentsType  `json:"type"`
	Raw  json.RawMessage `json:"-"`
}

func (c *RawComponent) UnmarshalJSON(data []byte) error {
	type Alias RawComponent
	aux := &struct{ *Alias }{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Raw = data
	return nil
}

func (c *RawComponent) ToComponent() (Component, error) {
	switch c.Type {
	case ActionsRowType:
		var row ActionRow
		if err := json.Unmarshal(c.Raw, &row); err != nil {
			return nil, err
		}
		return &row, nil
	case ButtonType:
		var btn Button
		if err := json.Unmarshal(c.Raw, &btn); err != nil {
			return nil, err
		}
		return &btn, nil
	case StringSelectType:
		var stringSelect StringSelect
		if err := json.Unmarshal(c.Raw, &stringSelect); err != nil {
			return nil, err
		}
		return &stringSelect, nil
	case TextInputType:
		var textInput TextInput
		if err := json.Unmarshal(c.Raw, &textInput); err != nil {
			return nil, err
		}
		return &textInput, nil
	case UserSelectType:
		var userSelect UserSelect
		if err := json.Unmarshal(c.Raw, &userSelect); err != nil {
			return nil, err
		}
		return &userSelect, nil
	case RoleSelectType:
		var roleSelect RoleSelect
		if err := json.Unmarshal(c.Raw, &roleSelect); err != nil {
			return nil, err
		}
		return &roleSelect, nil
	case MentionableSelectType:
		var mentionableSelect MentionableSelect
		if err := json.Unmarshal(c.Raw, &mentionableSelect); err != nil {
			return nil, err
		}
		return &mentionableSelect, nil
	case ChannelSelectType:
		var channelSelect ChannelSelect
		if err := json.Unmarshal(c.Raw, &channelSelect); err != nil {
			return nil, err
		}
		return &channelSelect, nil
	case SectionType:
		var section Section
		if err := json.Unmarshal(c.Raw, &section); err != nil {
			return nil, err
		}
		return &section, nil
	case TextDisplayType:
		var textDisplay TextDisplay
		if err := json.Unmarshal(c.Raw, &textDisplay); err != nil {
			return nil, err
		}
		return &textDisplay, nil
	case ThumbnailType:
		var thumbnail Thumbnail
		if err := json.Unmarshal(c.Raw, &thumbnail); err != nil {
			return nil, err
		}
		return &thumbnail, nil
	case MediaGalleryType:
		var mediaGallery MediaGallery
		if err := json.Unmarshal(c.Raw, &mediaGallery); err != nil {
			return nil, err
		}
		return &mediaGallery, nil
	case FileType:
		var file File
		if err := json.Unmarshal(c.Raw, &file); err != nil {
			return nil, err
		}
		return &file, nil
	case SeparatorType:
		var separator Separator
		if err := json.Unmarshal(c.Raw, &separator); err != nil {
			return nil, err
		}
		return &separator, nil
	case ContainerType:
		var container Container
		if err := json.Unmarshal(c.Raw, &container); err != nil {
			return nil, err
		}
		return &container, nil
	}

	return nil, errors.New("unknown component type")
}

type ActionRow struct {
	Type       ComponentsType `json:"type"`
	Id         int            `json:"id,omitempty"`
	Components []Component    `json:"components"`
}

func (a *ActionRow) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type ButtonStyle int

const (
	ButtonPrimary   ButtonStyle = 1
	ButtonSecondary ButtonStyle = 2
	ButtonSuccess   ButtonStyle = 3
	ButtonDanger    ButtonStyle = 4
	ButtonLink      ButtonStyle = 5
	ButtonPremium   ButtonStyle = 6
)

type Button struct {
	Type     ComponentsType `json:"type"`
	Id       int            `json:"id,omitempty"`
	Style    ButtonStyle    `json:"style"`
	Label    string         `json:"label,omitempty"`
	Emoji    *Emoji         `json:"emoji,omitempty"`
	CustomId string         `json:"custom_id"`
	SkuId    string         `json:"sku_id,omitempty"`
	Url      string         `json:"url,omitempty"`
	Disable  bool           `json:"disable,omitempty"`
}

func (a *Button) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type StringSelect struct {
	Type        ComponentsType `json:"type"`
	Id          int            `json:"id,omitempty"`
	CustomId    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   int            `json:"min_values,omitempty"`
	MaxValues   int            `json:"max_values,omitempty"`
	Disable     bool           `json:"disable,omitempty"`
}

func (a *StringSelect) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default,omitempty"`
}

type TextInput struct {
	Type        ComponentsType `json:"type"`
	Id          int            `json:"id,omitempty"`
	CustomId    string         `json:"custom_id"`
	Style       int            `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Value       string         `json:"value,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
}

func (a *TextInput) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type UserSelect struct {
	Type          ComponentsType `json:"type"`
	Id            int            `json:"id,omitempty"`
	CustomId      string         `json:"custom_id"`
	Placeholder   string         `json:"placeholder,omitempty"`
	DefaultValues []DefaultValue `json:"default_values,omitempty"`
	MinValues     int            `json:"min_values,omitempty"`
	MaxValues     int            `json:"max_values,omitempty"`
	Disable       bool           `json:"disabled,omitempty"`
}

func (a *UserSelect) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type DefaultValue struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type RoleSelect struct {
	Type          ComponentsType `json:"type"`
	Id            int            `json:"id,omitempty"`
	CustomId      string         `json:"custom_id"`
	Placeholder   string         `json:"placeholder,omitempty"`
	DefaultValues []DefaultValue `json:"default_values,omitempty"`
	MinValues     int            `json:"min_values,omitempty"`
	MaxValues     int            `json:"max_values,omitempty"`
	Disabled      bool           `json:"disabled,omitempty"`
}

func (a *RoleSelect) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type MentionableSelect struct {
	Type          ComponentsType `json:"type"`
	Id            int            `json:"id,omitempty"`
	CustomId      string         `json:"custom_id"`
	Placeholder   string         `json:"placeholder,omitempty"`
	DefaultValues []DefaultValue `json:"default_values,omitemtpy"`
	MinValues     int            `json:"min_values,omitempty"`
	MaxValues     int            `json:"max_values,omitempty"`
	Disabled      bool           `json:"disabled,omitempty"`
}

func (a *MentionableSelect) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type ChannelSelect struct {
	Type          ComponentsType `json:"type"`
	Id            int            `json:"id,omitempty"`
	CustomId      string         `json:"custom_id"`
	ChannelTypes  []int          `json:"channel_types,omitempty"`
	Placeholder   string         `json:"placeholder,omitempty"`
	DefaultValues []DefaultValue `json:"default_values,omitempty"`
	MinValues     int            `json:"min_values,omitempty"`
	MaxValues     int            `json:"max_values,omitempty"`
	Disabled      bool           `json:"disabled,omitempty"`
}

func (a *ChannelSelect) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type Section struct {
	Type       ComponentsType `json:"type"`
	Id         int            `json:"id,omitempty"`
	Components []TextDisplay  `json:"components"`
	Accessory  []Component    `json:"accessory"`
}

func (a *Section) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type TextDisplay struct {
	Type    ComponentsType `json:"type"`
	Id      int            `json:"id,omitempty"`
	Content string         `json:"content"`
}

func (a *TextDisplay) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type Thumbnail struct {
	Type        ComponentsType     `json:"type"`
	Id          int                `json:"id,omitempty"`
	Media       *UnfurledMediaItem `json:"media"`
	Description string             `json:"description,omitempty"`
	Spoiler     bool               `json:"spoiler,omitempty"`
}

func (a *Thumbnail) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type UnfurledMediaItem struct {
	Url          string `json:"url"`
	ProxyUrl     string `json:"proxy_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
	ContentType  string `json:"content_type,omitempty"`
	AttachmentId string `json:"attachment_id,omitempty"`
}

type MediaGallery struct {
	Type  ComponentsType `json:"type"`
	Id    int            `json:"id,omitempty"`
	Items []MediaGalleryItem
}

func (a *MediaGallery) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type MediaGalleryItem struct {
	Media       *UnfurledMediaItem `json:"media"`
	Description string             `json:"description,omitempty"`
	Spoiler     bool               `json:"spoiler,omitempty"`
}

type File struct {
	Type    ComponentsType     `json:"type"`
	Id      int                `json:"id,omitempty"`
	File    *UnfurledMediaItem `json:"file"`
	Spoiler bool               `json:"spoiler,omitempty"`
	Name    string             `json:"name"`
	Size    int                `json:"size"`
}

func (a *File) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type Separator struct {
	Type    ComponentsType `json:"type"`
	Id      int            `json:"id,omitempty"`
	Divider bool           `json:"divider,omitempty"`
	Spacing int            `json:"spacing,omitempty"`
}

func (a *Separator) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type Container struct {
	Type        ComponentsType `json:"type"`
	Id          int            `json:"id,omitempty"`
	Components  []Component    `json:"components"`
	AccentColor int            `json:"accent_color,omitempty"`
	Spoiler     bool           `json:"spoiler,omitempty"`
}

func (a *Container) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

func Raw(c Component) RawComponent {
	assertType(c)
	cType := c.GetType()
	data, _ := json.Marshal(c)

	return RawComponent{
		Type: cType,
		Raw:  data,
	}
}
