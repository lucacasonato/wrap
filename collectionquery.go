package wrap

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Sorter is for sorting
type Sorter struct {
	field string
	order interface{}
}

// Ascending means sorted from small to big
func Ascending(field string) *Sorter {
	return &Sorter{field, 1}
}

// Descending means sorted from big to small
func Descending(field string) *Sorter {
	return &Sorter{field, -1}
}

// Skip skips the first n documents
func (cq *CollectionQuery) Skip(n int) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$skip": n,
	})

	return &c
}

// Limit to only return n documents
func (cq *CollectionQuery) Limit(n int) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$limit": n,
	})

	return &c
}

// Count returns the amount of documents in a single document under the field 'field'
func (cq *CollectionQuery) Count(field string) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$count": field,
	})

	return &c
}

// Sample returns n amount of documents randomly picked from the document pool
func (cq *CollectionQuery) Sample(n int) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$sample": n,
	})

	return &c
}

// Sort sorts a collection by a certain order
func (cq *CollectionQuery) Sort(sorters ...Sorter) *CollectionQuery {
	c := *cq

	finalSorters := bson.M{}

	for _, s := range sorters {
		finalSorters[s.field] = s.order
	}

	c.pipes = append(c.pipes, &bson.M{
		"$sort": finalSorters,
	})

	return &c
}

// Join gets all documents from a foreign collection in this database, filters them by
// checking if the value of foreignField matches the value of localField and if so adding
// these to the original document as an array under the 'as' key
func (cq *CollectionQuery) Join(localField string, foreignCollection string, foreignField string, as string) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$lookup": bson.M{
			"from":         foreignCollection,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as,
		},
	})

	return &c
}

// Modify changes the data structure of the field like specified by the specification
func (cq *CollectionQuery) Modify(spec map[string]interface{}) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$project": spec,
	})

	return &c
}

// AddFields adds some fields to the returned documents
func (cq *CollectionQuery) AddFields(spec map[string]interface{}) *CollectionQuery {
	c := *cq

	c.pipes = append(c.pipes, &bson.M{
		"$addFields": spec,
	})

	return &c
}

// DocumentIterator gives you an iterator to loop over the documents
func (cq *CollectionQuery) DocumentIterator() (*Iterator, error) {
	cursor, err := cq.Collection.collection.Aggregate(cq.Collection.Database.Client.ctx(), cq.pipes)
	if err != nil {
		return nil, err
	}

	return &Iterator{
		Collection: cq.Collection,
		cursor:     cursor,
	}, nil
}
