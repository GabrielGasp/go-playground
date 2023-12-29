package forwarded_test

import (
	"bench/forwarded"
	"testing"
)

func BenchmarkRealIpWithoutSplit(b *testing.B) {
	forwardedHeader := `for=192.168.0.1, for=192.168.0.2;host=api.homolog.boongcloud.net;proto=https`

	for i := 0; i < b.N; i++ {
		forwarded.GetIpWithoutSplit(forwardedHeader)
	}
}

func BenchmarkRealIpWithSplit(b *testing.B) {
	forwardedHeader := `for=192.168.0.1, for=192.168.0.2;host=api.homolog.boongcloud.net;proto=https`

	for i := 0; i < b.N; i++ {
		forwarded.GetIpWithSplit(forwardedHeader)
	}
}
