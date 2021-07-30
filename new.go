package nilcheck

import "database/sql"

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt64(s int64) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}

func NewNullFloat64(s float64) sql.NullFloat64 {
	if s == 0 {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: s,
		Valid:   true,
	}
}
