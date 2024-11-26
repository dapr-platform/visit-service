package config

import "os"

var (
	ZLMEDIAKIT_API_URL = "http://zlmediakit/index/api"
	ZLMEDIAKIT_SECRET = "035c73f7-bb6b-4889-a715-d9eb2d1925cc"
)

func init() {
	if val := os.Getenv("ZLMEDIAKIT_API_URL"); val != "" {
		ZLMEDIAKIT_API_URL = val
	}
	if val := os.Getenv("ZLMEDIAKIT_SECRET"); val != "" {
		ZLMEDIAKIT_SECRET = val
	}
}
