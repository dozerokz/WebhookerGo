package webhookergo

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// minDecimalRGBValue is the smallest valid decimal RGB color (0).
	minDecimalRGBValue = 0

	// maxDecimalRGBValue is the largest valid decimal RGB color (16777215, or 0xFFFFFF).
	maxDecimalRGBValue = 16777215

	// minRGBValue is the minimum valid value for an individual R/G/B component (0).
	minRGBValue = 0

	// maxRGBValue is the maximum valid value for an individual R/G/B component (255).
	maxRGBValue = 255
)

// Embed represents a rich embed object for Discord.
type Embed struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Timestamp   string    `json:"timestamp,omitempty"`
	Color       int       `json:"color,omitempty"`
	Footer      Footer    `json:"footer,omitempty"`
	Image       Image     `json:"image,omitempty"`
	Thumbnail   Thumbnail `json:"thumbnail,omitempty"`
	Author      Author    `json:"author,omitempty"`
	Fields      []Field   `json:"fields,omitempty"`
	err         error
}

// Footer represents the footer section of an embed.
type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Image represents an image in an embed.
type Image struct {
	URL string `json:"url,omitempty"`
}

// Thumbnail represents a thumbnail image in an embed.
type Thumbnail struct {
	URL string `json:"url,omitempty"`
}

// Author represents the author section of an embed.
type Author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Field represents a key-value pair in the embed's fields section.
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// RGB is a struct representing an RGB color.
type RGB struct {
	R, G, B int
}

// NewEmbed creates a new Embed object.
func NewEmbed() *Embed {
	return &Embed{}
}

// Error returns the last error that occurred when configuring the embed, or nil if no errors have occurred.
func (e *Embed) Error() error {
	return e.err
}

// SetTitle sets the title of the embed.
func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

// SetDescription sets the description text of the embed.
func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

// SetURL sets the url for the embed.
// If provided, the title becomes a clickable link.
func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

// SetColorInt sets the embed color using an integer color representation (0 <= color <= 16777215).
func (e *Embed) SetColorInt(color int) *Embed {
	switch {
	case color > maxDecimalRGBValue:
		e.Color = maxDecimalRGBValue
	case color < minDecimalRGBValue:
		e.Color = minDecimalRGBValue
	default:
		e.Color = color
	}
	return e
}

// SetColorHex sets the embed color using a hex string (e.g. "#FF5733").
// If the hex string is invalid, the error can be retrieved with Error().
func (e *Embed) SetColorHex(hex string) *Embed {
	color, err := hexToColorInt(hex)
	if err != nil {
		e.err = err
	}
	e.Color = color
	return e
}

// SetColorRGB sets the embed color using an RGB type (e.g. webhook.RGB{R: 100, G: 200, B: 300}).
func (e *Embed) SetColorRGB(rgb RGB) *Embed {
	e.Color = rgbToInt(rgb)
	return e
}

// SetFooter sets the footer text and optional icon url for the embed.
// If iconURL is an empty string, only the text will be displayed.
func (e *Embed) SetFooter(text, iconURL string) *Embed {
	e.Footer = Footer{Text: text, IconURL: iconURL}
	return e
}

// SetTimestamp sets the embed timestamp to the given time.
// The time is always converted to UTC and formatted according to RFC3339.
func (e *Embed) SetTimestamp(t time.Time) *Embed {
	e.Timestamp = t.UTC().Format(time.RFC3339)
	return e
}

// SetTimestampNow sets the embed timestamp to the current time (UTC).
func (e *Embed) SetTimestampNow() *Embed {
	return e.SetTimestamp(time.Now())
}

// SetAuthorName sets the name of the embed's author section.
func (e *Embed) SetAuthorName(name string) *Embed {
	e.Author.Name = name
	return e
}

// SetAuthorURL sets the url of the embed's author section.
// The authors name becomes a clickable link.
func (e *Embed) SetAuthorURL(url string) *Embed {
	e.Author.URL = url
	return e
}

// SetAuthorIcon sets the icon url for the embed's author section.
// This is the small image displayed next to the author's name.
func (e *Embed) SetAuthorIcon(iconURL string) *Embed {
	e.Author.IconURL = iconURL
	return e
}

// SetImage sets the url of the embed's main image.
func (e *Embed) SetImage(url string) *Embed {
	e.Image.URL = url
	return e
}

// AddField adds a new field to the embed.
// Fields are key-value pairs displayed in the embed.
func (e *Embed) AddField(field *Field) *Embed {
	e.Fields = append(e.Fields, *field)
	return e
}

// NewEmptyField creates a new Field with default empty values.
func NewEmptyField() *Field {
	return &Field{}
}

// NewField creates a new Field with the given name, value, and inline flag.
func NewField(name, value string, inline bool) *Field {
	return &Field{Name: name, Value: value, Inline: inline}
}

// SetName sets the name of the field.
func (f *Field) SetName(name string) *Field {
	f.Name = name
	return f
}

// SetValue sets the value of the field.
func (f *Field) SetValue(value string) *Field {
	f.Value = value
	return f
}

// SetInline sets whether the field should be displayed inline.
func (f *Field) SetInline(inline bool) *Field {
	f.Inline = inline
	return f
}

// ClearFields removes all fields from the embed.
func (e *Embed) ClearFields() *Embed {
	e.Fields = []Field{}
	return e
}

// SetThumbnail sets the url of the embed's thumbnail image.
func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail.URL = url
	return e
}

// hexToColorInt converts a hex string to an integer representation
func hexToColorInt(hex string) (int, error) {
	if len(hex) == 7 && hex[0] == '#' {
		hex = hex[1:]
	} else if len(hex) != 6 {
		return 0, fmt.Errorf("invalid hex color format")
	}

	// Parse the hex string to an integer
	decimalValue, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return 0, fmt.Errorf("error parsing hex to int: %v", err)
	}

	// Return the integer value directly
	return int(decimalValue), nil
}

// rgbToInt converts an RGB struct to an integer representation
func rgbToInt(rgb RGB) int {
	switch {
	case rgb.R > maxRGBValue:
		rgb.R = maxRGBValue
	case rgb.G > maxRGBValue:
		rgb.G = maxRGBValue
	case rgb.B > maxRGBValue:
		rgb.B = maxRGBValue
	case rgb.R < minRGBValue:
		rgb.R = minRGBValue
	case rgb.G < minRGBValue:
		rgb.G = minRGBValue
	case rgb.B < minRGBValue:
		rgb.B = minRGBValue
	}

	return (rgb.R << 16) | (rgb.G << 8) | rgb.B
}
