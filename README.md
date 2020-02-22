[![](https://godoc.org/github.com/fmmajd/gopostgres?status.svg)](https://godoc.org/github.com/fmmajd/goevent)

# Semi-ORM for postgres in GO

this package is an attempt to have the minimum requirements of an orm for postgres and go-based apps

## How to install
LATER

# Important Note
This package uses the package:[github.com/jackc/pgx](github.com/jackc/pgx) underneath

## Usage

### Initialization and Logger
Before using any functionality, you should call the InitDB function to create a connection to your database:
```go
gopostgres.InitDB("database", "username", "password", "host", "logger", "logLevel")
```

you can use LogLevels from package pgx, i.e.: `pgx.LogLevelWarn`

logger should be an implementation of pgx.Logger interface. If you want a simple logger to log everything to standard output, you can use the provided StdLogger as in below:
```go
gopostgres.InitDB("database", "username", "password", "host", gopostgres.StdLogger{}, "logLevel")
```