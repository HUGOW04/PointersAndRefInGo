package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
	printing := "printingPointer"

	var pointing *string = &printing
	fmt.Println("adress: ", pointing)
	fmt.Println("value: ", *pointing)

}
