package grid

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	RowTag      = "row"
	ColIndexTag = "colindex"
)

func DecodeHeaderRowData(hrd *HeaderRowData, out interface{}) error {
	return assignhrd(reflect.ValueOf(out).Elem(), hrd)
}

func DecodeRowData(rd *RowData, out interface{}) error {
	return assignrd(reflect.ValueOf(out).Elem(), rd)
}

func assignhrd(dst reflect.Value, hrd *HeaderRowData) error {
	switch dst.Kind() {

	// only support struct
	case reflect.Struct:

		t := dst.Type()

		for i := 0; i < dst.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(RowTag)
			name, nonzero := parseRowTag(tag)
			indexstring := ""
			if len(name) == 0 {
				tag = field.Tag.Get(ColIndexTag)
				indexstring, nonzero = parseRowTag(tag)
				if len(indexstring) == 0 {
					return nil
				}
			}
			if len(indexstring) > 0 {
				index, cerr := strconv.Atoi(indexstring)
				if cerr != nil {
					return cerr
				}
				if nonzero {
					index--
				}
				switch dst.Field(i).Kind() {
				case reflect.Slice:
					if arr, err := hrd.GetIValStringArray(index, ","); err == nil {
						slice := reflect.MakeSlice(dst.Field(i).Type(), len(arr), len(arr))
						for i := range arr {
							if err := set(slice.Index(i), arr[i]); err != nil {
								return err
							}
						}
						dst.Field(i).Set(slice)
					} else {
						return err
					}
				case reflect.String:
					if val, err := hrd.GetIValString(index); err == nil {
						dst.Field(i).SetString(fmt.Sprintf("%v", val))
					} else {
						return err
					}
				case reflect.Int, reflect.Int64:
					if val, err := hrd.GetIValInt(index); err == nil {
						dst.Field(i).SetInt(int64(val))
					} else {
						return err
					}
				}
			} else {
				switch dst.Field(i).Kind() {
				case reflect.Slice:
					if arr, err := hrd.GetValStringArray(name, ","); err == nil {
						slice := reflect.MakeSlice(dst.Field(i).Type(), len(arr), len(arr))
						for i := range arr {
							if err := set(slice.Index(i), arr[i]); err != nil {
								return err
							}
						}
						dst.Field(i).Set(slice)
					} else {
						return err
					}
				case reflect.String:
					if val, err := hrd.GetValString(name); err == nil {
						dst.Field(i).SetString(fmt.Sprintf("%v", val))
					} else {
						return err
					}
				case reflect.Int, reflect.Int64:
					if val, err := hrd.GetValInt(name); err == nil {
						dst.Field(i).SetInt(int64(val))
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}

func assignrd(dst reflect.Value, rd *RowData) error {
	switch dst.Kind() {

	// only support struct
	case reflect.Struct:

		t := dst.Type()

		for i := 0; i < dst.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(ColIndexTag)
			indexstring, nonzero := parseRowTag(tag)
			if len(indexstring) == 0 {
				return nil
			}
			index, cerr := strconv.Atoi(indexstring)
			if cerr != nil {
				return cerr
			}
			if nonzero {
				index--
			}
			switch dst.Field(i).Kind() {
			case reflect.Slice:
				if arr, err := rd.GetValStringArray(index, ","); err == nil {
					slice := reflect.MakeSlice(dst.Field(i).Type(), len(arr), len(arr))
					for i := range arr {
						if err := set(slice.Index(i), arr[i]); err != nil {
							return err
						}
					}
					dst.Field(i).Set(slice)
				} else {
					return err
				}
			case reflect.String:
				if val, err := rd.GetValString(index); err == nil {
					dst.Field(i).SetString(fmt.Sprintf("%v", val))
				} else {
					return err
				}
			case reflect.Int, reflect.Int64:
				if val, err := rd.GetValInt(index); err == nil {
					dst.Field(i).SetInt(int64(val))
				} else {
					return err
				}
			}

		}
	}

	return nil
}

func set(dst reflect.Value, src interface{}) error {
	switch dst.Kind() {
	case reflect.Slice:
		arr, ok := src.([]interface{})
		if !ok {
			return errors.New("expected array")
		}
		slice := reflect.MakeSlice(dst.Type(), len(arr), len(arr))
		for i := range arr {
			if err := set(slice.Index(i), arr[i]); err != nil {
				return err
			}
		}
		dst.Set(slice)

	case reflect.String:
		dst.SetString(fmt.Sprintf("%v", src))

	case reflect.Int, reflect.Int64:
		switch v := src.(type) {
		case int:
			dst.SetInt(int64(v))
		case float64:
			dst.SetInt(int64(v))
		}
	}
	return nil
}

func parseRowTag(tag string) (name string, nonzero bool) {
	if tag == "" {
		return "", false
	}
	parts := strings.Split(tag, ",")
	name = parts[0]
	if len(parts) > 1 && parts[1] == "nonzero" {
		nonzero = true
	}

	return
}
