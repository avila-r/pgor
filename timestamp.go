package pgor

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	NullTimestamp = pgtype.Timestamp{
		Time:  time.Time{},
		Valid: false,
	}
)

func Timestamp(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}
}

func FromTimestamp(t pgtype.Timestamp) *time.Time {
	return &t.Time
}
