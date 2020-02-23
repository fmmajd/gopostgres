[![](https://godoc.org/github.com/fmmajd/gopostgres?status.svg)](https://godoc.org/github.com/fmmajd/goevent)

# Semi-ORM for Postgres in GO

this package is an attempt to have the minimum requirements of an ORM for Postgres and go-based apps

## How to install
```bash
go get -u github.com/fmmajd/gopostgres@v0.0.1
```

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
***important note***: currently, this package only works on tables with a serial primary key. support for other primary keys will be added later

This package provides an interface for Postgres objects. the interface goes like this:
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

`PostgresValue` Returns the value of each column, can return nil.

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

### Inserting a new record in the database
When you have populated a record and want to insert it into the database, you can call the function Insert and pass the record as an argument.

***important note*** the function PostgresId() MUST return 0 for the records that are not yet saved in the database, otherwise, a NewRecordWithUnZeroId error would be returned.

example:
```go
//...
//rec is previously populated
err:= gopostgres.DB.Insert(rec)
if err != nil {
  log.Fatalln(err)
} 
```

### Updating a record in the database
To update a record in the database, you can call the function Update and pass the record as an argument.

***important note*** the function PostgresId() MUST return non-zero for the records that are being updated,
 otherwise a UpdatingRecordWithZeroId error would be returned.

example:
```go
//...
//rec is previously populated
err:= gopostgres.DB.Insert(rec)
if err != nil {
  log.Fatalln(err)
} 
//..do things
rec.IncreaseViewsBy(3)
gopostgres.DB.Update(rec)
```

### Where Objects
In order to add conditions to the queries, you can use Where objects.
Currently, There are two types of where helpers, but more will be added soon.

```go
//to produce a "where column like %string%" condition
wehreLike := gopostgres.whereLike("column", "string")

//to produce a "where column = value" condition
wehreEquals := gopostgres.whereEquals("column", "value")
```

### FindAllWhere
To find ALL the rows via one or more where conditions, you can use the FindAllWhere function.

The first argument is the table name, and the second is the list of columns you want to be selected. If you want all the columns to be returned, simply put "*" here. 

example:
```go
whereTitleLikeRings := gopostgres.WhereLike("title", "rings")
whereRatingHigh := gopostgres.WhereEquals("rating", "5")
val, err := gopostgres.DB.FindAllWhere("movies", []string{"title", "rating","publish_year"}, whereTitleLikeRings, whereRatingHigh)
```
If no row is found, a NoRecordFound error would be returned.
***important note*** This package only reports AND conditions for now.

### FindBy
To find a single row via a unique column value, you can use FindBy function:
```go
val, err := gopostgres.DB.FindBy("record_table", "column", "value")
```
Remember that FindBy is used when you want ONLY one row, so if more than one row is found, a MoreThanOneRecordFound error will be returned.
And if no row is found, a NoRecordFound error would be returned. 

To fetch more than one row, you can use the function `FindAllBy`. 

### FindAllBy
To find ALL the rows via a column's value, you can use FindAllBy function:
```go
val, err := gopostgres.DB.FindAllBy("record_table", "column", "value")
```
If no row is found, a NoRecordFound error would be returned. 


### Errors
Some specific errors can be returned from query functions:

- `NoRecordFound` This error returns when there was no result for your specific condition.
- `MoreThanOneRecordFound` This error returns when the function is called to return ONE record, but more than one result is found.

## Tests

This package does not have a 100% test coverage yet. Tread with caution. 
functions without tests: FindAllWhere, Insert, Update
