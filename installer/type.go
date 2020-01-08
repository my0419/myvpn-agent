package installer

import "errors"

const (
	TypeL2TP      = "l2tp"
	TypePPTP      = "pptp"
	TypeOpenVPN   = "openvpn"
	TypeWireGuard = "wireguard"

	TypeTestSuccess = "testSuccess"
	TypeTestFail    = "testFail"
)

type Type string

func listOfTypes() map[string]string {
	return map[string]string{
		TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/4357bafd1e25192b42818cac785d205b4e0bd84b/l2tp.sh",
		TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/80d309fc7a08817a94d4d1ffc15eddf2bca4d29f/pptp.sh",
		TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/6e87ac14b3ce94ee94938dcfac9eb80f8d974ef5/openvpn.sh",
		TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/f40f94435d75efefb9bbe5fa5fbc1d7783484387/wireguard.sh",
		TypeTestSuccess: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
		TypeTestFail:    "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
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
