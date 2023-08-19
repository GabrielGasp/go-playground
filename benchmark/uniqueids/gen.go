package uniqueids

import (
	gofrsUUID "github.com/gofrs/uuid"
	googleUUID "github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

func UUIDV4Google() string {
	return googleUUID.New().String()
}

func UUIDV4Gofrs() string {
	return gofrsUUID.Must(gofrsUUID.NewV4()).String()
}

func UUIDV7() string {
	return gofrsUUID.Must(gofrsUUID.NewV7()).String()
}

func ULID() string {
	return ulid.Make().String()
}

func KSUID() string {
	return ksuid.New().String()
}
