package csv2

import (
	"strings"
	"testing"
)

type stv struct {
	Name  string `csv:"name" valid:"stringlength(3|100)" json:"name"`
	Stars int    `csv:"stars" valid:"range(1|5)" json:"stars"`
	URI   string `csv:"uri" valid:"url" json:"uri"`
	Ok    bool   `csv:"ok" json:"ok"`
}

func TestUnmarchalValidate(t *testing.T) {

	sa := []stv{}
	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,0,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	err := UnmarshalValidate(strings.NewReader(in), &sa)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkUnmarchalValidate(b *testing.B) {
	b.ResetTimer()

	in := `name,stars,ok,uri
"Rob",5,true,http://alif.tj/rob
Ken,3,false,https://github.com/dineshs
"Robert",4,true,"https://golang.org/pkg/reflect/"
`
	for n := 0; n < b.N; n++ {
		sa := []st{}
		nr := strings.NewReader(in)
		err := UnmarshalValidate(nr, &sa)
		if err != nil {
			b.Error(err)
		}
	}
}
