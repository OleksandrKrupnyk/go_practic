/**
List 1.2
*/
package main

import "fmt"

func Names() (first string, second string) {
	first = "Maxim List 1.2"
	second = "Aljona List 1.2"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
