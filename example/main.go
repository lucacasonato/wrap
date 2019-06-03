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
	Name            string
	Email           string
	FavoriteNumbers []int
	LastEdited      time.Time
}

func main() {
	client, err := wrap.Connect("mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs", 5*time.Second)
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
		Name:            "Luca Casonato",
		Email:           "luca.casonato@antipy.com",
		FavoriteNumbers: []int{5, 10, 15},
		LastEdited:      time.Now(),
	})
	if err != nil {
		panic(err)
	}

	jaap, err := users.Add(&User{
		Name:            "Jaap Aarts",
		Email:           "jaap.aarts@antipy.com",
		FavoriteNumbers: []int{20, 4, 100},
		LastEdited:      time.Now(),
	})
	if err != nil {
		panic(err)
	}

	<-time.NewTimer(10 * time.Millisecond).C

	err = users.Transaction(func(users *wrap.Collection) error {
		now := time.Now()

		err := users.Document(luca.ID).Update(true, update.Set("lastedited", now), update.Set("email", "luca@antipy.com"))
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

	if err != nil {
		panic(err)
	}

	iterator, err := users.
		Where(
			filter.TextSearch("luca"),
		).
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
}
