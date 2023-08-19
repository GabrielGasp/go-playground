package main

import (
	"bench/uniqueids"
	"fmt"
)

func main() {
	fmt.Println("UUID V4 (Google): ", uniqueids.UUIDV4Google())
	fmt.Println("UUID V4 (Gofrs):  ", uniqueids.UUIDV4Gofrs())
	fmt.Println("UUID V7:          ", uniqueids.UUIDV7())
	fmt.Println("ULID:             ", uniqueids.ULID())
	fmt.Println("KSUID:            ", uniqueids.KSUID())
}
