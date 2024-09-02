# Interfaces & Generics

## Interfaces
Interfaces in Go work in a different way that in other langauges.
A function that accepts a type of an interface, like the method "saveData" expects the parameter that is sent there, to have "Save()" function / method in it.

Also, you can embed interfaces inside interfaces, for example "outputable" contains the "saver" functionality.

```
type saver interface {
	Save() error
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

type outputable interface {
	saver
	Display()
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}
```

- Interfaces are also used as "Any", the function below can recieve any value.
```
saveData(value interface{}){} 
```

## Generics
Generics can be used like this, pretty self explanatory.

```
package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
```