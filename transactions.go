package wrap

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Transaction means all operations executed in the run function are atomic
func (c *Client) Transaction(run func(client *Client) error) error {
	session, err := c.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(c.context)

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	err = mongo.WithSession(c.context, session, func(sc mongo.SessionContext) error {
		client := &Client{
			client:  c.client,
			context: sc,
		}

		err = run(client)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(sc)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Transaction means all operations executed in the run function are atomic
func (db *Database) Transaction(run func(db *Database) error) error {
	err := db.Client.Transaction(func(client *Client) error {
		newDB := *db
		newDB.Client = client

		return run(&newDB)
	})
	if err != nil {
		return err
	}

	return nil
}

// Transaction means all operations executed in the run function are atomic
func (c *Collection) Transaction(run func(c *Collection) error) error {
	err := c.Database.Transaction(func(db *Database) error {
		newCollection := *c
		newCollection.Database = db

		return run(&newCollection)
	})
	if err != nil {
		return err
	}

	return nil
}

// Transaction means all operations executed in the run function are atomic
func (cq *CollectionQuery) Transaction(run func(cq *CollectionQuery) error) error {
	err := cq.Collection.Transaction(func(c *Collection) error {
		newCollectionQuery := *cq
		newCollectionQuery.Collection = c

		return run(&newCollectionQuery)
	})
	if err != nil {
		return err
	}

	return nil
}

// Transaction means all operations executed in the run function are atomic
func (d *Document) Transaction(run func(d *Document) error) error {
	err := d.Collection.Transaction(func(c *Collection) error {
		newDocument := *d
		newDocument.Collection = c

		return run(&newDocument)
	})
	if err != nil {
		return err
	}

	return nil
}
