package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	patEnvName = "DO_ACCESS_TOKEN"
)

var (
	pat    tokenSource
	client *godo.Client
	zone   string
)

func init() {
	pat = tokenSource(os.Getenv(patEnvName))

	if pat == "" {
		bail("Please set the environment variable: %s\n", patEnvName)
	}

	flag.StringVar(&zone, "z", "", "Zone to extract")

	flag.Parse()
}

func getClient() *godo.Client {
	if client != nil {
		return client
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, pat)
	client = godo.NewClient(oauthClient)

	return client
}

func bail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	client := getClient()

	dom := NewDomain(NewZoneName(zone))

	err := dom.Find(client)
	if err != nil {
		bail("Error: %s\n", err.Error())
	}

	fmt.Printf("%s", dom.ZoneFile)

	os.Exit(0)
}
