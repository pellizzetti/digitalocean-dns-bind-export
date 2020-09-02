package main

import (
	"golang.org/x/oauth2"
)

type (
	tokenSource string
)

func (t tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: string(t),
	}

	return token, nil
}
