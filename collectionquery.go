package wrap

import "context"

// DocumentIterator gives you an iterator to loop over the documents
func (c *CollectionQuery) DocumentIterator() (*Iterator, error) {
	cursor, err := c.collection.collection.Find(context.Background(), c.filter)
	if err != nil {
		return nil, err
	}

	return &Iterator{
		Collection: c.collection,
		cursor:     cursor,
	}, nil
}
