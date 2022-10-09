package util

import "fmt"

func Print(m [][]bool) {
	for i := len(m) - 1; i >= 0; i-- {
		for _, item := range m[i] {
			if item {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
