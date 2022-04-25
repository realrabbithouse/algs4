package main

import (
	"fmt"
	"reflect"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"api-version,omitempty" protobuf:"bytes,2,opt,name=api-version"`
}

func ShowTag() {
	var t TypeMeta
	tt := reflect.TypeOf(t)
	for i := 0; i < tt.NumField(); i++ {
		fmt.Printf("Field: %v, Type: %v, Tag: %v\n", tt.Field(i).Name, tt.Field(i).Type, tt.Field(i).Tag)
	}
}

func main() {
	fmt.Println("Hello, world!")
	ShowTag()
}
