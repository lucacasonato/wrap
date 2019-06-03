package update

import "go.mongodb.org/mongo-driver/bson"

// Update fields
type Update *bson.M

// CurrentDateType for the CurrentDate update operator
type CurrentDateType string

const (
	// Date as type for update operator
	Date CurrentDateType = "date"
	// Timestamp as type for update operator
	Timestamp CurrentDateType = "timestamp"
)

// CurrentDate sets a certain field to the current date or timestamp
func CurrentDate(field string, typ CurrentDateType) Update {
	return Update(&bson.M{
		"$currentDate": bson.M{
			field: bson.M{
				"$type": typ,
			},
		},
	})
}

// Increment a field by a certain amount
func Increment(field string, amount float64) Update {
	return Update(&bson.M{
		"$inc": bson.M{
			field: amount,
		},
	})
}

// SetIfLess operator updates only when current stored value is less than the new value
func SetIfLess(field string, value interface{}) Update {
	return Update(&bson.M{
		"$min": bson.M{
			field: value,
		},
	})
}

// SetIfGreater operator updates only when current stored value is greater than the new value
func SetIfGreater(field string, value interface{}) Update {
	return Update(&bson.M{
		"$max": bson.M{
			field: value,
		},
	})
}

// Multiply a field by a certain amount
func Multiply(field string, amount float64) Update {
	return Update(&bson.M{
		"$mul": bson.M{
			field: amount,
		},
	})
}

// Rename a field to a new name
func Rename(field string, newName string) Update {
	return Update(&bson.M{
		"$rename": bson.M{
			field: newName,
		},
	})
}

// Set a field to a new value
func Set(field string, value interface{}) Update {
	return Update(&bson.M{
		"$set": bson.M{
			field: value,
		},
	})
}

// SetIfNew a field to a new value if the document gets created by this update call
func SetIfNew(field string, value interface{}) Update {
	return Update(&bson.M{
		"$setOnInsert": bson.M{
			field: value,
		},
	})
}

// Unset removes a field from the document
func Unset(field string) Update {
	return Update(&bson.M{
		"$unset": bson.M{
			field: "",
		},
	})
}

// AddToSet adds the value to the array if it isnt already in the array
func AddToSet(field string, value interface{}) Update {
	return Update(&bson.M{
		"$addToSet": bson.M{
			field: value,
		},
	})
}

// PopFirst removes the first element from the array
func PopFirst(field string) Update {
	return Update(&bson.M{
		"$pop": bson.M{
			field: -1,
		},
	})
}

// PopLast removes the last element from the array
func PopLast(field string) Update {
	return Update(&bson.M{
		"$pop": bson.M{
			field: 1,
		},
	})
}

// RemoveAll element(s) from an array that are in remove array
func RemoveAll(field string, array []interface{}) Update {
	return Update(&bson.M{
		"$pullAll": bson.M{
			field: array,
		},
	})
}

// Push all elements onto the array
func Push(field string, newItems []interface{}) Update {
	return Update(&bson.M{
		"$push": bson.M{
			"$each": bson.M{
				field: newItems,
			},
		},
	})
}

// BitAND does a bitwise AND between the selected field and the integer
func BitAND(field string, value int) Update {
	return Update(&bson.M{
		"$bit": bson.M{
			field: bson.M{
				"and": value,
			},
		},
	})
}

// BitOR does a bitwise or between the selected field and the integer
func BitOR(field string, value int) Update {
	return Update(&bson.M{
		"$bit": bson.M{
			field: bson.M{
				"or": value,
			},
		},
	})
}

// BitXOR does a bitwise XOR between the selected field and the integer
func BitXOR(field string, value int) Update {
	return Update(&bson.M{
		"$bit": bson.M{
			field: bson.M{
				"xor": value,
			},
		},
	})
}
