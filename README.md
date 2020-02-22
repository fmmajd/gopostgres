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

### Creating Postgres Objects
***important note***: this package currently only works on tables with a serial primary key. support for other primary keys will be added later

This package provides an interface for postgres objects. the interface goes like this:
```go
type Record interface {
	PostgresTable() string
	PostgresColumns() []string
	PostgresValue(column string) interface{}
	PostgresValues() map[string]interface{}
	PostgresId() uint
}
```

`PostgresTable` Name of the correlating table.

`PostgresColumns` A slice of strings containing all the column names of that table.

`PostgresValue` Returns value of each column, can return nil.

`PostgresValues` Returns a map of columns for insert and update queries. Remember not to include the column "id" here.

`PostgresId` Returns the ID of the row. can return 0, if the row has not been saved yet.

example: 
```go
type User struct {
  id uint
  username string
  email string
  update_date time.Time
}

func (u User) PostgresTable() string {
  return "users"
}

func (u User) PostgresColumns() []string {
  return []string{
    "id",
    "username",
    "email",
    "update_date",
  }
}

func (u User) PostgresValue(column string) interface{} {
  switch column:
  case "id":
    return u.id
  case "username":
    return u.username
  case "email":
    return u.email
  case "update_date":
    return u.update_date
  default:
    return nil
}

func (u User) PostgresValues() map[string]interface{} {
  return map[string]interface{} {
    "username": u.PostgresValue("username"),
    "email": u.PostgresValue("email"),
    "update_date": u.PostgresValue("update_date"),
  }   
}

func (u User) PostgresId() uint {
  return u.id
}
```