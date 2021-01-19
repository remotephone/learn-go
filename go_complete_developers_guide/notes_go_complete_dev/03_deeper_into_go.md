## Intro to our program

This will be a program that simulates a deck of playing cards. 

### Functions

newDeck - create an array of strings
print - log out the contents of a deck of cards
shuffle - shuffle deck of cards
deal - create a hand of cards
saveToFile - save list of cards to a file
newDeckFromFile - load list of cards

### Declaring Variables

Create a new variable and assign it a string that represents a playing card.

Use fmt package to print the string. 

In the example, we did not add the `import "fmt"` string. VsCode did that for us, which is nice, but be aware of it. 

~~~
var card string = "Ace of Spades"
var = variable, tells go we got one coming
card = name of variable, it is "card"
string = tell's go compiler that only value of type string will ever be assigned here
"Ace of Spades" = value of variable card
~~~

### Type Declaration
Go is a `statically typed` language. Dnyamically typed languages don't care what value is assigned to a variable. staically typed languages do

These basic types include, but aren't limited to:

string - sequence of characters
bool - true or false
int - number, non decimal
float64 - number with decimal


These two statements are the same! In the second case, we rely on go compiler to figure out it contains a string
~~~
    var card string = "Ace of Spades"
	card := "Ace of Spades"
~~~

We only use `:=` syntax on NEW variables, we do not do that on redefining it. You only do this when you `inititalize` the variable. 


### Keep in Mind

- You can initialize a variable in a function and assign a value to it at the same time. 
- If you inititalize a variable OUTSIDE a function,  yuo cannot assign a value to it. You can do that inside the function instead. 

```go
//Works:
package main
 
import "fmt"
 
var deckSize int
 
func main() {
  deckSize = 50
  fmt.Println(deckSize)
}

//Does not Work:
package main
 
import "fmt"
 
deckSize := 50
 
func main() {
  fmt.Println(deckSize)
}
```

### Return Values

Go expects you to tell the compiler what type of value will be returned when a function runs, if it returns something

```go
// This gives an error
func newCard() {
	return "Five of Diamonds"
}

// This works becase we told the compiler what kind of value to expect
func newCard() string {
	return "Five of Diamonds"
}

```

### Arrays in Go

Array = Fixed length list of things
  - primitive data structure for holding list of records
  - fixed length
  - must be defined with data type, all identical types
Slice = An array that can shrink or grow
  - more "array" like features
  - no fixed length
  - must be of same type
  

