package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

var (
	repoPath string
)

func init() {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	repoPath = strings.TrimSpace(string(out))
}

func main() {
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}

func _main() error {
	layouts := []string{
		"RFC1123",
		"RFC1123Z",
		"RFC822",
		"RFC822Z",
		"RFC850",
		"Kitchen",
	}

	b := bytes.NewBuffer(nil)

	fprintf(b, "package extratime")
	fprintf(b, "// This file is auto-generated by internal/gen.go. DO NOT EDIT.")
	fprintf(b, "//go:generate go run internal/gen.go")

	// type alias and named type
	fprintf(b, "type (")
	fprintf(b, "    Time = time.Time")
	for _, l := range layouts {
		fprintf(b, "%s Time", l)
	}
	fprintf(b, ")")

	pkgs := []string{
		"json",
		"xml",
	}
	ifs := []string{
		"Unmarshaler",
		"Marshaler",
	}
	// interface checker
	fprintf(b, "var (")
	for _, l := range layouts {
		for _, p := range pkgs {
			for _, i := range ifs {
				fprintf(b, "_ %s.%s = (*%s)(nil)", p, i, l)
			}
		}
	}
	fprintf(b, ")")

	interfaces := map[string]struct {
		in, out string
	}{
		"MarshalJSON":   {in: "", out: "([]byte, error)"},
		"MarshalXML":    {in: "e *xml.Encoder, start xml.StartElement", out: "error"},
		"UnmarshalJSON": {in: "data []byte", out: "error"},
		"UnmarshalXML":  {in: "d *xml.Decoder, start xml.StartElement", out: "error"},
	}
	keys := make([]string, 0, len(interfaces))
	for k := range interfaces {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// implements
	for _, l := range layouts {
		for _, k := range keys {
			inf := interfaces[k]
			var rt string
			if strings.HasPrefix(k, "Marshal") {
				rt = ""
			} else {
				rt = "*"
			}
			fprintf(b, "func (i %s%s) %s(%s) %s {", rt, l, k, inf.in, inf.out)
			switch k {
			case "UnmarshalJSON":
				fprintf(b, "t, err := time.Parse(time.%s, cut(data))", l)
				fprintf(b, "if err != nil { return err }")
				fprintf(b, "*i = %s(t)", l)
				fprintf(b, "return nil")
			case "MarshalJSON":
				fprintf(b, "return []byte(`\"`+Time(i).Format(time.%s)+`\"`), nil", l)
			case "MarshalXML":
				fprintf(b, "return e.EncodeElement(Time(i).Format(time.%s), start)", l)
			case "UnmarshalXML":
				fprintf(b, "var s string")
				fprintf(b, "if err := d.DecodeElement(&s, &start); err != nil { return err }")
				fprintf(b, "t, err := time.Parse(time.%s, s)", l)
				fprintf(b, "if err != nil { return err }")
				fprintf(b, "*i = %s(t)", l)
				fprintf(b, "return nil")
			case "UnmarshalYAML":
				fprintf(b, "return nil")
			case "MarshalYAML":
				fprintf(b, "return nil, nil")
			}
			fprintf(b, "}")
		}
	}

	// test code
	t := bytes.NewBuffer(nil)
	fprintf(t, "package extratime")
	fprintf(t, "// This file is auto-generated by internal/gen.go. DO NOT EDIT.")
	fprintf(t, "//go:generate go run internal/gen.go")
	fprintf(t, `import "github.com/stretchr/testify/assert"`)

	for _, l := range layouts {
		fprintf(t, "func Test%s_json(t *testing.T) {", l)
		fprintf(t, "    j := fmt.Sprintf(`{\"t\": \"`+time.%s+`\"}`)", l)
		fprintf(t, "    t.Log(j)")
		fprintf(t, "    var m map[string]%s", l)
		fprintf(t, "    assert.NoError(t, json.Unmarshal([]byte(j), &m))")
		fprintf(t, "    b, err := json.Marshal(m)")
		fprintf(t, "    assert.Nil(t, err)")
		fprintf(t, "    assert.JSONEq(t, j, string(b))")
		fprintf(t, "}")
	}

	for _, l := range layouts {
		fprintf(t, "func Test%s_xml(t *testing.T) {", l)
		fprintf(t, "    type A struct {")
		fprintf(t, "        XMLName xml.Name `xml:\"a\"`")
		fprintf(t, "        Text    string   `xml:\",chardata\"`")
		fprintf(t, "        B       %s       `xml:\"b\"`", l)
		fprintf(t, "    }")
		fprintf(t, `    v := time.%s`, l)
		fprintf(t, "    x := `<a><b>`+v+`</b></a>`")
		fprintf(t, "    t.Log(x)")
		fprintf(t, "    var a A")
		fprintf(t, "    assert.NoError(t, xml.Unmarshal([]byte(x), &a))")
		fprintf(t, "    b, err := xml.Marshal(a)")
		fprintf(t, "    assert.Nil(t, err)")
		fprintf(t, "    assert.Equal(t, x, string(b))")
		fprintf(t, "}")
	}

	test := filepath.Join(repoPath, "extratime_gen_test.go")
	if err := WriteFormattedCodeToFile(test, t); err != nil {
		return err
	}

	dst := filepath.Join(repoPath, "extratime_gen.go")
	return WriteFormattedCodeToFile(dst, b)
}

func fprintf(w io.Writer, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(w, "\n"+format, args...)
}

// stolen from https://github.com/lestrrat-go/jwx/blob/master/internal/codegen/codegen.go
func WriteFormattedCodeToFile(filename string, src io.Reader) error {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return errors.Wrap(err, `failed to read from source`)
	}

	formatted, err := imports.Process("", buf, nil)
	if err != nil {
		scanner := bufio.NewScanner(bytes.NewReader(buf))
		lineno := 1
		for scanner.Scan() {
			txt := scanner.Text()
			fmt.Fprintf(os.Stdout, "%03d: %s\n", lineno, txt)
			lineno++
		}
		return errors.Wrap(err, `failed to format code`)
	}

	if dir := filepath.Dir(filename); dir != "." {
		if _, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return errors.Wrapf(err, `failed to create directory %s`, dir)
			}
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return errors.Wrapf(err, `failed to open %s.go`, filename)
	}
	defer f.Close()
	if _, err := f.Write(formatted); err != nil {
		return err
	}
	return nil
}
