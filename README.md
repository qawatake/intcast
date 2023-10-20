# intcast

`intcast` identifies integer type casts that can potentially cause overflow.

```go
func f(i int) uint {
  return uint(i) // unsafe cast
}
```

`intcast` also handles casts on defined types.

```go
type MyInt int
func f(i MyInt) uint {
  return uint(i) // unsafe cast
}
```

`intcast` ignores lines with ignore comments.

```go
func f(i int) uint {
  //lint:ignore intcast reason
  return uint(i)
}
```

## How to use

```sh
go install github.com/qawatake/intcast/cmd/intcast@latest
intcast ./...
```
