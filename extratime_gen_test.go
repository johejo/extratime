package extratime

// This file is auto-generated by internal/gen.go. DO NOT EDIT.
//go:generate go run internal/gen.go
import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRFC1123_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.RFC1123 + `"}`)
	t.Log(j)
	var m map[string]RFC1123
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestRFC1123Z_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.RFC1123Z + `"}`)
	t.Log(j)
	var m map[string]RFC1123Z
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestRFC822_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.RFC822 + `"}`)
	t.Log(j)
	var m map[string]RFC822
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestRFC822Z_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.RFC822Z + `"}`)
	t.Log(j)
	var m map[string]RFC822Z
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestRFC850_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.RFC850 + `"}`)
	t.Log(j)
	var m map[string]RFC850
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestKitchen_json(t *testing.T) {
	j := fmt.Sprintf(`{"t": "` + time.Kitchen + `"}`)
	t.Log(j)
	var m map[string]Kitchen
	assert.NoError(t, json.Unmarshal([]byte(j), &m))
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.JSONEq(t, j, string(b))
}
func TestRFC1123_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       RFC1123  `xml:"b"`
	}
	v := time.RFC1123
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
func TestRFC1123Z_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       RFC1123Z `xml:"b"`
	}
	v := time.RFC1123Z
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
func TestRFC822_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       RFC822   `xml:"b"`
	}
	v := time.RFC822
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
func TestRFC822Z_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       RFC822Z  `xml:"b"`
	}
	v := time.RFC822Z
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
func TestRFC850_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       RFC850   `xml:"b"`
	}
	v := time.RFC850
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
func TestKitchen_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name `xml:"a"`
		Text    string   `xml:",chardata"`
		B       Kitchen  `xml:"b"`
	}
	v := time.Kitchen
	x := `<a><b>` + v + `</b></a>`
	t.Log(x)
	var a A
	assert.NoError(t, xml.Unmarshal([]byte(x), &a))
	b, err := xml.Marshal(a)
	assert.Nil(t, err)
	assert.Equal(t, x, string(b))
}
