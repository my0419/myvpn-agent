package main

import (
	"myvpn-agent/installer"
	"myvpn-agent/handler"
	"net/http"
	"os"
	"log"
)

func main() {

	if len(os.Getenv("ENCRYPT_KEY")) != 32 {
		log.Fatal("Invalid ENCRYPT_KEY")
	}

	if os.Getenv("VPN_CLIENT_CONFIG_FILE") == "" {
		os.Setenv("VPN_CLIENT_CONFIG_FILE", "/tmp/myvpn-client-config")
	}

	setup, err := installer.CreateInstaller(os.Getenv("VPN_TYPE"))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		setup.Start()
	}()
	http.HandleFunc("/", handler.HandleState(setup, os.Getenv("ENCRYPT_KEY")))
	log.Fatal(http.ListenAndServe(":8400", nil))
}