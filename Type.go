package OperatorEvaluation

type Operator interface {
	Next() bool
	Execute() Record
}

// An entry is the smallesst unit. It ll have a field name and its corresponding value. Represents a cell
type Entry struct {
	field string
	value string
}

// A sequence of entries. Represents  a row
type Record struct {
	entries []Entry
}
