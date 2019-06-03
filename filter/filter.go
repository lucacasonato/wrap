package filter

import (
	"github.com/lucacasonato/wrap/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Filter fields
type Filter bson.M

// Equal matches if field value is equal to the specified value
func Equal(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$eq": value,
		},
	})
}

// GreaterThan matches if field value is greater than the specified value
func GreaterThan(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$gt": value,
		},
	})
}

// GreaterThanOrEqual matches if field value is greater than or equal to the specified value
func GreaterThanOrEqual(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$gte": value,
		},
	})
}

// LessThan matches if field value is less than the specified value
func LessThan(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$lt": value,
		},
	})
}

// LessThanOrEqual matches if field value is less than or equal to the specified value
func LessThanOrEqual(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$lte": value,
		},
	})
}

// ArrayContains matches if the array in the field contains the specified value
func ArrayContains(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$in": value,
		},
	})
}

// NotEqual matches if field value is not equal to the specified value
func NotEqual(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$ne": value,
		},
	})
}

// ArrayNotContains matches if the array in the field does not contain the specified value
func ArrayNotContains(field string, value interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$nin": value,
		},
	})
}

// AND matches if all of the filters match
func AND(filters ...Filter) Filter {
	return Filter(bson.M{
		"$and": filters,
	})
}

// OR matches if any of the filters match
func OR(filters ...Filter) Filter {
	return Filter(bson.M{
		"$or": filters,
	})
}

// NOT matches if the filter does not match and doesn't match if the filter matches
func NOT(filter Filter) Filter {
	return Filter(bson.M{
		"$not": bson.A{
			filter,
		},
	})
}

// NOR matches if all filters are false
func NOR(filters ...Filter) Filter {
	return Filter(bson.M{
		"$nor": filters,
	})
}

// Exists matches if the field exists or not
func Exists(field string, exists bool) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$exists": exists,
		},
	})
}

// IsType matches if the field is of the specified type
func IsType(field string, typ types.Type) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$type": typ,
		},
	})
}

// TODO: implement schema

// Regex matches if the field value matches the regular expression
func Regex(field string, regex string) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$regex": primitive.Regex{Pattern: regex},
		},
	})
}

// TextSearch matches if the field contains the specified text
func TextSearch(text string) Filter {
	return Filter(bson.M{
		"$text": bson.M{
			"$search": text,
		},
	})
}

// Modulo matches if the remainder of devision on the fields value with the deviser equals the specified remainder
func Modulo(field string, divisor int, remainder int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$mod": bson.A{
				divisor,
				remainder,
			},
		},
	})
}

// JavascriptExpression matches if the field contains the specified text
func JavascriptExpression(expression string) Filter {
	return Filter(bson.M{
		"$where": primitive.JavaScript(expression),
	})
}

// ArrayAll matches if the field contains all of the specified items
func ArrayAll(field string, items []interface{}) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$all": items,
		},
	})
}

// ArraySingleMatch matches if any of the items in the array match all filters
func ArraySingleMatch(field string, filters ...Filter) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$elemMatch": filters,
		},
	})
}

// ArraySize matches if the array has the specified size
func ArraySize(field string, size int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$size": size,
		},
	})
}

// BitsAll0 matches if all bits in a binary value (that are matched by the bitmask) are 0
func BitsAll0(field string, mask int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$bitsAllClear": mask,
		},
	})
}

// BitsAll1 matches if all bits in a binary value (that are matched by the bitmask) are 1
func BitsAll1(field string, mask int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$bitsAllSet": mask,
		},
	})
}

// BitsAny0 matches if any bit in a binary value (that is matched by the bitmask) is 0
func BitsAny0(field string, mask int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$bitsAnyClear": mask,
		},
	})
}

// BitsAny1 matches if any bit in a binary value (that is matched by the bitmask) is 1
func BitsAny1(field string, mask int) Filter {
	return Filter(bson.M{
		field: bson.M{
			"$bitsAnySet": mask,
		},
	})
}
