# How do we run code in the project

## go build 
compile a bunch of go source code files. just compile, don't execute

~~~
computer@computer:~/gits/learn-go/helloworld$ ls
main.go  notes.md
computer@computer:~/gits/learn-go/helloworld$ go build
lcomputer@computer:~/gits/learn-go/helloworld$ ls
helloworld  main.go  notes.md
computer@computer:~/gits/learn-go/helloworld$ go build main.go
computer@computer:~/gits/learn-go/helloworld$ ls
helloworld  main  main.go  notes.md
~~~


## go run
compile and execute one or two files. 

~~~
computer@computer:~/gits/learn-go/helloworld$ ls
main.go  notes.md
computer@computer:~/gits/learn-go/helloworld$ go run main.go 
Hi there!
computer@computer:~/gits/learn-go/helloworld$ ls
main.go  notes.md
~~~


# Packages
discreet sets of code. can include any files in a package. every file in a package must declare the package they belong to

## Types of packages

### Executable
generates files we can run

### Reusable
Code used as helpers, good place to put reusable logic

### How do you know when you're making one or the other

The name of the package you use determines if you're making a dependency or executable code. `Main` is a special term, only use it when making a package you want to create some runnable file

`package main` - defines a package that can be compiled and then  executed. Must have a function called `main`


# Imports
`import "fmt"` imports a package, "fmt" into our code. This makes it accessible. It's in the standard library, but if its not imported, its not accessible to the code. 

you can import things from the standard library easily, but other packages are available too

https://golang.org/pkg is a good place to go 


# What is `func` thing

~~~
Tells go we will declare a function
|     Sets the name
|     |   List of arguments to pass
|     |   |
V     V   V
func main() {
    // Function body, call function, run this code
}
~~~


# How is the main.go file organized

~~~
package main        //package delcaration

import "fmt"        //import others we need

func main() {       // declare functions, do stuff     
    fmt.Println("hi there") 
}