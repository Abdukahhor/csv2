package csv2

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {

	f, err := os.Open("hotels.csv")
	if err != nil {
		t.Error(err)
	}
	h := []hotel{}
	err = UnmarshalAll(f, &h)
	if err != nil {
		t.Error(err)
	}
	err = WriteFile(h, "json", "hotels.json")
	if err != nil {
		t.Error(err)
	}
}
