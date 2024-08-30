package main

import "fmt"

func getData(data string) (new_data string) {
	new_data = fmt.Sprintf("This is some data %v", data)
	return new_data
}

func main() {
	fmt.Println(getData("TEST"))
}
