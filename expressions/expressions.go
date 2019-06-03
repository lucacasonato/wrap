package expressions

import (
	"github.com/lucacasonato/wrap/types"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	// Include a certain field
	Include = true
	// Exclude a certain field
	Exclude = false
	// MetaTextScore returns the score of the match for a text filtered query
	MetaTextScore = &bson.M{
		"$meta": "textScore",
	}
)

// Value is the value of a field in the document
func Value(field string) interface{} {
	return "$" + field
}

// Literal makes arbitrary data an expression without parsing
func Literal(data interface{}) interface{} {
	return data
}

// Absolute is the absolute value of a value
func Absolute(n interface{}) interface{} {
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

// Average takes an expression, evaluates it to an array and returns the single average value of the items in the array
// or
// Average takes an array and returns the single average value of the items in the array
func Average(expression interface{}) interface{} {
	return &bson.M{
		"$avg": expression,
	}
}

// Ceil returns the smallest integer greater than or equal to the specified number
func Ceil(n interface{}) interface{} {
	return &bson.M{
		"$ceil": n,
	}
}

// Compare takes two values and compares them
// -1 if the first value is less than the second
// 1 if the first value is greater than the second
// 0 if the two values are equivalent
func Compare(first interface{}, second interface{}) interface{} {
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

// Condition takes an boolean expression and returns an expression depending on the value of the boolean expression
func Condition(expression interface{}, trueCase interface{}, falseCase interface{}) interface{} {
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

// DateFromString makes a date from a format and a string
func DateFromString(dateString interface{}, format interface{}, timezone interface{}, onError interface{}, onNull interface{}) interface{} {
	d := bson.M{
		"dateString": dateString,
		"format":     format,
		"timezone":   timezone,
	}

	if onError != nil {
		d["onError"] = onError
	}

	if onNull != nil {
		d["onNull"] = onNull
	}

	return &bson.M{
		"$dateFromString": d,
	}
}

// DateToString takes a date and formats it into a string
func DateToString(date interface{}, format interface{}, timezone interface{}, onNull interface{}) interface{} {
	d := bson.M{
		"date":     date,
		"format":   format,
		"timezone": timezone,
	}

	if onNull != nil {
		d["onNull"] = onNull
	}

	return &bson.M{
		"$dateToString": d,
	}
}

// DayOfMonth takes a date and returns the day of the month
func DayOfMonth(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$dayOfMonth": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// DayOfWeek takes a date and returns the day of the week
func DayOfWeek(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$dayOfWeek": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// DayOfYear takes a date and returns the day of the year
func DayOfYear(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$dayOfYear": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Divide devides the numerator by the denominator and returns the result
func Divide(numerator interface{}, denominator interface{}) interface{} {
	return &bson.M{
		"$divide": bson.A{
			numerator,
			denominator,
		},
	}
}

// Equals returns true if the two expressions are equal
func Equals(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$eq": bson.A{
			expression1,
			expression2,
		},
	}
}

// Exp returns the e^n (e = Euler's number)
func Exp(n interface{}) interface{} {
	return &bson.M{
		"$exp": n,
	}
}

// Filter an array by checking if the condition matches for an element.
// variableName is the variable to use for for the element
func Filter(array interface{}, variableName interface{}, condition interface{}) interface{} {
	return &bson.M{
		"$filter": &bson.M{
			"input": array,
			"as":    variableName,
			"cond":  condition,
		},
	}
}

// Floor returns the largest integer less than or equal to the specified number
func Floor(n interface{}) interface{} {
	return &bson.M{
		"$floor": n,
	}
}

// GreaterThan is true if expression1 is greater than expression2
func GreaterThan(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$gt": bson.A{
			expression1,
			expression2,
		},
	}
}

// GreaterThanOrEqual is true if expression1 is greater than or equal to expression2
func GreaterThanOrEqual(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$gte": bson.A{
			expression1,
			expression2,
		},
	}
}

// LessThan matches is true if expression1 is less to expression2
func LessThan(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$lt": bson.A{
			expression1,
			expression2,
		},
	}
}

// LessThanOrEqual is true if expression1 is less than or equal to expression2
func LessThanOrEqual(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$lte": bson.A{
			expression1,
			expression2,
		},
	}
}

