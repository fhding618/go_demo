package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	testDataBaseDir, _ = os.Getwd()
	databaseURL        = os.Getenv("MYSQL_URL")
)

func init() {
	fmt.Printf("path:%v\n", testDataBaseDir)
	fmt.Printf("MYSQL: %v\n", databaseURL)
}

func TestAB(t *testing.T) {

}
