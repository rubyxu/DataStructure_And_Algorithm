//朴素/暴力模式匹配
package main

import "fmt"

func main() {
	S := "dasddewfreggsdgdadasdadas"
	P := "dgda"

	local := match(S, P)
	if local == -1 {
		fmt.Println("no match.")
		return
	}
	fmt.Println("matched, local is:", local)
}

func match(s, p string) int {
	sl := len(s)
	pl := len(p)
	fmt.Println(sl, pl)
	var i, j int = 0, 0
	for {
		if i < sl && j < pl {
			if s[i] == p[j] {
				i++
				j++
			} else {
				i = i - (j - 1)
				j = 0
			}
		} else {
			break
		}
	}
	if j == pl {
		return i - j
	} else {
		return -1
	}
}
