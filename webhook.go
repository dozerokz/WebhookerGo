package webhookergo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Webhook represents a Discord webhook message.
// It can contain content, username, avatar, TTS flag, and embeds.
type Webhook struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	TTS       bool    `json:"tts,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

// DiscordError represents an error returned by the Discord API.
type DiscordError struct {
	StatusCode int
	Body       string
	RetryAfter time.Duration // only set if 429 Too Many Requests
}

// Error returns a human-readable error string for DiscordError.
func (e *DiscordError) Error() string {
	if e.StatusCode == http.StatusTooManyRequests && e.RetryAfter > 0 {
		return fmt.Sprintf("discord webhook error: status=%d retry_after=%s body=%s",
			e.StatusCode, e.RetryAfter, e.Body)
	}
	return fmt.Sprintf("discord webhook error: status=%d body=%s", e.StatusCode, e.Body)
}

// NewWebhook creates a new Webhook object.
func NewWebhook() *Webhook {
	return &Webhook{}
}

// SetUsername sets the username for the webhook message.
func (w *Webhook) SetUsername(username string) *Webhook {
	w.Username = username
	return w
}

// SetAvatarURL sets the avatar URL for the webhook message.
func (w *Webhook) SetAvatarURL(url string) *Webhook {
	w.AvatarURL = url
	return w
}

// SetContent sets the main content (text) of the webhook message.
func (w *Webhook) SetContent(content string) *Webhook {
	w.Content = content
	return w
}

// SetTTS sets whether the message should be sent as a text-to-speech message.
func (w *Webhook) SetTTS(tts bool) *Webhook {
	w.TTS = tts
	return w
}

// AddEmbed appends an Embed to the webhook message.
func (w *Webhook) AddEmbed(embed *Embed) *Webhook {
	w.Embeds = append(w.Embeds, *embed)
	return w
}

// SendSimple sends a plain text message to the given webhook URL.
func SendSimple(webhookURL string, message string) error {
	w := NewWebhook().SetContent(message)
	return w.Send(webhookURL)
}

// SendEmbed sends a webhook message containing a single Embed to the given webhook URL.
func SendEmbed(webhookURL string, embed *Embed) error {
	w := NewWebhook().AddEmbed(embed)
	return w.Send(webhookURL)
}

// Send sends the webhook message to the specified Discord webhook URL.
// Returns a DiscordError if Discord returns a non-success status code.
func (w *Webhook) Send(webhookURL string) error {
	data, err := json.Marshal(w)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to post webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

		discErr := &DiscordError{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}

		// Handle 429 retry-after
		if resp.StatusCode == http.StatusTooManyRequests {
			if ra := resp.Header.Get("Retry-After"); ra != "" {
				if ms, err := strconv.Atoi(ra); err == nil {
					discErr.RetryAfter = time.Duration(ms) * time.Millisecond
				}
			}
		}

		return discErr
	}

	return nil
}
