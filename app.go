package main

import (
	"log"
	"reflect"
	"strings"
)

const supportedTag = "bind"

type exampleData struct {
	Name string `bind:"src=url, field=name"`
	Age  int    `bind:"src=form, field=age"`
}

func main() {

	getTag(&exampleData{"foo", 1})

}

func getTag(s interface{}) {
	typ := reflect.TypeOf(s).Elem()
	// val := reflect.ValueOf(s).Elem()

	for i := 0; i < typ.NumField(); i++ {

		isBind := typ.Field(i).Tag.Get(supportedTag)
		if isBind != "" {

			inputFieldNameList := strings.Split(isBind, ",")

			for _, list := range inputFieldNameList {

				list = strings.Replace(list, " ", "", -1)

				log.Println(list)
			}
			// log.Println(val.Field(i))
		}
	}

}
