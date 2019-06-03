package wrap

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

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

// DocumentIterator gives you an iterator to loop over the documents
func (cq *CollectionQuery) DocumentIterator() (*Iterator, error) {
	cursor, err := cq.collection.collection.Aggregate(context.Background(), cq.pipes)
	if err != nil {
		return nil, err
	}

	return &Iterator{
		Collection: cq.collection,
		cursor:     cursor,
	}, nil
}
