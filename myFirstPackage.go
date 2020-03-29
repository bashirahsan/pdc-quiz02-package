
package pdc-quiz02-package

import "fmt"

func MyPackage(Choice int) int {
		Cases := []string{"Pakistan has 1526 cases.","South Korea has 9583 cases.","Farance has 37575 cases.","Go Cororna GO"}
		if Choice == 1 {
			fmt.Println(Cases[0])
		}else if Choice == 2 {
			fmt.Println(Cases[1])
		}else if Choice == 3 {
			fmt.Println(Cases[2])
		}else if Choice == 4 {
			fmt.Println(Cases[3])
		}else if Choice == 0 {
			fmt.Println("You can't exit without knowing cases for all countries named in the menu.")
		}else{
			fmt.Println("Option not valid.")

		}
		return 1
}
