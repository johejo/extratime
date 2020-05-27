// Package extratime is an extra package for unmarshalling and marshalling time format to json and xml.
package extratime

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

func (u UnixTimeStamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(u.Time().Unix(), 10) + `"`), nil
}

func (u UnixTimeStamp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strconv.FormatInt(u.Time().Unix(), 10), start)
}

func (u *UnixTimeStamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(trim(data), 10, 64)
	if err != nil {
		return err
	}
	*u = UnixTimeStamp(time.Unix(i, 0))
	return nil
}

func (u *UnixTimeStamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*u = UnixTimeStamp(time.Unix(i, 0))
	return nil
}

func trim(b []byte) string {
	return strings.Trim(string(b), `"`)
}
