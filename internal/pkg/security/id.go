package security

import "github.com/sqids/sqids-go"

func HashUserId(userId uint) string {
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	id, _ := s.Encode([]uint64{uint64(userId)})

	return id
}
