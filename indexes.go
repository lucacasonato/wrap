package wrap

// Index used to index data
type Index interface{}

var (
	// AscendingIndex goes from start to end
	AscendingIndex Index = 1
	// DescendingIndex goes from end to start
	DescendingIndex Index = -1
	// TextIndex indexes text
	TextIndex Index = "text"
)
