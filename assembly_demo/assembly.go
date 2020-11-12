/**
 * 生成汇编
 * GOOS=linux GOARCH=386 go tool compile -S assembly.go >> assembly.S
 **/

package main

func g(p int) int {
	return p + 1
}

func main() {
	c := g(4) + 1
	_ = c
}
