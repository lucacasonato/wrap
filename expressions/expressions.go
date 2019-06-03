package expressions

import (
	"github.com/lucacasonato/wrap/types"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	// Include a certain field
	Include interface{} = true
	// Exclude a certain field
	Exclude interface{} = true
)

// Value is the value of a field in the document
func Value(field string) interface{} {
	return "$" + field
}

// Literal makes arbitrary data an expression without parsing
func Literal(data interface{}) interface{} {
	return data
}

// Abs is the absolute value of a value
func Abs(n interface{}) interface{} {
	return &bson.M{
		"$abs": n,
	}
}

// Add adds all values together
func Add(expressions ...interface{}) interface{} {
	return &bson.M{
		"$add": expressions,
	}
}

// IsAllTrue returns true if all elements in the array are not 'false, null, 0 or undefined'
func IsAllTrue(array interface{}) interface{} {
	return &bson.M{
		"$allElementsTrue": bson.A{
			array,
		},
	}
}

// AND is true if all expressions evaluate as true
func AND(expressions ...interface{}) interface{} {
	return &bson.M{
		"$and": expressions,
	}
}

// IsAnyTrue returns true if any element in the array is not 'false, null, 0 or undefined'
func IsAnyTrue(expressions ...interface{}) interface{} {
	return &bson.M{
		"$anyElementTrue": expressions,
	}
}

// ArrayElementAt gets the element at a specfifed location in the array
func ArrayElementAt(array interface{}, i interface{}) interface{} {
	return &bson.M{
		"$arrayElemAt": bson.A{
			array,
			i,
		},
	}
}

// ArrayToObject takes an array of kv arrays or objects and maps them to one object
func ArrayToObject(array interface{}) interface{} {
	return &bson.M{
		"$arrayToObject": array,
	}
}

// Avg takes an expression, evaluates it to an array and returns the single average value of the items in the array
func Avg(expression interface{}) interface{} {
	return &bson.M{
		"$avg": expression,
	}
}

// AvgArray takes an array and returns the single average value of the items in the array
func AvgArray(expressions []interface{}) interface{} {
	return &bson.M{
		"$avg": expressions,
	}
}

// Ceil returns the smallest integer greater than or equal to the specified number
func Ceil(n interface{}) interface{} {
	return &bson.M{
		"$ceil": n,
	}
}

// Cmp takes two values and compares them
// -1 if the first value is less than the second
// 1 if the first value is greater than the second
// 0 if the two values are equivalent
func Cmp(first interface{}, second interface{}) interface{} {
	return &bson.M{
		"$cmp": bson.A{
			first,
			second,
		},
	}
}

// Concat takes an array of expressions that evaluate to string and return the joined string of these
func Concat(strings ...interface{}) interface{} {
	return &bson.M{
		"$concat": strings,
	}
}

// ConcatArray takes an array of expressions that evaluate to array and return the joined array of these
func ConcatArray(arrays ...interface{}) interface{} {
	return &bson.M{
		"$concatArrays": arrays,
	}
}

// Cond takes an boolean expression and returns an expression depending on the value of the boolean expression
func Cond(expression interface{}, trueCase interface{}, falseCase interface{}) interface{} {
	return &bson.M{
		"$cond": &bson.M{
			"if":   expression,
			"then": trueCase,
			"else": falseCase,
		},
	}
}

// Convert one type too another. onError and onNull are optional and thus may be nil
func Convert(input interface{}, typ types.Type, onError interface{}, onNull interface{}) interface{} {
	d := bson.M{
		"input": input,
		"to":    typ,
	}

	if onError != nil {
		d["onError"] = onError
	}

	if onNull != nil {
		d["onNull"] = onNull
	}

	return &bson.M{
		"$convert": d,
	}
}

// DateFromParts makes a date from the parts of that date
func DateFromParts(year interface{}, month interface{}, day interface{}, hour interface{}, minute interface{}, second interface{}, millisecond interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$dateFromParts": &bson.M{
			"year":        year,
			"month":       month,
			"day":         day,
			"hour":        hour,
			"minute":      minute,
			"second":      second,
			"millisecond": millisecond,
			"timezone":    timezone,
		},
	}
}

// DateFromPartsISO makes an ISO date from the parts of that date
func DateFromPartsISO(weekYear interface{}, week interface{}, dayOfWeek interface{}, hour interface{}, minute interface{}, second interface{}, millisecond interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$dateFromParts": &bson.M{
			"isoWeekYear":  weekYear,
			"isoWeek":      week,
			"isoDayOfWeek": dayOfWeek,
			"hour":         hour,
			"minute":       minute,
			"second":       second,
			"millisecond":  millisecond,
			"timezone":     timezone,
		},
	}
}
