# go-idnum Numeric ID struct

If you have some id or integer used on a program then you'll waste some memory and CPU for conversions to string and from string.
Instead you can create the IdNum struct that is just a pair of integer and it's string. 

## Install

    go get -u github.com/stokito/go-idnum

## Usage

```go
idNum1 := NewIdNumFromStr("42")
idNum2 := NewIdNum(42)
```

## License

[0BSD](https://opensource.org/licenses/0BSD) (similar to Public Domain)
