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
		TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/7f46031a95f36890830a94c69decea3a2ba17428/l2tp.sh",
		TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/706b441fbb8877368967e3d2ac49886ad53ba225/pptp.sh",
		TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/6b6272ca48cf1fca9926287e6391fee4779dd1d3/openvpn.sh",
		TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/ab139b3c4e0fa825c2bb513f4d7599002fce127e/wireguard.sh",
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
