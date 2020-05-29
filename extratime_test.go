package extratime

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnixTimeStamp_xml(t *testing.T) {
	type A struct {
		XMLName xml.Name      `xml:"a"`
		Text    string        `xml:",chardata"`
		B       UnixTimeStamp `xml:"b"`
	}
	t.Run("Unix", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).Unix()
		x := fmt.Sprintf(`<a><b>%s</b></a>`, strconv.FormatInt(v, 10))
		t.Log(x)
		var a A
		assert.NoError(t, xml.Unmarshal([]byte(x), &a))
		b, err := xml.Marshal(a)
		assert.Nil(t, err)
		assert.Equal(t, x, string(b))
	})
	t.Run("UnixNano", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).UnixNano()
		x := fmt.Sprintf(`<a><b>%s</b></a>`, strconv.FormatInt(v, 10))
		t.Log(x)
		var a A
		assert.NoError(t, xml.Unmarshal([]byte(x), &a))
		b, err := xml.Marshal(a)
		assert.Nil(t, err)
		assert.Equal(t, x, string(b))
	})
}

func TestUnixTimeStamp_json(t *testing.T) {
	t.Run("Unix", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).Unix()
		vint := strconv.FormatInt(v, 10)
		want := fmt.Sprintf(`{"t": %s}`, vint)

		tests := []struct {
			name string
			j    string
		}{
			{
				name: "as string",
				j:    fmt.Sprintf(`{"t": "%s"}`, vint), // with double quote
			},
			{
				name: "as number",
				j:    fmt.Sprintf(`{"t": %s}`, vint),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var m map[string]UnixTimeStamp
				assert.NoError(t, json.Unmarshal([]byte(tt.j), &m))
				b, err := json.Marshal(m)
				assert.Nil(t, err)
				assert.JSONEq(t, want, string(b))
			})
		}
	})

	t.Run("UnixNano", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).UnixNano()
		vint := strconv.FormatInt(v, 10)
		want := fmt.Sprintf(`{"t": %s}`, vint)

		tests := []struct {
			name string
			j    string
		}{
			{
				name: "as string",
				j:    fmt.Sprintf(`{"t": "%s"}`, vint), // with double quote
			},
			{
				name: "as number",
				j:    fmt.Sprintf(`{"t": %s}`, vint),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var m map[string]UnixTimeStamp
				assert.NoError(t, json.Unmarshal([]byte(tt.j), &m))
				b, err := json.Marshal(m)
				assert.Nil(t, err)
				assert.JSONEq(t, want, string(b))
			})
		}
	})
}

func Test_trim(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{`"hello"`, `hello`},
		{`world`, `world`},
	}
	for _, tt := range tests {
		t.Run(tt.in+`->`+tt.out, func(t *testing.T) {
			got := trim([]byte(tt.in))
			assert.Equal(t, tt.out, got)
		})
	}
}
