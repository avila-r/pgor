package pgor

import "github.com/jackc/pgx/v5/pgtype"

var (
	NullBigint = pgtype.Int8{
		Int64: 0,
		Valid: false,
	}
)

func Bigint(i int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: i,
		Valid: true,
	}
}
