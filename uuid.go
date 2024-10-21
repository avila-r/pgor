package pgor

import "github.com/jackc/pgx/v5/pgtype"

var (
	NullUUID = pgtype.UUID{
		Bytes: [16]byte{},
		Valid: false,
	}
)

func UUID(bytes [16]byte) pgtype.UUID {
	return pgtype.UUID{
		Bytes: bytes,
		Valid: true,
	}
}

func FromUUID(t pgtype.UUID) [16]byte {
	return t.Bytes
}
