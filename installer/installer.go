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

func (i *Installer) runScript(script string) (*bytes.Buffer, error) {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("wget -O - %s | bash", script))
	cmd.Env = os.Environ()

	buf := new(bytes.Buffer)

	if system.DebugEnabled() {
		cmd.Stdout = buf
		cmd.Stderr = buf
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return buf, cmd.Run()
}

func (i *Installer) RunPreStage() {
	log.Println("Start pre script execution")
	i.State.Status.setSetup()
	buf, err := i.runScript(i.Type.script(StagePre))

	if err != nil {
		fmt.Printf("Finish script execution. Error %s", err.Error())
		if system.DebugEnabled() {
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

func (i *Installer) RunPostStage() {
	log.Println("Start post script execution")

	s := i.Type.script(StagePost)
	if s == "" {
		log.Println("No post script execution")
		return
	}

	buf, err := i.runScript(s)
	if err != nil {
		fmt.Printf("Post script execution error. %s\nScript output:\n%s", err.Error(), buf.String())
		return
	}
}

func CreateInstaller(t string) (*Installer, error) {
	typeItem, err := createType(t)
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
