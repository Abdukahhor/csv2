package csv2

import (
	"os"
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

//"first_name","last_name","company_name","address","city","county","state","zip","phone1","phone2","email","web"
type hotel struct {
	Name    string `csv:"first_name" json:"first_name"`
	Stars   int    `csv:"stars" json:"stars"`
	URI     string `csv:"uri" json:"uri"`
	Address string `csv:"address" json:"address"`
	Contact string `csv:"contact" json:"contact"`
	Phone   string `csv:"phone" json:"phone"`
}

func TestUnmarchalFile(t *testing.T) {

	f, err := os.Open("hotels.csv")
	if err != nil {
		t.Error(err)
	}
	h := []hotel{}
	err = UnmarshalAll(f, &h)
	if err != nil {
		t.Error(err)
	}
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

func TestUnmarchalAll(t *testing.T) {

	sa := []st{}
	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	err := UnmarshalAll(strings.NewReader(in), &sa)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkUnmarchalAll(b *testing.B) {
	b.ResetTimer()

	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	for n := 0; n < b.N; n++ {
		sa := []st{}
		nr := strings.NewReader(in)
		err := UnmarshalAll(nr, &sa)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnmarchalAllFile(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		f, err := os.Open("hotels.csv")
		if err != nil {
			b.Error(err)
		}
		h := []hotel{}
		err = UnmarshalAll(f, &h)
		if err != nil {
			b.Error(err)
		}
	}
}
