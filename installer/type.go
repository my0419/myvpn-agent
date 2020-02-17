package installer

import "errors"

const (
	TypeL2TP        = "l2tp"
	TypePPTP        = "pptp"
	TypeOpenVPN     = "openvpn"
	TypeWireGuard   = "wireguard"
	TypeShadowsocks = "shadowsocks"

	TypeTestSuccess = "testSuccess"
	TypeTestFail    = "testFail"
)

type Type string

func listOfTypes() map[string]string {
	return map[string]string{
		TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/4dd457a9215b87e68a8c5033edcf845b79a93e48/l2tp.sh",
		TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/9d17eab8bf61285cb84f4eb03340b7118d615574/pptp.sh",
		TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/92fb4523f5d211b4a09f8c8b6f15df041202bf96/openvpn.sh",
		TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/350380fb54fbc89f9c28026911c4c179ae6831b9/wireguard.sh",
		TypeTestSuccess: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
		TypeTestFail:    "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
		TypeShadowsocks: "https://gist.githubusercontent.com/my0419/44f0e8aa1b70fa887432ad6d1d376830/raw/f43e9360e080d9c749d7acad8155d436cd3cfa29/shadowsocks.sh",
	}
}

func (t Type) script() string {
	return listOfTypes()[string(t)]
}

func (t Type) valid() bool {
	return t.script() != ""
}

func createType(alias string) (*Type, error) {
	typeItem := Type(alias)
	if false == typeItem.valid() {
		return nil, errors.New("This type is not supported")
	}
	return &typeItem, nil
}
