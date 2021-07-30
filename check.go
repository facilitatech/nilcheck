package nilcheck

import "database/sql"

func String(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}

func Int64(s sql.NullInt64) int64 {
	if s.Valid {
		return s.Int64
	}

	return 0
}

func Float64(s sql.NullFloat64) float64 {
	if s.Valid {
		return s.Float64
	}

	return 0
}

func Bool(s sql.NullBool) bool {
	if s.Valid {
		return s.Bool
	}

	return false
}
