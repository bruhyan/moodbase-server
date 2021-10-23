## MoodBase server

built with Go

## Running the server

Run

```
go run main.go
```

### Health check

```
curl localhost:9090/health
```

### Self learning help

- `&` goes in front of a variable when you want to get that variable's memory address, i.e. gives a memory address value
- `*` goes in front of a variable that holds a memory address and resolves it (counterpart of `&`). i.e. gets the value given a memory address variable
- `*` goes in front of a type --> becomes part of type declaration --> "var str_pointer \*string: this variable holds a pointer to a string"
