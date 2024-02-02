# Learning Go

## Generating Testing Docs


You can add examples in the *_test.go files and they will show up in the godoc.
It will take the comment that is defined above to function and make it part of the docs e.g.

```
// Add(x, y) takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
```

Provide an // Output
to also show an output in the godoc example.

adder_test.go
```
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

```

Seemingly `go test -v` will only run tests that also have an //Output defined
```
$ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
```