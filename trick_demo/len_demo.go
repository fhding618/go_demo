/**
 * https://mp.weixin.qq.com/s/GqdzjXuaZOmjqit32Vxuzg
 **/

package main

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << uint(len(s)) / 128
var b byte = 1 << uint(len(s[:])) / 128

func main() {
	println(a, b)
}
