package main

import "fmt"

func main() {
	S := "dasddewfreggsdgdadasdadas"
	P := "ewfregg"

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
	diff := sl - pl

	for i := 0; i < diff; i++ {
		j := 0
		for j = 0; j < pl; j++ {
			if s[i+j] != p[j] {
				break
			}
		}
		if j == pl {
			return i
		}
	}
	return -1
}
