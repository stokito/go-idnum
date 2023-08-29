# go-idnum Numeric ID struct

If you have some id or integer used on a program then you'll waste some memory and CPU for conversions to string and from string.
Instead you can create the IdNum struct that is just a pair of integer and it's string.
Then you can call multiple times it's String() method without any new allocations.
The type supports JSON serialization and deserialization

## Usage

```go
idNum1 := NewIdNum(42)
idNum2 := NewIdNumFromStr("42")
idNum3 := NewIdNumFromBytes([]byte("42"))

idNum1.Num == 42
idNum1.Str == "42"
idNum1.String() == "42"

// You can use as a field type in DTO
type User struct {
    Id IdNum
}

body, err := json.Marshal(u)
body == `{"Id":42}`
```

## Install

    go get -u github.com/stokito/go-idnum


## License

[0BSD](https://opensource.org/licenses/0BSD) (similar to Public Domain)
