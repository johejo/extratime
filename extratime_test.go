package extratime

import (
	"encoding/json"
	"encoding/xml"
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
		x := `<a><b>` + strconv.FormatInt(v, 10) + `</b></a>`
		t.Log(x)
		var a A
		assert.NoError(t, xml.Unmarshal([]byte(x), &a))
		b, err := xml.Marshal(a)
		assert.Nil(t, err)
		assert.Equal(t, x, string(b))
	})
	t.Run("UnixNano", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).UnixNano()
		x := `<a><b>` + strconv.FormatInt(v, 10) + `</b></a>`
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
		j := `{"t": "` + strconv.FormatInt(v, 10) + `"}`
		t.Log(j)
		var m map[string]UnixTimeStamp
		assert.NoError(t, json.Unmarshal([]byte(j), &m))
		b, err := json.Marshal(m)
		assert.Nil(t, err)
		assert.JSONEq(t, j, string(b))
	})
	t.Run("UnixName", func(t *testing.T) {
		v := time.Date(2020, 5, 20, 1, 23, 44, 33, time.Local).UnixNano()
		j := `{"t": "` + strconv.FormatInt(v, 10) + `"}`
		t.Log(j)
		var m map[string]UnixTimeStamp
		assert.NoError(t, json.Unmarshal([]byte(j), &m))
		b, err := json.Marshal(m)
		assert.Nil(t, err)
		assert.JSONEq(t, j, string(b))
	})
}

func Test_cut(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{`"hello"`, `hello`},
		{`world`, `world`},
	}
	for _, tt := range tests {
		t.Run(tt.in+`->`+tt.out, func(t *testing.T) {
			got := cut([]byte(tt.in))
			assert.Equal(t, tt.out, got)
		})
	}
}
