package OperatorEvaluation

type ProjectionOperator struct {
	column string
	child  Operator
	record Record
}

// NewProjectionOperator creates a new ProjectionOperator.
func NewProjectionOperator(columnName string, child Operator) Operator {
	return &ProjectionOperator{
		column: columnName,
		child:  child,
	}
}

func (s *ProjectionOperator) Next() bool {
	return s.child.Next()
}

func (s *ProjectionOperator) Execute() Record {
	record := s.child.Execute()

	for i := 0; i < len(record.entries); i++ {
		entry := record.entries[i]
		if entry.field == s.column {
			s.record = record
			break
		}
	}
	return record
}
