package csv2

import (
	"strings"
	"testing"
)

type st struct {
	Name  string `csv:"name" json:"name"`
	Stars int    `csv:"stars" json:"stars"`
	URI   string `csv:"uri" json:"uri"`
	Ok    bool   `csv:"ok" json:"ok"`
}

type stNoHeader struct {
	Name  string `json:"name"`
	Stars int    `json:"stars"`
	Ok    bool   `json:"ok"`
	URI   string `json:"uri"`
}

func TestUnmarchalNoHeader(t *testing.T) {
	sa := []stNoHeader{}
	in := `"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	err := UnmarshalNoHeader(strings.NewReader(in), &sa)
	if err != nil {
		t.Error(err)
	}
}

func TestUnmarchal(t *testing.T) {

	sa := []st{}
	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	err := Unmarshal(strings.NewReader(in), &sa)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkUnmarchal(b *testing.B) {
	b.ResetTimer()

	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	for n := 0; n < b.N; n++ {
		sa := []st{}
		nr := strings.NewReader(in)
		err := Unmarshal(nr, &sa)
		if err != nil {
			b.Error(err)
		}
	}
}
