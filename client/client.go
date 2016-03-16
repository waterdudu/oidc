package client

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/oidc"
	"log"
	"net/http"
	"time"

	// phttp "github.com/coreos/dex/pkg/http"
)

func NewOidcClient(clientID, clientSecret, discovery, redirectURL string) (*oidc.Client, error) {

	cc := oidc.ClientCredentials{
		ID:     clientID,
		Secret: clientSecret,
	}

	var tlsConfig tls.Config
	// if *caFile != "" {
	//     roots := x509.NewCertPool()
	//     pemBlock, err := ioutil.ReadFile(*caFile)
	//     if err != nil {
	//         log.Fatalf("Unable to read ca file: %v", err)
	//     }
	//     roots.AppendCertsFromPEM(pemBlock)
	//     tlsConfig.RootCAs = roots
	// }

	httpClient := &http.Client{Transport: &http.Transport{TLSClientConfig: &tlsConfig}}

	var cfg oidc.ProviderConfig
	var err error
	var count int = 1
	for {
		cfg, err = oidc.FetchProviderConfig(httpClient, discovery)
		if err == nil {
			break
		}

		sleep := 3 * time.Second
		fmt.Printf("Failed fetching provider config, trying again in %v: %v", sleep, err)
		time.Sleep(sleep)
		count++
		if count == 3 {
			return &oidc.Client{}, errors.New("discovery timeout error")
		}
	}

	log.Printf("Fetched provider config from %s: %#v", discovery, cfg)

	scopes := append(oidc.DefaultScope, "offline_access")

	ccfg := oidc.ClientConfig{
		HTTPClient:     httpClient,
		ProviderConfig: cfg,
		Credentials:    cc,
		RedirectURL:    redirectURL,
		Scope:          scopes,
	}

	client, err := oidc.NewClient(ccfg)
	if err != nil {
		fmt.Printf("Unable to create Client: %v", err)
		return &oidc.Client{}, err
	}

	client.SyncProviderConfig(discovery)

	return client, nil
}
