# Bytes

Format bytes to string

## Installation

```go get github.com/labstack/gommon/bytes```

## [Usage](https://github.com/labstack/gommon/blob/master/bytes/bytes_test.go)

### Decimal prefix

```go
fmt.Println(bytes.Format(1323))
```

`1.32 KB`

### Binary prefix

```go
fmt.Println(bytes.FormatBin(1323))
```

`1.29 KiB`
