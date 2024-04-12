package pkg

import (
	"database/sql"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}

func NullIntToInt(n sql.NullInt64) int {
	if n.Valid {
		return int(n.Int64)
	}

	return 0
}

func NullFloatToFloat(f sql.NullFloat64) float64 {
	if f.Valid {
		return f.Float64
	}

	return 0.0
}