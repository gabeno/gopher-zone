# Ramp up Golang

### Resources

1. [Awesome Go](https://awesome-go.com)
2. [Go by Example](https://gobyexample.com)
3. [Errors in Go](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
4. [Maps are not reference types](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
5. [Constant Errors](https://dave.cheney.net/2016/04/07/constant-errors)

### Language References

1. [Named Result Parameters](https://go.dev/wiki/CodeReviewComments#named-result-parameters)
2. [Slices](https://go.dev/blog/slices-intro)
3. [fmt options](https://pkg.go.dev/fmt)
4. [methods](https://go.dev/ref/spec#Method_declarations)
5. [Table drived tests](https://go.dev/wiki/TableDrivenTests)
6. [Method values](https://go.dev/ref/spec#Method_values)
7. [Using Go Modules](https://go.dev/blog/using-go-modules)
8. [Maps](https://go.dev/blog/maps)


### Gotchas

- This is the only way I could get `errcheck` to work on my setup (Mac M2)

```
# fix go package installs
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

- Maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. never initialize an empty map variable:

```
var m map[string]string
```

instead do:

```
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Both approaches create an empty hash map and point dictionary at it.