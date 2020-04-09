package capture

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const tagName = "regexpGroup"

func Parse(regexpString, text string, d interface{}) error {
	if d == nil || text == "" || regexpString == "" {
		return nil
	}

	r, err := regexp.Compile(regexpString)
	if err != nil {
		return err
	}
	matches := r.FindStringSubmatch(text)
	names := r.SubexpNames()
	mapParams := mapSubexpNames(matches, names)

	if reflect.TypeOf(d).Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer to struct")
	}

	val := reflect.ValueOf(d).Elem()
	typ := reflect.TypeOf(d).Elem()
	if typ.Kind() != reflect.Struct {
		return errors.New("destination must be a struct")
	}

	for j := 0; j < typ.NumField(); j++ {
		typField := typ.Field(j)
		valField := val.Field(j)
		if !valField.IsValid() || !valField.CanSet() {
			continue
		}
		tagValue := typField.Tag.Get(tagName)

		if paramValue, ok := mapParams[tagValue]; ok {
			if !valField.IsValid() || !valField.CanSet() {
				continue
			}

			err := setValue(valField, paramValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setValue(valField reflect.Value, value string) error {
	switch valField.Kind() {
	case reflect.String:
		valField.SetString(value)

	case reflect.Int:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}

		valField.SetInt(v)

	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}

		valField.SetBool(v)

	case reflect.Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		valField.SetFloat(v)

	case reflect.Ptr:
		if value != "" {
			if valField.IsNil() {
				valField.Set(reflect.New(valField.Type().Elem()))
			}

			err := setValue(valField.Elem(), value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		r[n[i]] = m[i]
	}
	return r
}
