package gotify

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

// Gotify provides "gotification" of domain specific identifier names:
//   underscored names translates into camel case ones.
// How the translation is done:
//   idenitifier name is splitted into chunks and each chunk is to be lowered, samples:
//   abc_def -> [abc, def]
//   AbdDef  -> [abd, def]
//   userId  -> [user, id]
// Then there are choice:
//   package name (just concat)
//   private identifiers (first chunk is kept lowered, the rest is treated as for public identifiers)
//   public identifiers:
// Eeach chunk is looked up in provided dictionary (id -> ID translation is always here)
// and is replaced with either found translation or just with titled word and then concatenated
// into the string
//   abc_def -> AbdDef
//   AbdDef  -> AbdDef
//   userId  -> UserID
type Gotify struct {
	dict map[string]string
}

// New constructs Gotify object
func New(src map[string]string) *Gotify {
	if src == nil {
		src = map[string]string{}
	}
	src["id"] = "ID"
	return &Gotify{
		dict: src,
	}
}

func acceptableHead(value byte) bool {
	return ('a' <= value && value <= 'z') || ('A' <= value && value <= 'Z') || value == '_'
}

func acceptableTail(value byte) bool {
	return acceptableHead(value) || ('0' <= value && value <= '9')
}

func filter(rawData string) string {
	buf := &bytes.Buffer{}
	rawData = strings.Replace(rawData, ".", "_", -1)
	data := []byte(rawData)
	for i, value := range data {
		if acceptableHead(value) {
			if err := buf.WriteByte(value); err != nil {
				log.Fatal(err)
			}
			for _, val := range data[i+1:] {
				if acceptableTail(val) || val == ' ' {
					if err := buf.WriteByte(val); err != nil {
						log.Fatal(err)
					}
				}
			}
			break
		}
	}
	return string(buf.Bytes())
}

var goishMask *regexp.Regexp

// split obviously splits
func split(name string) []string {
	name = strings.Replace(strings.Replace(name, ".", " ", -1), "_", " ", -1)
	deserveUppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, letter := range deserveUppercase {
		name = strings.Replace(name, string(letter), " "+string(letter), -1)
	}
	res := strings.Split(name, " ")
	real := []string{}
	for _, piece := range res {
		if len(piece) > 0 {
			real = append(real, piece)
		}
	}
	return real
}

func (g *Gotify) title(value string) string {
	if val, ok := g.dict[strings.ToLower(value)]; ok {
		return val
	}
	if len(value) == 0 {
		return value
	}
	i := 0
	for i < len(value) && !('0' <= value[i] && value[i] <= '9') {
		i++
	}
	var head string
	if i == len(value) {
		head = value
		value = ""
	} else {
		head = value[:i]
		value = value[i:]
	}
	i = 0
	for i < len(value) && '0' <= value[i] && value[i] <= '9' {
		i++
	}
	var tail string
	if i == len(value) {
		tail = value
		value = ""
	} else {
		tail = value[:i]
		value = value[i:]
	}
	if val, ok := g.dict[head]; ok {
		return val + tail + g.title(value)
	} else if len(head) == 0 {
		return tail + g.title(value)
	}
	return strings.ToUpper(head[:1]) + head[1:] + tail + g.title(value)
}

// Public translates identifier into Go public identifier name
func (g *Gotify) Public(name string) string {
	splits := split(filter(name))
	res := &bytes.Buffer{}
	for _, part := range splits {
		res.WriteString(g.title(part))
	}
	return res.String()
}

// Private translates identifier into Go private identifier name
func (g *Gotify) Private(name string) string {
	splits := split(filter(name))
	res := &bytes.Buffer{}
	res.WriteString(strings.ToLower(splits[0]))
	for _, part := range splits[1:] {
		res.WriteString(g.title(part))
	}
	return res.String()
}

// Package generates acceptable package name
func (g *Gotify) Package(name string) string {
	return strings.ToLower(strings.Replace(strings.Replace(name, "_", "", -1), ".", "", -1))
}

// Goimports generates directory name that can be consumed by goformat utility
func (g *Gotify) Goimports(data string) string {
	return strings.Replace(data, "_", "-", -1)
}

// True checks if this a goish identifier
func (g *Gotify) True(name string) bool {
	return goishMask.MatchString(name)
}

func init() {
	goishMask = regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$")
}
