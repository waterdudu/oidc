package client

import (
	"testing"
)

func TestNewOidcClient(t *testing.T) {

	clientID := "Ixfx7fwqdgZYT_HHJz_zloBR2-XXXXXX-XXXXXXXXXX=@code.dudu.me"
	clientSecret := "L6dGbU_pdNcjtg-XXXXXXXXXXXXXXXXP9Rr8eRCKF5Brb4VzcoEJpIzl8yIflvWYWIUT9vWV2FYkiUKYUm57XDhCaw9jZUKN"
	discovery := "http://127.0.0.1:5556"
	redirectURL := "http://code.dudu.me:9999/callback"

	oidcClient, err := NewOidcClient(clientID, clientSecret, discovery, redirectURL)

	if err != nil {
		t.Fatalf("NewOidcClient error : %v", err)
	}

	t.Log("NewOidcClient created OK!")
	t.Logf("oidc client : %v", oidcClient)
}
