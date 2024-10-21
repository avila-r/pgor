package pgor

import "github.com/jackc/pgx/v5/pgtype"

var (
	NullText = pgtype.Text{
		String: "",
		Valid:  false,
	}
)

func Text(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  true,
	}
}

func FromText(t pgtype.Text) *string {
	return &t.String
}
