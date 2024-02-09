package hash

import (
	"time"

	"github.com/mitchellh/hashstructure/v2"
)

type Affiliate struct {
	ID                       string
	Username                 string
	Name                     string
	Email                    string
	Phone                    string
	DocumentType             string
	DocumentNumber           string
	Type                     string
	MinimumGuaranteedEnabled bool
	MinimumGuaranteedValue   int
	Status                   string
	UserID                   string
	ManagerID                string
	CollaboratorID           string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                time.Time
}

func Hash(v any) {
	_, err := hashstructure.Hash(v, hashstructure.FormatV2, nil)
	if err != nil {
		panic(err)
	}

	// fmt.Println(hash)
}
