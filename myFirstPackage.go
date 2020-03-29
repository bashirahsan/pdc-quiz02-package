
package ahsan

import "fmt"

func myPackage(choice int) int {
		cases := []string{"Pakistan has 1526 cases.","South Korea has 9583 cases.","Farance has 37575 cases.","Go Cororna GO"}
		if choice == 1 {
			fmt.Println(cases[0])
		}else if choice == 2 {
			fmt.Println(cases[1])
		}else if choice == 3 {
			fmt.Println(cases[2])
		}else if choice == 4 {
			fmt.Println(cases[3])
		}else if choice == 0 {
			fmt.Println("You can't exit without knowing cases for all countries named in the menu.")
		}else{
			fmt.Println("Option not valid.")

		}
		return 1
}