# Ramp up Golang

### Resources

1. [Awesome Go](https://awesome-go.com)
2. [Go by Example](https://gobyexample.com)
3. [Errors in Go](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
4. [Maps are not reference types](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
5. [Constant Errors](https://dave.cheney.net/2016/04/07/constant-errors)
6. [Laws of Reflection](https://go.dev/blog/laws-of-reflection)

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

- Go lets us get around this with the type `interface{}` which you can think of as just *any* type.

_So why not use `interface{}` for everything and have really flexible functions?_

As a user of a function that takes `interface{}` you lose type safety. What if you meant to pass `Herd.species` of type `string` into a function but instead did `Herd.count` which is an `int`? The compiler won't be able to inform you of your mistake. You also have no idea what you're allowed to pass to a function. Knowing that a function takes a `UserService` for instance is very useful.

As a writer of such a function, you have to be able to inspect anything that has been passed to you and try and figure out what the type is and what you can do with it. This is done using reflection. This can be quite clumsy and difficult to read and is generally less performant (as you have to do checks at runtime).

Use reflection if you really need to. If you want polymorphic functions, consider if you could design it around an interface (not `interface{}`, confusingly) so that users can use your function with multiple types if they implement whatever methods you need for your function to work.

- [When to use locks over channels and goroutines?](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#when-to-use-locks-over-channels-and-goroutines)

  - use channels when passing ownership of data
  - use mutexes for managing state 