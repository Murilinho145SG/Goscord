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
	}

	return nil, errors.New("unknown component type")
}

type ActionRow struct {
	Type int `json:"type"`
	// Id         int            `json:"id,omitempty"`
	Components []RawComponent `json:"components"`
}

func (a *ActionRow) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type Button struct {
	Type     int    `json:"type"`
	Id       int    `json:"id,omitempty"`
	Style    int    `json:"style"`
	Label    string `json:"label,omitempty"`
	Emoji    *Emoji `json:"emoji,omitempty"`
	CustomId string `json:"custom_id"`
	SkuId    string `json:"sku_id,omitempty"`
	Url      string `json:"url,omitempty"`
	Disable  bool   `json:"disable,omitempty"`
}

func (a *Button) GetType() ComponentsType {
	return ComponentsType(a.Type)
}

type StringSelect struct {
	Type        int            `json:"type"`
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
	Type        int    `json:"type"`
	Id          int    `json:"id,omitempty"`
	CustomId    string `json:"custom_id"`
	Style       int    `json:"style"`
	Label       string `json:"label"`
	MinLength   int    `json:"min_length,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Value       string `json:"value,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
}

func (a *TextInput) GetType() ComponentsType {
	return ComponentsType(a.Type)
}
