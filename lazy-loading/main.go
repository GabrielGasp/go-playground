package main

import (
	"fmt"
	"lazy-loading/database"
)

func main() {
	database.Connect()

	post := database.GetPost(1)

	fmt.Println(post)

	post.LoadAuthor()

	fmt.Println(post)

	post.LoadAuthor()
}
