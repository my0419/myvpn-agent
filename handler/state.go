package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/my0419/myvpn-agent/crypto"
	"github.com/my0419/myvpn-agent/installer"
	"log"
	"net/http"
	"os"
	"time"
)

func HandleState(installer *installer.Installer, stopHttp chan bool, encryptKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			return
		}
		ctx := r.Context()
		state := installer.State
		go func() {
			select {
			case <-ctx.Done():
				if state.Status.IsCompleted() {
					stopHttp <- true
				}
				return
			default:
			}
		}()
		state.TimeRunning = int(time.Now().Sub(state.TimeStarting).Seconds())
		jsonData, _ := json.Marshal(state)
		encryptData, err := crypto.EncryptAES(jsonData, encryptKey)
		if err != nil {
			log.Println(fmt.Sprintf("Error encrypt response %s", err.Error()))
			w.WriteHeader(500)
			return
		}
		if os.Getenv("DEBUG_AGENT") != "" {
			fmt.Printf("Debug Response: %s\n", string(jsonData))
		}

		w.Write([]byte(base64.StdEncoding.EncodeToString([]byte(encryptData))))
	}
}
