package util

import "fmt"

func Print(m [][]bool) {
	for _, row := range m {
		for _, item := range row {
			if item {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
