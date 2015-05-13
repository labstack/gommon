# Bytes

Format bytes to string

## Installation

```go
go get github.com/labstack/gommon/bytes
```

## [Usage](https://github.com/labstack/gommon/blob/master/bytes/bytes_test.go)

```sh
import gytes github.com/labstack/gommon/bytes
```

### Decimal prefix

```go
fmt.Println(gytes.Format(1323))
```

`1.32 KB`

### Binary prefix

```go
fmt.Println(gytes.FormatB(1323))
```

`1.29 KiB`
