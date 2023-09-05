package main

import (
	"fmt"
	"os"
)

var (
	words    []string
	fileName string
	str      string
)

func create(words []string, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return
	}
	for i := range words {
		for j := len(words) - 1; j >= 0; j-- {
			for f := len(words) - 1; f >= 0; f-- {
				for k := len(words) - 1; k >= 0; k-- {
					fmt.Fprintln(file, fmt.Sprintf("%s%s%s%s", words[i], words[j], words[f], words[k]))
				}
			}
		}
	}

}
func main() {
	for i := 1; i <= 4; i++ {
		fmt.Print("Enter key word:")
		fmt.Scanln(&str)
		words = append(words, str)
	}
	fmt.Print("Enter your password list file name:")
	fmt.Scanln(&fileName)
	create(words, fileName)
}
