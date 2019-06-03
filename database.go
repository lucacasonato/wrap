package wrap

// Database gets a database instance from a client
func (c *Client) Database(id string) *Database {
	database := c.client.Database(id)

	return &Database{ID: id, database: database, Client: c}
}

// Delete a database
func (d *Database) Delete() error {
	return d.database.Drop(d.Client.context)
}
