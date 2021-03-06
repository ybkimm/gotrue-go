package gotrueapi

import (
	"net/http"

	"go.lair.cx/gotrue-go/internal/reqbuilder"
)

func Logout(host string, accessToken string) (*http.Request, error) {
	return reqbuilder.New().
		Method("POST").
		Headers("Authorization", "Bearer "+accessToken).
		Host(host).
		Path("/logout").
		Build()
}
