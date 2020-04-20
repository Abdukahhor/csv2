package csv2

import (
	"encoding/csv"
	"errors"
	"io"
	"reflect"
	"strconv"
)

//parse header for position of CSV column with tag of struct
func headerTag(rec []string, out interface{}) (htags []int) {
	el := reflect.TypeOf(out)
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	}
	el = el.Elem()
	for j := 0; j < len(rec); j++ {
		for i := 0; i < el.NumField(); i++ {
			f := el.Field(i)
			if tg := f.Tag.Get("csv"); tg == rec[j] {
				htags = append(htags, i)
				break
			}
		}
	}
	return
}

//UnmarshalNoHeader parses the CSV from the reader in the interface without csv header.
func UnmarshalNoHeader(in io.Reader, out interface{}) error {
	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr && rv.IsNil() {
		return errors.New("nil or is not pointer " + reflect.TypeOf(out).String())
	}
	rv = rv.Elem()
	rvType := rv.Type()
	if rvType.Kind() != reflect.Slice {
		return errors.New("cannot use " + rv.String() + ", only slice or array supported")
	}
	innerType := rvType.Elem()
	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	csvLen := len(records)
	//make slice to hold the CSV content
	rv.Set(reflect.MakeSlice(rvType, csvLen, csvLen))
	for i, record := range records {
		el := reflect.New(innerType).Elem()
		for k, v := range record {
			f := el.Field(k)
			err = setField(f, v)
			if err != nil {
				return err
			}
		}
		//set struct to slice
		rv.Index(i).Set(el)
	}
	return nil
}

//Unmarshal parses the CSV from the reader in the interface.
//Read all from io.Reader and after that proccess
func Unmarshal(in io.Reader, out interface{}) error {

	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr && rv.IsNil() {
		return errors.New("nil or is not pointer " + reflect.TypeOf(out).String())
	}
	rv = rv.Elem()
	rvType := rv.Type()
	if rvType.Kind() != reflect.Slice {
		return errors.New("cannot use " + rv.String() + ", only slice or array supported")
	}
	innerType := rvType.Elem()

	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("empty csv")
	}
	//parse header for position of CSV column
	htags := headerTag(records[0], out)
	//parse body of csv
	body := records[1:]
	csvLen := len(body)
	//make slice to hold the CSV content
	rv.Set(reflect.MakeSlice(rvType, csvLen, csvLen))
	for i, record := range body {
		//new reflect of struct
		el := reflect.New(innerType).Elem()
		for k, v := range record {
			// Position found accordingly to header
			f := el.Field(htags[k])
			err = setField(f, v)
			if err != nil {
				return err
			}
		}
		//set struct to slice
		rv.Index(i).Set(el)
	}
	return nil
}

func setField(f reflect.Value, v string) error {
	if !f.IsValid() {
		return errors.New("invalid value")
	}
	//set value by type
	switch f.Kind() {
	case reflect.String:
		f.SetString(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		f.SetInt(i64)
	case reflect.Float32, reflect.Float64:
		fl, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		f.SetFloat(fl)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		f.SetBool(b)
	}
	return nil
}
