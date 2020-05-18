# extratime

An extra time package for unmarshalling and marshalling json and xml.

[![ci](https://github.com/johejo/extratime/workflows/ci/badge.svg)](https://github.com/johejo/extratime/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/johejo/extratime/branch/master/graph/badge.svg)](https://codecov.io/gh/johejo/extratime)

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
package main_test

import (
    "encoding/json"
    "encoding/xml"
    "fmt"

    "github.com/johejo/extratime"
)

type A struct {
	XMLName xml.Name            `xml:"a"`
	Text    string              `xml:",chardata"`
	B       extratime.RFC1123   `xml:"b"`
} 

func main() {
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
```
