package capture

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

const tagName = "regexpGroup"

func Parse(regexpString, text string, d interface{}) error {
	if d == nil || len(text) == 0 || len(regexpString) == 0 {
		return nil
	}

	r, err := regexp.Compile(regexpString)
	if err != nil {
		return err
	}
	matches := r.FindStringSubmatch(text)
	names := r.SubexpNames()
	mapParams := mapSubexpNames(matches, names)

	typ := reflect.TypeOf(d).Elem()
	val := reflect.ValueOf(d).Elem()

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

			switch typField.Type.Kind() {
			case reflect.String:
				valField.SetString(paramValue)

			case reflect.Int:
				v, err := strconv.ParseInt(paramValue, 10, 64)
				if err != nil {
					return err
				}

				valField.SetInt(v)

			case reflect.Bool:
				v, err := strconv.ParseBool(paramValue)
				if err != nil {
					return err
				}

				valField.SetBool(v)

			case reflect.Float64:
				v, err := strconv.ParseFloat(paramValue, 64)
				if err != nil {
					return err
				}

				valField.SetFloat(v)
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
