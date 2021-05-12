package config

import (
	"net/http"

	"github.com/fortinetdev/forti-sdk-go/fortimanager2/auth"
)

// Config provides configuration to a FMG client instance
// It saves authentication information and a http connection
// for FMG Client instance to create New connction to FMG
// and Send data to FMG,  etc. (needs to be extended later.)
type Config struct {
	Auth     *auth.Auth
	HTTPCon  *http.Client
	FwTarget string
}
