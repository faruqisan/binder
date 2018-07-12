package main

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

const supportedTag = "bind"

type bindRule struct {
	source string
	field  string
}

type exampleData struct {
	Name string `bind:"src=url, field=name"`
	Age  int    `bind:"src=form, field=age"`

	// Name string `bind:"-"`
	// Age  int    `bind:"-"`
}

func main() {

	Bind(&exampleData{})

}

func Bind(s interface{}) (interface{}, error) {
	typ := reflect.TypeOf(s).Elem()
	val := reflect.ValueOf(s).Elem()

	br := new(bindRule)	

	for i := 0; i < typ.NumField(); i++ {

		log.Println(val.Field(i))

		isBind := typ.Field(i).Tag.Get(supportedTag)

		if isBind == "" {
			return &bindRule{}, errors.New("Bind tag not found")
		}

		if isBind != "-" {
			br.setBindRule(isBind)
			log.Printf("%+v", br)
		}

	}

	return br, nil

}

func (b *bindRule) setBindRule(rules string) {

	inputFieldNameList := strings.Split(rules, ",")

	for _, list := range inputFieldNameList {
		tagKV := strings.Replace(list, " ", "", -1)
		tmp := strings.Split(tagKV, "=")

		tagK := tmp[0]
		tagV := tmp[1]

		if tagK == "src" {
			b.source = tagV
		} else if tagK == "field" {
			b.field = tagV
		}

	}

}
