# wrap

[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/lucacasonato/wrap)

Wrap enables the full MongoDB feature set without ever having to touch BSON or other weird libraries. Everything is structured in a way that BSON and MongoDB stay completly hidden.

## install

```bash
go get github.com/lucacasonato/wrap
```

```go
import "github.com/lucacasonato/wrap"
```

## usage

#### connect

```go
client, err := wrap.Connect("mongodb://localhost:27017", 5*time.Second)
if err != nil {
  panic(err)
}
```

#### open a database

```go
db := client.Database("production")
```

> note: you are only getting a refrence to the database here. you are not actually creating it yet

#### get a collection

```go
users := db.Collection("users")
```

> note: you are only getting a refrence to the collection here. you are not actually creating it yet

#### add data

```go
doc, err := users.Add(&User{
  Name:            "Luca Casonato",
  Email:           "luca.casonato@antipy.com",
  FavoriteNumbers: []int{5, 10, 15},
  LastEdited:      time.Now(),
})
if err != nil {
  panic(err)
}
```

#### get data

```go
data, err := doc.Get()
if err != nil {
  panic(err)
}

user := User{}

data.DataTo(&user)

fmt.Println(user.Name)
```

#### update data

```go
err = doc.Update(update.Set("email", "luca@antipy.com"), true)
if err != nil {
  panic(err)
}
```

#### create index

```go
err = users.CreateIndex(map[string]wrap.Index{
  "name":  wrap.TextIndex,
  "email": wrap.AscendingIndex,
})
if err != nil {
  panic(err)
}
```

#### get filtered data

```go
iterator, err := users.
  Where(filter.AND(
    filter.TextSearch("luca"),
    filter.Equal("email", "luca.casonato@antipy.com"),
  )).
  DocumentIterator()
if err != nil {
  panic(err)
}
defer iterator.Close()

for iterator.Next() {
  user := User{}
  err := iterator.DataTo(&user)
  if err != nil {
    panic(err)
  }

  fmt.Println(user)
}
```

#### get structurally modified data (aggregation)

```go
iterator, err = users.All().
  Modify(map[string]interface{}{
    "email": expressions.Exclude,
  }).
  AddFields(map[string]interface{}{
    "averagefavoritenumber": expressions.MathAvg(expressions.Value("favoritenumbers")),
  }).
  DocumentIterator()

if err != nil {
  panic(err)
}
defer iterator.Close()

for iterator.Next() {
  user := map[string]interface{}{}
  err := iterator.DataTo(&user)
  if err != nil {
    panic(err)
  }

  fmt.Println(user)
}
```

#### transactions

```go
err := users.Transaction(func(users *wrap.Collection) error {
  now := time.Now()

  err := users.Document(luca.ID).Update(true, update.Set("lastedited", now))
  if err != nil {
    return err
  }

  err = users.Document(jaap.ID).Update(true, update.Set("lastedited", now))
  if err != nil {
    return err
  }

  return nil
})
if err != nil {
  err = nil
}
```

#### example

A full example can be found in the "example" folder.

## planning

- implement schema filters (im lazy)
- automatic index creation
- more tests

## contributing

to build start a mongo server on localhost:27017

run tests with `go test`

1. open an issue about your idea / suggestion / bug
2. wait for response
3. have someone fix it or pr yourself
4. thanks

## licence

Copyright (c) 2019 Luca Casonato

This project is licenced under the MIT licence. More details in the LICENCE file.
