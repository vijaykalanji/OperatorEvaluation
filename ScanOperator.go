package OperatorEvaluation

// ScanOperator accesses the bufferpool to fetch a bunch of rows that belong to a table.
type ScanOperator struct {
	records []Record
	idx     int
}

func NewScanOperator(rows []Record) Operator {
	return &ScanOperator{
		records: rows,
		idx:     -1,
	}
}

func (s *ScanOperator) Next() bool {
	s.idx++
	return s.idx < len(s.records)
}

func (s *ScanOperator) Execute() Record {
	return s.records[s.idx]
}
