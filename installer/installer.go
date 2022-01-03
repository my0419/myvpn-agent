package installer

import (
	"bytes"
	"fmt"
	"github.com/my0419/myvpn-agent/system"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

type Installer struct {
	Type  Type
	State State
}

func (i *Installer) Start() {
	debug := system.DebugEnabled()
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("wget -O - %s | bash", i.Type.script()))
	cmd.Env = os.Environ()

	buf := new(bytes.Buffer)

	if true == debug {
		cmd.Stdout = buf
		cmd.Stderr = buf
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	log.Println("Start script execution")
	i.State.Status.setSetup()
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprintf("Finish script execution. Error %s", err.Error()))
		if true == debug {
			i.State.Status.setError(fmt.Sprintf("Debug mode trace. Error %s\nScript output:\n%s", err.Error(), buf.String()))
		} else {
			i.State.Status.setError(err.Error())
		}
		return
	}
	clientConfigFile, err := os.Open(os.Getenv("VPN_CLIENT_CONFIG_FILE"))
	if err != nil {
		log.Println("Failed open client config file")
		i.State.Status.setError(err.Error())
		return
	}
	clientConfig, _ := ioutil.ReadAll(clientConfigFile)
	i.State.Status.setCompleted(string(clientConfig))
}

func CreateInstaller(t string, os string) (*Installer, error) {
	if os == "" {
		os = "debian9"
	}
	typeItem, err := createType(t, os)
	if err != nil {
		return nil, err
	}
	status := Status{}
	status.setIdle()

	state := State{Status: status, TimeStarting: time.Now()}
	return &Installer{
		Type:  *typeItem,
		State: state,
	}, nil
}
