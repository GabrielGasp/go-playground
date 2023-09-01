package main

import (
	"fmt"
	"io"
	"net/http"
)

func getImagesURLs() {
	// file, err := os.Create("ulids.txt")
	// if err != nil {
	// 	panic(err)
	// }

	req, _ := http.NewRequest("GET", "https://onabet.com/api/gs/v2/game?currentSlice=84&group=CASINO&name=", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {

	getImagesURLs()
	// file, err := os.Create("ulids.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// for i := 0; i < 100; i++ {
	// 	ulid := ulid.Make()

	// 	_, err := fmt.Fprintln(file, ulid.String())
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	time.Sleep(1 * time.Millisecond)
	// }

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
