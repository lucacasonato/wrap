package types

// Type is the type of a value
type Type string

const (
	// Number matches doubles, ints and decimals
	Number Type = "number"
	// Double is a floating point number
	Double Type = "double"
	// String is a string
	String Type = "string"
	// Object is a map
	Object Type = "object"
	// Array a list of values
	Array Type = "array"
	// BinaryData is a blob of bytes
	BinaryData Type = "binData"
	// ObjectID is the id for an object
	ObjectID Type = "objectId"
	// Boolean is a true or false value
	Boolean Type = "bool"
	// Date is a date in ISO 8601 format
	Date Type = "date"
	// Null means there is no value
	Null Type = "null"
	// RegularExpression is a way to match something
	RegularExpression Type = "regex"
	// JavaScript is some javascript code
	JavaScript Type = "javascript"
	// JavaScriptWithScope is some javascript code that has a scope
	JavaScriptWithScope Type = "javascriptWithScope"
	// Int is a non decimal 32 bit number
	Int Type = "int"
	// Long is a non decimal 64 bit number
	Long Type = "long"
	// Decimal is a 128 bit floating point number
	Decimal Type = "decimal"
	// Timestamp is a timestamp in seconds after January 1, 1970
	Timestamp Type = "timestamp"
)
