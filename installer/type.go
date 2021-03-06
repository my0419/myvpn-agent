package installer

import "errors"

const (
	TypeL2TP        = "l2tp"
	TypePPTP        = "pptp"
	TypeOpenVPN     = "openvpn"
	TypeWireGuard   = "wireguard"
	TypeShadowsocks = "shadowsocks"
	TypeSocksFive   = "socks5"

	TypeTestSuccess = "testSuccess"
	TypeTestFail    = "testFail"
)

type Type string

func listOfTypes() map[string]string {
	return map[string]string{
		TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/f6a1f2426f8ed7af889f4ed39faf0e1505ee2a6f/l2tp.sh",
		TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/2889061402fbe1dd247316caed0ca45722ca9b07/pptp.sh",
		TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/c20873b922cb06224d3a9b2be9a38676dec4365a/openvpn.sh",
		TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/168aa42b994625d4160602d52787d068ee84b9bd/wireguard.sh",
		TypeTestSuccess: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
		TypeTestFail:    "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
		TypeShadowsocks: "https://gist.githubusercontent.com/my0419/44f0e8aa1b70fa887432ad6d1d376830/raw/efb5bf933038ef9bb5291f7a506ee0ac2267fc0f/shadowsocks.sh",
		TypeSocksFive:   "https://gist.githubusercontent.com/my0419/5ddf74cb80eed50d19ae799d37fcaecc/raw/60439e05f9251b9e1e7c18c2331360b4793c10a2/socks5.sh",
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
