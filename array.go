package pgor

import "github.com/jackc/pgx/v5/pgtype"

func ListOf[T any](n ...T) pgtype.Array[T] {
	var (
		list []T
	)

	list = append(list, n...)

	return pgtype.Array[T]{
		Elements: list,
		Valid:    true,
	}
}
