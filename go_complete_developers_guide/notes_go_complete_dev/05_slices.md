## Slices

Indexed from zero. Calling an item is familair too

```go
fruits := []string{"Apple", "Pear", "Banana", "Peach"}

fruits[0] = Apple
fruits [2] = Banana

// to get a range of slices

fruits[startIndexIncludign : upToNotIncluding]
fruits[0:2] = "Apple", "Pear"
fruits[:2] = "Apple", "Pear"
fruits[2:] = "Banana", "Peach"
```