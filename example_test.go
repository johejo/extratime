package extratime_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/johejo/extratime"
)

type A struct {
	XMLName xml.Name          `xml:"a"`
	Text    string            `xml:",chardata"`
	B       extratime.RFC1123 `xml:"b"`
}

func Example() {
	// xml
	const x = "<a><b>Mon, 02 Jan 2006 15:04:05 MST</a></b>"
	var a A
	if err := xml.Unmarshal([]byte(x), &a); err != nil {
		panic(err)
	}
	xb, err := xml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xb))

	// json
	const j = `{"t": "Mon, 02 Jan 2006 15:04:05 MST"}`
	var m map[string]extratime.RFC1123
	if err := json.Unmarshal([]byte(x), &m); err != nil {
		panic(err)
	}
	jb, err := xml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jb))
}
