package pgor

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	NullNumeric = pgtype.Numeric{
		Int:   nil,
		Exp:   0,
		Valid: false,
	}
)

func Numeric(n big.Int, e ...int32) pgtype.Numeric {
	var exp int32 = 0
	if len(e) > 0 {
		exp = e[0]
	}

	return pgtype.Numeric{
		Int:   &n,
		Exp:   exp,
		Valid: true,
	}
}
