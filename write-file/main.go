package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
)

func main() {
	file, err := os.Create("ulids.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < 100; i++ {
		ulid := ulid.Make()

		_, err := fmt.Fprintln(file, ulid.String())
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Millisecond)
	}

	// for i := 0; i < 100; i++ {
	// 	ulid := ulid.Make()
	// 	ulidAsUuid := uuid.Must(uuid.FromBytes(ulid.Bytes()))
	// 	_, err := fmt.Fprintln(file, ulidAsUuid.String())
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	time.Sleep(1 * time.Millisecond)
	// }

	fmt.Println("Done!")
}
