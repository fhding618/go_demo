package main

import (
	"fmt"
	"net/url"
	"os"
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

func main() {
	a := []int{}

	store := &Store{dbName: "ab"}
	if store.config != nil {
		fmt.Printf("ab")
	}
	fmt.Printf("End")

	var testDataBaseDir, _ = os.Getwd()
	fmt.Printf("%v", testDataBaseDir)
}
