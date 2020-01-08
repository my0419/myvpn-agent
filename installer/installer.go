package installer

import (
	"os/exec"
	"os"
	"time"
	"fmt"
	"log"
	"io/ioutil"
)

type Installer struct {
	Type  Type
	State State
}

func (i *Installer) Start() {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("wget -O - %s | bash", i.Type.script()))
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Start script execution")
	i.State.Status.setSetup()
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprintf("Finish script execution. Error %s", err.Error()))
		i.State.Status.setError(err.Error())
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

func CreateInstaller(typeAlias string) (*Installer, error) {
	typeItem, err := createType(typeAlias)
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
