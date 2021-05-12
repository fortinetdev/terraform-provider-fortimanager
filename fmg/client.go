package fortimanager

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/fortinetdev/forti-sdk-go/fortimanager2/auth"
	forticlient "github.com/fortinetdev/forti-sdk-go/fortimanager2/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname  string
	User      string
	Passwd    string
	Insecure  *bool
	CABundle  string
	ScopeType string
	Adom      string
}

// FortiClient contains the basic FMG SDK connection information to FMG
// It can be used to as a client of FMG for the plugin
type FortiClient struct {
	Client *forticlient.FortiSDKClient
	Cfg    *Config
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	err := createFMGClient(&fClient, c)
	if err != nil {
		return nil, fmt.Errorf("Error create fortimanager client: %v", err)
	}

	return &fClient, nil
}

func createFMGClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Hostname, c.User, c.Passwd, c.CABundle)

	if auth.Hostname == "" {
		_, err := auth.GetEnvHostname()
		if err != nil {
			return fmt.Errorf("Error reading Hostname")
		}
	}

	if auth.User == "" {
		_, err := auth.GetEnvUsername()
		if err != nil {
			return fmt.Errorf("Error reading Username")
		}
	}

	if auth.Passwd == "" {
		_, err := auth.GetEnvPassword()
		if err != nil {
			return fmt.Errorf("Error reading Password")
		}
	}

	if auth.CABundle == "" {
		auth.GetEnvCABundle()
	}

	if auth.CABundle != "" {
		f, err := os.Open(auth.CABundle)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}
		defer f.Close()

		caBundle, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}

		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(caBundle)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	}

	if c.Insecure == nil {
		// defaut value for Insecure is false
		b, _ := auth.GetEnvInsecure()
		config.InsecureSkipVerify = b
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if config.InsecureSkipVerify == false && auth.CABundle == "" {
		return fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 250,
	}

	fc := forticlient.NewClient(auth, client)

	fClient.Cfg = c
	fClient.Client = fc

	return nil
}
