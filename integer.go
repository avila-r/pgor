package pgor

import "github.com/jackc/pgx/v5/pgtype"

var (
	NullInteger = pgtype.Int4{
		Int32: 0,
		Valid: false,
	}
)

func Integer(i int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: i,
		Valid: true,
	}
}
