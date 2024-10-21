package pgor

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	NullDate = pgtype.Date{
		Time:  time.Time{},
		Valid: false,
	}
)

func Time(t time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  t,
		Valid: true,
	}
}

func Now() pgtype.Date {
	t := time.Now()

	return pgtype.Date{
		Time:  t,
		Valid: true,
	}
}

func FromDate(t pgtype.Date) time.Time {
	return t.Time
}
