## Byte slices

we need a package that can help with writing to disk and reading it back

ioutil is a good one. It has a write file function

~~~go
func WriteFile(filename string, data []byte, perm os.FileMode) error
~~~

First argument is the name of the file. Second is byte slice, 3rd is permissions for the file.

### What is a byte slice

Think of a string of characters. 

"Hi there!" 

72 105 32 116 104 101 114 101 33

We typically work with strings because they're human readable. but the function from the standard library expects a byte slice. We need to essentially pass it a string in a format it expects.

We have Deck type. How do we turn that into a byte slice?

Make use of a process called `Type Conversion` - take one type and convert it to another

List out type you want, set of parentheses, and value you have. 

[]byte("Hi there!")

We have a deck, we want to turn it into a slice of bytes.

Take deck, which is a slice of strings, to a string, to a byte slice.

We need an intermediate helper function to convert a deck to a string