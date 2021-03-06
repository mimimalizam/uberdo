# Journey notes

```
$ cobra init ...

Your Cobra application is ready at
$GOPATH/src/github.com/<handle>/<app-name>


$ cd $GOPATH/src/github.com/<handle>/<app-name> && go build


$ ./uberdo

A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.


$ cobra add add

add created at $GOPATH/src/github.com/mimimalizam/uberdo/cmd/add.go

```

- `cobra add` added new command add. It is loaded into, what appears to be a "root" command,
with package `cmd` in `root.go` file. Eventhough, the command added `cmd/add.go` file
it is still in the same package `cmd` (different file, same package).
- Folder <-> Go package

```

$ go build
```

# Functions in Go

It is a type, first class citizen in Go. Functions can have multiple input values.
Also, functions don't have access to anything in the calling function unless it's
passed explicitly. Each time we call a function, we push it onto the call stack,
and each time we return from a function, we pop the last function off of the stack.

# Arrays and slices

It seems that in Go we are seing more slices then arrays. An array has fixed lenght,
so adding or removing elements requires changing the length inside the brackets.
Contrary, slice's length is dynamic. It is always associated with an array.

# For loop

```go
for x,y := range Foo
```

~ `each`, `x` is index/key, `y` is value. `_` allows ignoring naming variables.

# Special functions

- `init()` - runs before `main()`.

# Misc

- `:=` - assignement operator; declares & assigns in one operation
- Capital letter means that type, method or function is available to `other` packages (example `todo/todo.go`).
Lower case means only available to `this` package.
- `go run main.go add "one two" three` <-> `go build` && `./uberdo add ...`

# Errors

```
# cmd/root.go

func Execute() {
  err := RootCmd.Execute()
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
```

There are no exceptions in Go, errors seem to be just values which are being handled
when they occur.
