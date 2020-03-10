package csv2

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

//Write marshal JSON, XML
func Write(v interface{}, format string) (b []byte, err error) {
	switch format {
	case "json":
		b, err = json.Marshal(v)
		if err != nil {
			return
		}
	case "xml":
		b, err = xml.Marshal(v)
		if err != nil {
			return
		}
	}

	return
}

//WriteFile to JSON, XML file
func WriteFile(v interface{}, format, filename string) error {
	b, err := Write(v, format)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
