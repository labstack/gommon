# Gytes

Format bytes to string

## Installation

```go
go get github.com/labstack/gommon/gytes
```

## [Usage](https://github.com/labstack/gommon/blob/master/gytes/gytes_test.go)

```sh
import github.com/labstack/gommon/gytes
```

### Decimal prefix

```go
fmt.Println(gytes.Format(1323))
```

`1.32 KB`

### Binary prefix

```go
gytes.BinaryPrefix(true)
fmt.Println(gytes.Format(1323))
```

`1.29 KiB`

Above examples operate on global instance of Gytes. To create a new instance

```go

