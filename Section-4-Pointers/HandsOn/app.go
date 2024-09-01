package main

import "fmt"

func main() {
	age := 32 // regular value

	agePointer := &age // Settings the pointer of the value to agePointer

	// Before using function to change pointer value
	fmt.Println("Age Pointer Address: ", agePointer)
	fmt.Println("Age Pointer Value: ", *agePointer)

	getAdultYears(agePointer) // Passing pointer to the function.

	// After using function to change pointer value
	fmt.Println("Age Pointer Address: ", agePointer)
	fmt.Println("Age Pointer Value: ", *agePointer)

}

func getAdultYears(age *int) { // We
	*age = *age - 18
}
