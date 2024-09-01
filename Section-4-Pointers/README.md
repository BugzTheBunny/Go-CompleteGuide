# Pointers.

### _Why Pointers?_

- #### Avoiding Unnecessary value copies.
    By default, Go creates a copy when passing a value to a function  
    meaning you will have the save value two times in the memmory, until the end of the function run, when the garbage collector will remove it.  
    Which means, that for large and complex values, this can take a lot of space for no reason.

     When using pointers, you can just pass the address of the value, so it will access it directly.
    ```
    ** IMPORTANT ** - You should not be avoiding passing small values, like Integers for example to functions, as these are really small values that bearly take any memmory, therefore the extra code and operations are not worth it.
    ```
- #### Directly Mutate Values.
     If we are passing the address, instead of the value, we can change directly the value, so we don't even need return statements.  
     This leads to less code, but can make some code less understandable, or unexpected behaviour.