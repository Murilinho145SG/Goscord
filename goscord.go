package goscord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"github.com/Murilinho145SG/gouter/log"
	"github.com/gorilla/websocket"
)

const (
	VERSION = "10"
)

type Bot struct {
	token        string
	c            *websocket.Conn
	events       map[reflect.Type][]reflect.Value
	url          string
	lastSequence string
	sessionId    string
}

type Payload struct {
	T  *string                `json:"t"`
	S  *int                   `json:"s"`
	OP int                    `json:"op"`
	D  map[string]interface{} `json:"d"`
}

type D struct {
	HeartbeatInterval int      `json:"heartbeat_interval"`
	Trace             []string `json:"_trace"`
}

func New(token string) *Bot {
	return &Bot{
		token:  token,
		events: make(map[reflect.Type][]reflect.Value),
	}
}

func (b *Bot) Start() error {
	b.url = fmt.Sprintf("wss://gateway.discord.gg/?v=%s&encoding=json", VERSION)
	c, _, err := websocket.DefaultDialer.Dial(b.url, http.Header{})
	if err != nil {
		return err
	}
	b.c = c

	_, msg, err := c.ReadMessage()
	if err != nil {
		return err
	}

	var payload Payload
	if err := json.Unmarshal(msg, &payload); err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(time.Duration(payload.D["heartbeat_interval"].(float64)) * time.Millisecond)
		for range ticker.C {
			heartbeat := map[string]any{
				"op": 1,
				"d":  nil,
			}
			c.WriteJSON(heartbeat)
			fmt.Println("→ Heartbeat enviado")
		}
	}()

	identify := map[string]any{
		"op": 2,
		"d": map[string]any{
			"token":      b.token,
			"session_id": b.sessionId,
			"intents":    131071,
			"properties": map[string]string{
				"os":      "windows",
				"browser": "goscord",
				"device":  "goscord",
			},
		},
	}
	if err := c.WriteJSON(identify); err != nil {
		return err
	}

	fmt.Println("→ Identify enviado")

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					go b.reconnect()
					continue
				}
				log.Error(err)
				continue
			}

			var p Payload
			if err := json.Unmarshal(msg, &p); err != nil {
				log.Error(err)
				return
			}

			if p.T == nil {
				continue
			}

			log.Info(*p.T)
			switch *p.T {
			case "READY":
				{
					b.sessionId = p.D["session_id"].(string)
				}
			case "RESUMED":
				{
				}

			// Guilds
			case "GUILD_CREATE":
				{
				}
			case "GUILD_DELETE":
				{
				}
			case "GUILD_UPDATE":
				{
				}
			case "GUILD_AUDIT_LOG_ENTRY_CREATE":
				{
				}

			// Guild Roles
			case "GUILD_ROLE_CREATE":
				{
				}
			case "GUILD_ROLE_UPDATE":
				{
				}
			case "GUILD_ROLE_DELETE":
				{
				}

			// Guild Members
			case "GUILD_MEMBER_ADD":
				{
				}
			case "GUILD_MEMBER_UPDATE":
				{
				}
			case "GUILD_MEMBER_REMOVE":
				{
				}

			// Guild Moderation
			case "GUILD_BAN_ADD":
				{
				}
			case "GUILD_BAN_REMOVE":
				{
				}
			case "GUILD_EMOJIS_UPDATE":
				{
				}
			case "GUILD_STICKERS_UPDATE":
				{
				}
			case "GUILD_INTEGRATIONS_UPDATE":
				{
				}
			case "GUILD_WEBHOOKS_UPDATE":
				{
				}

			// Guild Scheduled Events
			case "GUILD_SCHEDULED_EVENT_CREATE":
				{
				}
			case "GUILD_SCHEDULED_EVENT_UPDATE":
				{
				}
			case "GUILD_SCHEDULED_EVENT_DELETE":
				{
				}
			case "GUILD_SCHEDULED_EVENT_USER_ADD":
				{
				}
			case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
				{
				}

			// Invite
			case "INVITE_CREATE":
				{
				}
			case "INVITE_DELETE":
				{
				}

			// Thread
			case "THREAD_CREATE":
				{
				}
			case "THREAD_UPDATE":
				{
				}
			case "THREAD_DELETE":
				{
				}
			case "THREAD_LIST_SYNC":
				{
				}
			case "THREAD_MEMBER_UPDATE":
				{
				}
			case "THREAD_MEMBERS_UPDATE":
				{
				}

			// Messages
			case "MESSAGE_CREATE":
				{
					b.TriggerEvent(NewMessageCreate(b, p.D))
				}
			case "MESSAGE_UPDATE":
				{
				}
			case "MESSAGE_DELETE":
				{
				}
			case "MESSAGE_DELETE_BULK":
				{
				}
			case "MESSAGE_REACTION_ADD":
				{
				}
			case "MESSAGE_REACTION_REMOVE":
				{
				}
			case "MESSAGE_REACTION_REMOVE_ALL":
				{
				}
			case "MESSAGE_REACTION_REMOVE_EMOJI":
				{
				}

			// Channels
			case "CHANNEL_CREATE":
				{
				}
			case "CHANNEL_UPDATE":
				{
				}
			case "CHANNEL_DELETE":
				{
				}
			case "CHANNEL_PINS_UPDATE":
				{
				}

			// Voices
			case "VOICE_SERVER_UPDATE":
				{
				}
			case "VOICE_STATE_UPDATE":
				{
				}
			case "VOICE_CHANNEL_STATUS_UPDATE":
				{
				}

			// User
			case "USER_UPDATE":
				{
				}

			// Interaction
			case "INTERACTION_CREATE":
				{
					b.TriggerEvent(NewInteractionCreate(b, p.D))
				}

			// Others
			case "PRESENCE_UPDATE":
				{
				}
			case "TYPING_START":
				{
				}
			}
		}
	}()

	return nil
}

func (b *Bot) reconnect() {
	time.Sleep(5 * time.Second)

	c, _, err := websocket.DefaultDialer.Dial(b.url, http.Header{})
	if err != nil {
		log.Error(err)
		return
	}
	b.c = c

	if b.sessionId != "" {
		resume := map[string]any{
			"op": 6,
			"d": map[string]any{
				"token":      b.token,
				"session_id": b.sessionId,
			},
		}
		c.WriteJSON(resume)
	}
}

func (b *Bot) Stop() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	return b.Close()
}

func (b *Bot) Close() error {
	return b.c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "shutdown"), time.Now().Add(time.Second))
}

func (b *Bot) OnEvent(event interface{}) {
	t := reflect.TypeOf(event)
	if t.Kind() != reflect.Func {
		return
	}

	paramType := t.In(0)
	b.events[paramType] = append(b.events[paramType], reflect.ValueOf(event))
}

func (b *Bot) TriggerEvent(event interface{}) {
	t := reflect.TypeOf(event)

	for _, handler := range b.events[t] {
		go handler.Call([]reflect.Value{reflect.ValueOf(event)})
	}
}

func (b *Bot) sendPost(endpoint string, body []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", "https://discord.com/api/v"+VERSION+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+b.token)
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	if r.StatusCode == 429 {
		return b.sendPost(endpoint, body)
	}

	return r.Body, nil
}

func (b *Bot) sendGet(endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", "https://discord.com/api/v"+VERSION+endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+b.token)
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode == 429 {
		return b.sendGet(endpoint)
	}

	return r.Body, nil
}