// Hour takes a date and returnst the hour
func Hour(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$hour": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// IfNull evaluates expression and if it is null uses replacement expression instead
func IfNull(expression interface{}, replacement interface{}) interface{} {
	return &bson.M{
		"$ifNull": &bson.A{
			expression,
			replacement,
		},
	}
}

// ArrayContains returns true if array contains the evaluated value of expression
func ArrayContains(arrayToBeSearched interface{}, searchExpression interface{}) interface{} {
	return &bson.M{
		"$in": &bson.A{
			searchExpression,
			arrayToBeSearched,
		},
	}
}

// ArrayIndex returns the index of the first match of search on the array between start and end indexes (zero based).
// Returns -1 if search can not be satisfied
func ArrayIndex(arrayToBeSearched interface{}, searchExpression interface{}, start interface{}, end interface{}) interface{} {
	return &bson.M{
		"$indexOfArray": &bson.A{
			searchExpression,
			arrayToBeSearched,
			start,
			end,
		},
	}
}

// StringIndexByte returns the index of the first match of substing on the string between start and end indexes (zero based UTF-8 byte index).
// Returns -1 if substring can not be found
func StringIndexByte(stringToBeSearched interface{}, substring interface{}, start interface{}, end interface{}) interface{} {
	return &bson.M{
		"$indexOfBytes": &bson.A{
			stringToBeSearched,
			substring,
			start,
			end,
		},
	}
}

// StringIndexCP returns the index of the first match of substing on the string between start and end indexes (zero based UTF-8 code point).
// Returns -1 if substring can not be found
func StringIndexCP(stringToBeSearched interface{}, substring interface{}, start interface{}, end interface{}) interface{} {
	return &bson.M{
		"$indexOfCP": &bson.A{
			stringToBeSearched,
			substring,
			start,
			end,
		},
	}
}

// IsArray returns true if the expression evaluates to an array
func IsArray(expression interface{}) interface{} {
	return &bson.M{
		"$isArray": bson.A{
			expression,
		},
	}
}

// ISODayOfWeek takes a date and returns the ISO day of the week
func ISODayOfWeek(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$isoDayOfWeek": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// ISOWeek takes a date and returns the ISO week
func ISOWeek(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$isoWeek": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// ISOWeekYear takes a date and returns the ISO week year
func ISOWeekYear(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$isoWeekYear": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Let specifies some variables to use in the expression
func Let(vars interface{}, expression interface{}) interface{} {
	return &bson.M{
		"$let": &bson.M{
			"vars": vars,
			"in":   expression,
		},
	}
}

// Ln returns the natural log of n (loge(n))
func Ln(n interface{}) interface{} {
	return &bson.M{
		"$ln": n,
	}
}

// Log returns the log of n with base 'base'
func Log(n interface{}, base interface{}) interface{} {
	return &bson.M{
		"$log": &bson.A{
			n,
			base,
		},
	}
}

// Log10 returns the base10 log of n
func Log10(n interface{}) interface{} {
	return &bson.M{
		"$log10": n,
	}
}

// TrimPrefix returns the input with chars trimmed from prefix
func TrimPrefix(input interface{}, chars interface{}) interface{} {
	return &bson.M{
		"$ltrim": &bson.M{
			"input": input,
			"chars": chars,
		},
	}
}

// Map loops over input and calls expression with the array element as a variable with the
// variableName name and appends its output to a new array
func Map(input interface{}, variableName interface{}, expression interface{}) interface{} {
	return &bson.M{
		"$map": &bson.M{
			"input": input,
			"as":    variableName,
			"in":    expression,
		},
	}
}

// Max takes an expression, evaluates it to an array and returns the largest of the values in the array
// or
// Max takes an array and returns the largest of the values in the array
func Max(expression interface{}) interface{} {
	return &bson.M{
		"$max": expression,
	}
}

// Min takes an expression, evaluates it to an array and returns the smallest of the values in the array
// or
// Min takes an array and returns the smallest of the values in the array
func Min(expression interface{}) interface{} {
	return &bson.M{
		"$min": expression,
	}
}

// Merge takes an expression, evaluates it to an array and returns the merged object of the elements in the array
// or
// Merge takes an array and returns the merged object of the elements in the array
func Merge(expression interface{}) interface{} {
	return &bson.M{
		"$mergeObjects": expression,
	}
}

// Millisecond takes a date and returns the milliseconds
func Millisecond(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$millisecond": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Second takes a date and returns the seconds
func Second(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$second": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Minute takes a date and returns the minutes
func Minute(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$minute": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Month takes a date and returns the month
func Month(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$minute": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Mod devides the numerator by the denominator and returns the remainder
func Mod(numerator interface{}, denominator interface{}) interface{} {
	return &bson.M{
		"$mod": bson.A{
			numerator,
			denominator,
		},
	}
}

// Multiply multiplies expression1 and expression2 and returns the result
func Multiply(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$multiply": bson.A{
			expression1,
			expression2,
		},
	}
}

// NotEqual checks if expression1 and expression2 are equal
func NotEqual(expression1 interface{}, expression2 interface{}) interface{} {
	return &bson.M{
		"$ne": bson.A{
			expression1,
			expression2,
		},
	}
}

// Not inverts the expression
func Not(expression interface{}) interface{} {
	return &bson.M{
		"$not": bson.A{
			expression,
		},
	}
}

// ObjectToArray turns an object into an array of kv objects
func ObjectToArray(expression interface{}) interface{} {
	return &bson.M{
		"$objectToArray": expression,
	}
}

// OR is true if all expressions evaluate as true
func OR(expressions ...interface{}) interface{} {
	return &bson.M{
		"$or": expressions,
	}
}

// Power returns base to the power of exponent
func Power(base interface{}, exponent interface{}) interface{} {
	return &bson.M{
		"$pow": bson.A{
			base,
			exponent,
		},
	}
}

// Fori returns an array with numbers between start and end incremented by step
func Fori(start interface{}, end interface{}, step interface{}) interface{} {
	return &bson.M{
		"$range": bson.A{
			start,
			end,
			step,
		},
	}
}

// Reduce returns a single value computed from the array with the reducer expression
func Reduce(array interface{}, startValue interface{}, reducer interface{}) interface{} {
	return &bson.M{
		"$reduce": bson.M{
			"input":        array,
			"initialValue": startValue,
			"in":           reducer,
		},
	}
}

// ArrayReverse reverses an array
func ArrayReverse(array interface{}) interface{} {
	return &bson.M{
		"$reverseArray": array,
	}
}

// TrimSuffix returns the input with chars trimmed from end
func TrimSuffix(input interface{}, chars interface{}) interface{} {
	return &bson.M{
		"$rtrim": &bson.M{
			"input": input,
			"chars": chars,
		},
	}
}

// SetDifference returns all distinct items that exist in array 1 but not array 2
func SetDifference(array1 interface{}, array2 interface{}) interface{} {
	return &bson.M{
		"$setDifference": &bson.A{
			array1,
			array2,
		},
	}
}

// SetEquals returns true if all arrays contain the same distinct elements
// (order and amount does not matter)
func SetEquals(arrays ...interface{}) interface{} {
	return &bson.M{
		"$setEquals": arrays,
	}
}

// SetIntersect returns all distinct items that appear in all arrays
func SetIntersect(arrays ...interface{}) interface{} {
	return &bson.M{
		"$setIntersection": arrays,
	}
}

// SetIsSubset returns true if all items from set1 exist in set2
func SetIsSubset(array1 interface{}, array2 interface{}) interface{} {
	return &bson.M{
		"$setIsSubset": &bson.A{
			array1,
			array2,
		},
	}
}

// SetUnion returns an array with all unique items from all arrays
func SetUnion(arrays ...interface{}) interface{} {
	return &bson.M{
		"$setUnion": arrays,
	}
}

// ArraySize returns the length of an array
func ArraySize(array interface{}) interface{} {
	return &bson.M{
		"$size": array,
	}
}

// ArraySlice returns n items from the array starting at the start position
func ArraySlice(array interface{}, n interface{}, startPosition interface{}) interface{} {
	return &bson.M{
		"$slice": bson.A{
			array,
			startPosition,
			n,
		},
	}
}

// StringSplit returns an array of substrings that come from the original string split at the delimiter
func StringSplit(stringToSplit interface{}, delimiter interface{}) interface{} {
	return &bson.M{
		"$split": bson.A{
			stringToSplit,
			delimiter,
		},
	}
}

// SquareRoot returns the square root of n
func SquareRoot(n interface{}) interface{} {
	return &bson.M{
		"$sqrt": n,
	}
}

// StandardDeviationPopulation returns the population standard deviation for the selected data
func StandardDeviationPopulation(population interface{}) interface{} {
	return &bson.M{
		"$stdDevPop": population,
	}
}

// StandardDeviationSample returns the sample standard deviation for the selected data
func StandardDeviationSample(sample interface{}) interface{} {
	return &bson.M{
		"$stdDevSamp": sample,
	}
}

// StringCaseCompare compares the string case insensitvely
// 1 if first string is “greater than” the second string
// 0 if the two strings are equal
// -1 if the first string is “less than” the second string
func StringCaseCompare(string1 interface{}, string2 interface{}) interface{} {
	return &bson.M{
		"$strcasecmp": bson.A{
			string1,
			string2,
		},
	}
}

// StringLengthBytes returns the number of UTF-8 encoded bytes in the specified string
func StringLengthBytes(stringToCheck interface{}) interface{} {
	return &bson.M{
		"$strLenBytes": stringToCheck,
	}
}

// StringLengthCP returns the number of UTF-8 code points in the specified string
func StringLengthCP(stringToCheck interface{}) interface{} {
	return &bson.M{
		"$strLenCP": stringToCheck,
	}
}

// StringSubstring returns the a substring of a specified length starting at a certain position (zero point char based)
func StringSubstring(intialString interface{}, start interface{}, length interface{}) interface{} {
	return &bson.M{
		"$substr": bson.A{
			intialString,
			start,
			length,
		},
	}
}

// StringSubstringBytes returns the a substring with a specified count of UTF-8 bytes starting at a certain index (zero point UTF-8 byte index based)
func StringSubstringBytes(intialString interface{}, index interface{}, count interface{}) interface{} {
	return &bson.M{
		"$substrBytes": bson.A{
			intialString,
			index,
			count,
		},
	}
}

// StringSubstringCP returns the a substring with a specified count of UTF-8 code points starting at a certain index (zero point UTF-8 code point based)
func StringSubstringCP(intialString interface{}, index interface{}, count interface{}) interface{} {
	return &bson.M{
		"$substrCP": bson.A{
			intialString,
			index,
			count,
		},
	}
}

// Subtract returns value1 - value2
func Subtract(value1 interface{}, value2 interface{}) interface{} {
	return &bson.M{
		"$subtract": bson.A{
			value1,
			value2,
		},
	}
}

// Sum adds all values in an array together
func Sum(array ...interface{}) interface{} {
	return &bson.M{
		"$sum": array,
	}
}

// SwitchBranch is a branch to switch on
type SwitchBranch struct {
	Case interface{}
	Then interface{}
}

// Switch checks which case applies and then evalutes that expression. If none match, default is used
func Switch(def interface{}, branches ...SwitchBranch) interface{} {
	return &bson.M{
		"$switch": bson.M{
			"branches": branches,
			"default":  def,
		},
	}
}

// ToBool turns a value into a bool
func ToBool(value interface{}) interface{} {
	return &bson.M{
		"$toBool": value,
	}
}

// ToDate turns a value into a date
func ToDate(value interface{}) interface{} {
	return &bson.M{
		"$toDate": value,
	}
}

// ToDecimal turns a value into a decimal
func ToDecimal(value interface{}) interface{} {
	return &bson.M{
		"$toDecimal": value,
	}
}

// ToDouble turns a value into a double
func ToDouble(value interface{}) interface{} {
	return &bson.M{
		"$toDouble": value,
	}
}

// ToInt turns a value into a int
func ToInt(value interface{}) interface{} {
	return &bson.M{
		"$toInt": value,
	}
}

// ToLong turns a value into a long
func ToLong(value interface{}) interface{} {
	return &bson.M{
		"$toLong": value,
	}
}

// ToObjectID turns a value into an object id
func ToObjectID(value interface{}) interface{} {
	return &bson.M{
		"$toObjectId": value,
	}
}

// ToString turns a value into a string
func ToString(value interface{}) interface{} {
	return &bson.M{
		"$toString": value,
	}
}

// ToLowerString turns a value into a lowercase string
func ToLowerString(value interface{}) interface{} {
	return &bson.M{
		"$toLower": value,
	}
}

// ToUpperString turns a value into an uppercase string
func ToUpperString(value interface{}) interface{} {
	return &bson.M{
		"$toUpper": value,
	}
}

// Trim returns the input with chars trimmed from prefix and suffix
func Trim(input interface{}, chars interface{}) interface{} {
	return &bson.M{
		"$ltrim": &bson.M{
			"input": input,
			"chars": chars,
		},
	}
}

// Trunc returns the number with all decimal places removed
func Trunc(n interface{}) interface{} {
	return &bson.M{
		"$trunc": n,
	}
}

// Type returns the type of the value
func Type(value interface{}) interface{} {
	return &bson.M{
		"$type": value,
	}
}

// Week takes a date and returns the week
func Week(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$week": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// Year takes a date and returns the year
func Year(date interface{}, timezone interface{}) interface{} {
	return &bson.M{
		"$year": &bson.M{
			"date":     date,
			"timezone": timezone,
		},
	}
}

// ArrayZip returns an array of arrays with one item from each array in the sub arrays
func ArrayZip(useLongest bool, def interface{}, arrays ...interface{}) interface{} {
	return &bson.M{
		"$zip": &bson.M{
			"inputs":           arrays,
			"useLongestLength": useLongest,
			"default":          def,
		},
	}
}
