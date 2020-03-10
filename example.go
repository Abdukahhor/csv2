package csv2

import (
	"fmt"
	"os"
	"sort"
)

//name,address,stars,contact,phone,uri
type hotelJSON struct {
	Name    string `csv:"name" valid:"stringlength(3|100)" json:"name"`
	Stars   int    `csv:"stars" valid:"range(1|5)" json:"stars"`
	URI     string `csv:"uri" valid:"url" json:"uri"`
	Address string `csv:"address" json:"address"`
	Contact string `csv:"contact" json:"contact"`
	Phone   string `csv:"phone" json:"phone"`
}

type hotelXML struct {
	Name    string `csv:"name" valid:"stringlength(3|100)" xml:"name"`
	Stars   int    `csv:"stars" valid:"range(1|5)" xml:"stars"`
	URI     string `csv:"uri" valid:"url" xml:"uri"`
	Address string `csv:"address" xml:"address"`
	Contact string `csv:"contact" xml:"contact"`
	Phone   string `csv:"phone" xml:"phone"`
}

func example() {
	f, err := os.Open("hotels.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	h := []hotelJSON{}
	//parse csv file and validates content of csv
	// only valid will be in slice
	err = UnmarshalValidate(f, &h)
	if err != nil {
		fmt.Println(err)
	}
	//sort slice by stars
	sort.Slice(h, func(p, q int) bool {
		return h[p].Stars > h[q].Stars
	})
	//write to json file
	err = WriteFile(h, "json", "hotels.json")
	if err != nil {
		fmt.Println(err)
	}
}

func exampleXML() {
	f, err := os.Open("hotels.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	h := []hotelXML{}
	//parse csv file and validates content of csv
	// only valid wil be slice
	err = UnmarshalValidate(f, &h)
	if err != nil {
		fmt.Println(err)
	}
	//sort slice by name
	sort.Slice(h, func(p, q int) bool {
		return h[p].Name < h[q].Name
	})
	//write to xml file
	err = WriteFile(h, "xml", "hotels.xml")
	if err != nil {
		fmt.Println(err)
	}
}
