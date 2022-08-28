package helpers

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var PathPrefix string
var Port int
var Bearers map[string]string

func init() {
	bearerPostPtr := flag.String("bearerPost", "", "Authentication bearer expected to push data. If empty, authentication is disabled for POST.")
	bearerGetPtr := flag.String("bearerGet", "", "Authentication bearer expected to pull data. If empty, authentication is disabled for GET.")
	portPtr := flag.Int("port", 9101, "Port to bind on")
	pathPrefixPtr := flag.String("pathPrefix", "", "Path Prefix (useful when server is not started at root context).")

	flag.Parse()
	Port = *portPtr
	PathPrefix = *pathPrefixPtr

	Bearers["GET"] = *bearerGetPtr
	Bearers["POST"] = *bearerPostPtr

	for k := range Bearers {
		if Bearers[k] == "" {
			Bearers[k] = os.Getenv(fmt.Sprintf("GREG_%s_BEARER", k))
		}

		if Bearers[k] != "" {
			Bearers[k] = "Bearer " + Bearers[k]
		}
	}

	if strings.HasSuffix(PathPrefix, "/") {
		PathPrefix = PathPrefix[0 : len(PathPrefix)-1]
	}
}
