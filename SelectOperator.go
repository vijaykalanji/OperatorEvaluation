package OperatorEvaluation

type SelectionOperator struct {
	key     string
	value   string
	child   Operator
	currRec Record
}

func (f *SelectionOperator) NewSelectionOperator(inputKey string, inputVal string, child Operator) Operator {
	return &SelectionOperator{
		key:   inputKey,
		value: inputVal,
		child: child,
	}
}

func (f *SelectionOperator) Next() bool {

	for f.child.Next() {
		record := f.child.Execute()

		// for each cell in the record we check if there is a cell matching this criterion
		for i := 0; i < len(record.entries); i++ {
			entry := record.entries[i]

			if entry.field == f.key && entry.value == f.value {
				f.currRec = record
				return true
			}
		}

	}

	return false
}

func (f *SelectionOperator) Execute() Record {
	return f.currRec
}
