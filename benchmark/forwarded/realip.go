package forwarded

import (
	"net"
	"strings"
)

func GetIpWithoutSplit(forwarded string) string {
	const forPrefix = "for="
	const ipv6Prefix = `"[`
	const ipv6Suffix = `]"`

	forIndex := strings.Index(forwarded, "for=")
	if forIndex == -1 {
		return ""
	}

	forValueStart := forIndex + len(forPrefix)
	forValueEnd := strings.Index(forwarded[forValueStart:], ";")
	if forValueEnd == -1 {
		forValueEnd = len(forwarded)
	} else {
		forValueEnd += forValueStart // add offset to get the real index considering the entire header
	}

	forComponent := forwarded[forValueStart:forValueEnd]
	i := strings.Index(forComponent, ",") // first comma separates multiple IPs, we only want the first one
	if i == -1 {
		i = len(forComponent)
	}

	realIP := forComponent[:i]

	// Normalize IPv6 addresses
	if strings.HasPrefix(realIP, ipv6Prefix) && strings.HasSuffix(realIP, ipv6Suffix) {
		realIP = strings.Trim(realIP, `[]"`)
	}

	if realIP == "" || net.ParseIP(realIP) == nil {
		return ""
	}

	return realIP
}

func GetIpWithSplit(forwarded string) string {
	const forPrefix = "for="
	const ipv6Prefix = `"[`
	const ipv6Suffix = `]"`

	components := strings.Split(forwarded, ";")

	var realIP string

	for _, component := range components {
		if strings.HasPrefix(component, forPrefix) {
			realIP = strings.Split(strings.TrimPrefix(component, forPrefix), ",")[0]
			break
		}
	}

	// Normalize IPv6 addresses
	if strings.HasPrefix(realIP, ipv6Prefix) && strings.HasSuffix(realIP, ipv6Suffix) {
		realIP = strings.Trim(realIP, `[]"`)
	}

	if realIP == "" || net.ParseIP(realIP) == nil {
		return ""
	}

	return realIP
}
