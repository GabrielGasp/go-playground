package hashstruct_test

import (
	"bench/hashstruct"
	"testing"
	"time"
)

func BenchmarkHash(b *testing.B) {
	v := hashstruct.Affiliate{
		ID:                       "01HP7NVMK8W5AYQW4GDJDVW53A",
		Username:                 "Xabss",
		Name:                     "Xablau da Silva",
		Email:                    "xablau@mail.com",
		Phone:                    "11999999999",
		DocumentType:             "CPF",
		DocumentNumber:           "12345678909",
		Type:                     "affiliate",
		MinimumGuaranteedEnabled: true,
		MinimumGuaranteedValue:   1000,
		Status:                   "active",
		UserID:                   "01HP7NYQQGV9S6872FS9MEFMJ1",
		ManagerID:                "01HP7NYV53R9WX8KNCCD09474S",
		CollaboratorID:           "01HP7NYY44N6FBABPX8ZEVW05W",
		CreatedAt:                time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:                time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt:                time.Time{},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hashstruct.Hash(v)
	}
}
