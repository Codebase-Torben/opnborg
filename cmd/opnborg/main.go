package main

import (
	"errors"
	"fmt"
	"os"

	"paepcke.de/opnborg"
)

const _app = opnborg._app

func main() {
	fmt.Println(_app + "[STARTUP][V0.0.1]")
	_, err := readConfig()
	if err != nil {
		fmt.Printf(_app+"[EXIT]%s\n", err)
		os.Exit(1)
	}
	fmt.Println(_app + "[END]")
}

func readConfig() (*opnborg.OPNCall, error) {

	if _, ok := os.LookupEnv("OPN_TARGETS"); !ok {
		return nil, errors.New(fmt.Sprintf("[ERROR] Add at least one target server to env var 'OPN_TARGETS' (multi valued, comma seperated)"))
	}

	if _, ok := os.LookupEnv("OPN_APIKEY"); !ok {
		return nil, errors.New(fmt.Sprintf("[ERROR] Set env var 'OPN_APIKEY' to your opnsense api key"))
	}

	if _, ok := os.LookupEnv("OPN_APISECRET"); !ok {
		return nil, errors.New(fmt.Sprintf("[ERROR] Set env var 'OPN_APISECRET' to your opnsense api key secret"))
	}
	return &opnborg.OPNCall{
		Targets:     os.Getenv("OPN_TARGETS"),
		Key:         os.Getenv("OPN_APIKEY"),
		Secret:      os.Getenv("OPN_APISECRET"),
		NoSSLVerify: os.Getenv("OPN_NOSSLVERIFY") == "1",
	}, nil
}