package main

import "fmt"

func main() {
	Backtracking(4)
}

// Рассмотрим задачу: дано число N,
// нужно сгенерировать все правильные скобочные последовательности
// из N открывающих и N закрывающих скобок.
func Backtracking(n int) {
	findSkobki(n, 0, 0, "")
}

func findSkobki(n, left, right int, s string) {
	if left == n && right == n {
		fmt.Println(s)
	} else {
		if left < n {
			findSkobki(n, left+1, right, fmt.Sprintf("%s(", s))
		}
		if right < left {
			findSkobki(n, left, right+1, fmt.Sprintf("%s)", s))
		}
	}
}
