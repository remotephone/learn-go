# Struct

## Data structure
Define a structure for Data. Think of cards app

suit -> string
value -> string

kinda like a python dictionary, but not really. don't get carried away with the metaphor

There are a few ways to define a struct

```go
	// Kinda wacky, relies on positional arguments, if you swap anything you break stuff
	alex := person{"Alex", "Anderson"}

	// more reasonable
	tim := person{firstName: "Tim", lastName: "Tomas"}

	// assigns zero value to each field
	var steve person
```
```go
$ go run main.go                     [22:21:05]
{Alex Anderson}
{Tim Tomas}
{ }
```

