package csv2

import (
	"encoding/csv"
	"errors"
	"io"
	"reflect"

	"github.com/asaskevich/govalidator"
)

//UnmarshalValidate parses the CSV from the reader in the interface.
//Read and process line by line
//Validates by valid tag in struct, if validation is true then adds to a slice
func UnmarshalValidate(in io.Reader, out interface{}) error {

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
	//read first line for header
	record, err := r.Read()
	if err != nil {
		return err
	}
	//parse header
	htags := headerTag(record, out)
	//read content of CSV
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
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
		result, _ := govalidator.ValidateStruct(el.Interface())
		if result {
			//appends the struct to slice
			rv.Set(reflect.Append(rv, el))
		}
	}
	return nil
}
