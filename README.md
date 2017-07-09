# uberdo

If you'd like to use Go plugin for Vim, here is [one](https://github.com/fatih/vim-go).

# Some initial notes about my first GO app

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
