package main

import (
	"fmt"
	"net/url"
	"reflect"
)

type Options struct {
	Name       string
	MasterAddr *url.URL
}

type ZhihuMySQLConfig struct {
	Name string
	Opt  *Options
}

type Store struct {
	dbName string
	config *ZhihuMySQLConfig
}

func AB() string {
	return "ab"
}

var s = "Go101.org"

var a uint32 = 1 << uint32(len(s)) / 128
var b uint32 = 1 << uint32(len(s[:])) / 128

func main() {
	ss := 'a'
	fmt.Printf("%v\n", ss+1)

	aa := len(s)
	fmt.Printf("%#v\n", reflect.TypeOf(aa))
	fmt.Printf("len: %d val: %d\n", uint32(len(s[:])), b)
	fmt.Printf("len: %d val: %d\n", uint32(len(s)), a)

	//var s = "ab"
	//fmt.Printf("%d\n", len(s))
	//
	//store := &Store{dbName: "ab"}
	//if store.config != nil {
	//	fmt.Printf("ab")
	//}
	//fmt.Printf("End")
	//
	//var testDataBaseDir, _ = os.Getwd()
	//fmt.Printf("%v", testDataBaseDir)
}
