package pgor

import "github.com/jackc/pgx/v5/pgtype"

var (
	NullBoolean = pgtype.Bool{
		Bool:  false,
		Valid: false,
	}
)

func Boolean(b bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  b,
		Valid: true,
	}
}
