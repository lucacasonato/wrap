package wrap

// Next means go to the next document in the iterator
func (i *Iterator) Next() bool {
	return i.cursor.Next(i.Collection.Database.Client.ctx())
}

// Data decodes some data and returns an interface
func (i *Iterator) Data() (interface{}, error) {
	var data interface{}
	err := i.cursor.Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DataTo decodes some data into an interface
func (i *Iterator) DataTo(data interface{}) error {
	err := i.cursor.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// Close stops the iterator
func (i *Iterator) Close() error {
	err := i.cursor.Close(i.Collection.Database.Client.ctx())
	if err != nil {
		return err
	}

	return nil
}
