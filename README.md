# extratime

An extra time package for unmarshalling and marshalling json and xml.

[![ci](https://github.com/johejo/extratime/workflows/ci/badge.svg)](https://github.com/johejo/extratime/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/johejo/extratime/branch/master/graph/badge.svg)](https://codecov.io/gh/johejo/extratime)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/johejo/extratime)
[![Go Report Card](https://goreportcard.com/badge/johejo/extratime)](https://goreportcard.com/report/johejo/extratime)

## Supported Formats

- RFC1123
- RFC1123Z
- RFC822
- RFC822Z
- Kitchen

## Install

```
go get github.com/johejo/extratime
```

## Example

```go
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
	const x = "<a><b>Mon, 02 Jan 2006 15:04:05 MST</b></a>"
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
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		panic(err)
	}
	jb, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jb))

	// Output:
	// <a><b>Mon, 02 Jan 2006 15:04:05 MST</b></a>
	// {"t":"Mon, 02 Jan 2006 15:04:05 MST"}
}
```

## License

MIT

## Author

Mitsuo Heijo (@johejo)
