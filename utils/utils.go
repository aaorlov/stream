package utils

import (
	"os"

	"github.com/aaorlov/stream/log"
)

const defaultPort = "8080"

// ResolveAddress address
func ResolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			log.Debugf("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		log.Debugf("Environment variable PORT is undefined. Using port :8080 by default")
		return ":" + defaultPort
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
