# Arrays, Slices & Maps

## `Arrays`
- Array init examples  
    You can init an empty array ofcourse, but it's important to specify the size.
    ```
    package main

    import "fmt"

    func main() {

        var products [4]string
        prices := [5]float64{10, 20, 30, 40, 44}

        fmt.Println(prices)
        fmt.Println(products)
    }
    ```
- Arrays are accesible by index
    ```
    package main

    import "fmt"

    func main() {
        var products [4]string
        prices := [5]float64{10, 20, 30, 40, 44}
        fmt.Println(prices)
        fmt.Println(products)

        products[0] = "potato"
        products[3] = "banana"

        fmt.Println(products)
        
    }

    ---
    [10 20 30 40 44]
    [   ]
    [potato   banana]
    ---
    ```

## `Slices`
- Using slices you can get values from an array for example, note that the FIRST value is always included, and the LAST value is always excluded.

    ```
    func main() {
        prices := [5]float64{10, 20, 30, 40, 44}
        featuredPrices := prices[1:3] // First value included, second value EXCLUDED
        fmt.Println(featuredPrices)
    }
    ---
    [20 30]
    ---
    ```
- Ofcourse you can also cut in these ways:
    ```
    featuredPrices := prices[:3] // Will take from the begining untill 3
    featuredPrices := prices[1:] // Will take from the 1 until the end.
    ```

### * **Slices are actually kind of a pointer to the particular part of an array.** *

## `Dynamic Lists / Slices`
`
- This is how you can declear an  slice with no size
    ```
    newPrices := []float64{10.90, 20.99, 30.99}
	fmt.Println(newPrices)
    ```
    now while the size is not declared, you still can't access indexes that do no exist:
    ```
    newPrices := []float64{10.90, 20.99, 30.99}
	fmt.Println(newPrices)

	newPrices[1] = 9
	//newPrices[3] = 20 // This will not work, as this index does not exist.
    ```
- This is how you "add" something to a slice, to be more specific, **it returns a new slice**, append also can take N amount of items.
    ```
    newPrices = append(newPrices, 9)
    newPrices = append(newPrices, 12,13,114,100,200,300)

	fmt.Println(newPrices)
    ```
- You can also add a whole slice to a slice, using this syntax **slice...**, the 3 dots will unpack everything from the slice.
    ```
    func main() {
        prices := [5]float64{10, 20, 30, 40, 44}
        pricesToMerge := []float64{100, 200, 300, 400}
        newPrices = append(newPrices, pricesToMerge...)
        fmt.Println(newPrices)
    }
    ```

## `Maps`
- Maps are pretty straight forward.  
when we specify a map, we need to declare which types are going in there, in this case **string : string** -> (key:value)  
Note that everything can be a key.
    ```
    func main() {
        // Creating a map.
        websites := map[string]string{
            "Google":              "google.com",
            "Amazon Web Services": "aws.com",
        }
        fmt.Println(websites)

        // Adding to a map
        websites["LinkedIn"] = "linkedin.com"
        fmt.Println(websites)

        // Removing from a Map
        delete(websites, "Google")
        fmt.Println(websites)
    }
    ```

## `make`
There is the `make` function, that provides you the opportunity to create arrays / slices with already delegated memmory amount.


```
    func main() {
        courseRatings := make(map[string]float64, 3)
        fmt.Println(courseRatings)
    }
```
You also can use type alias for this, and make it better from a developer standpoint.
```
    package main

    import "fmt"

    type floatMap map[string]float64

    func (m floatMap) output() {
        fmt.Println(m)
    }

    func main() {
        courseRatings := make(floatMap, 3)
        courseRatings.output()
    }

```

## `Iteration`
You can loop over Maps / Slices / Arrays in this way:
```
	users := make([]string, 2, 5)
	users[0] = "Tom"
	users = append(users, "Josh")
	users = append(users, "Tim")

	for index, value := range users {
		fmt.Println(index, " ", value)
	}
```