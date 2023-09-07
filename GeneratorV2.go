package main

import (
	"fmt"
	"os"
)

func create(Target []string, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return
	}
	for i := range Target {
		for j := len(Target) - 1; j >= 0; j-- {
			for k := len(Target) - 1; k >= 0; k-- {
				for l := len(Target) - 1; l >= 0; l-- {
					for m := len(Target) - 1; m >= 0; m-- {
						for n := len(Target) - 1; n >= 0; n-- {
							fmt.Fprintln(file, fmt.Sprintf("%s%s%s%s%s%s\n", Target[i], Target[j], Target[k], Target[l], Target[m], Target[n]))
						}
					}
				}
			}
		}
	}
	for i := 0; i <= len(Target)-1; i++ {
		for j := 0; j <= len(Target)-1; j++ {
			for k := 0; k <= len(Target)-1; k++ {
				for l := 0; l <= len(Target)-1; l++ {
					for m := 0; m <= len(Target)-1; m++ {
						for n := 0; n <= len(Target)-1; n++ {
							fmt.Fprintln(file, fmt.Sprintf("%s%s%s%s%s%s\n", Target[i], Target[j], Target[k], Target[l], Target[m], Target[n]))
						}
					}
				}
			}
		}
	}

}

var (
	FirstName, LastName, phoneNumber, year, month, day, fileName string
	Target                                                       []string
)

func main() {
	fmt.Print("Enter first name:")
	fmt.Scanln(&FirstName)
	Target = append(Target, FirstName)
	fmt.Print("Enter last name:")
	fmt.Scanln(&LastName)
	Target = append(Target, LastName)
	fmt.Print("Enter phone number:")
	fmt.Scanln(&phoneNumber)
	Target = append(Target, phoneNumber)
	fmt.Print("Enter year of Birth:")
	fmt.Scanln(&year)
	Target = append(Target, year)
	fmt.Print("Enter month of Birth:")
	fmt.Scanln(&month)
	Target = append(Target, month)
	fmt.Print("Enter day of Birth:")
	fmt.Scanln(&day)
	Target = append(Target, day)
	fmt.Print("Enter file name:")
	fmt.Scanln(&fileName)
	create(Target, fileName)
}
