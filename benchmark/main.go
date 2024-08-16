package main

import (
	"bench/hash"
	"fmt"
)

func main() {
	s := "55252f79-8874-473b-b0de-c6816a0ef06e.431091077177.rollout-tenants.0"
	fmt.Println(hash.Sum256(s))
	fmt.Println(hash.SumMD5(s))
}
