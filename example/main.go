package main

import (
	"fmt"
	"time"

	"github.com/lucacasonato/wrap/expressions"
	"github.com/lucacasonato/wrap/filter"

	"github.com/lucacasonato/wrap/update"

	"github.com/lucacasonato/wrap"
)

// User record
type User struct {
	Name       string
	Email      string
	Numbers    []int
	LastEdited time.Time
}

func main() {
	client, err := wrap.Connect("mongodb://localhost:27017", 5*time.Second)
	if err != nil {
		panic(err)
	}

	db := client.Database("production")
	users := db.Collection("users")

	err = users.Delete()
	if err != nil {
		panic(err)
	}

	err = users.CreateIndex(map[string]wrap.Index{
		"name":  wrap.TextIndex,
		"email": wrap.AscendingIndex,
	})
	if err != nil {
		panic(err)
	}

	luca, err := users.Add(&User{
		Name:       "Luca Casonato",
		Email:      "luca.casonato@antipy.com",
		Numbers:    []int{5, 10, 15},
		LastEdited: time.Now(),
	})
	if err != nil {
		panic(err)
	}

	_, err = users.Add(&User{
		Name:       "Jaap Aarts",
		Email:      "jaap.aarts@antipy.com",
		Numbers:    []int{20, 4, 100},
		LastEdited: time.Now(),
	})
	if err != nil {
		panic(err)
	}

	<-time.NewTimer(10 * time.Millisecond).C

	err = luca.Update(update.CurrentDate("lastedited", update.Date), true)
	if err != nil {
		panic(err)
	}

	iterator, err := users.Where(
		filter.AND(
			filter.TextSearch("luca"),
			filter.Equal("email", "luca.casonato@antipy.com"),
		),
	).Modify(map[string]interface{}{
		"average_number": expressions.Avg(expressions.Value("numbers")),
	}).DocumentIterator()
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
}
