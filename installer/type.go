package installer

import "errors"

const (
	TypeL2TP        = "l2tp"
	TypePPTP        = "pptp"
	TypeOpenVPN     = "openvpn"
	TypeWireGuard   = "wireguard"
	TypeShadowsocks = "shadowsocks"
	TypeSocksFive   = "socks5"
	TypeOwncloud    = "owncloud"
	TypeNextcloud   = "nextcloud"
	TypeTorBridge   = "torbridge"

	TypeTestSuccess = "testSuccess"
	TypeTestFail    = "testFail"
)

type Type struct {
	alias string
	os    string
}

func listOfTypes() map[string]map[string]string {
	return map[string]map[string]string{
		"debian9": {
			TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/f6a1f2426f8ed7af889f4ed39faf0e1505ee2a6f/l2tp.sh",
			TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/2889061402fbe1dd247316caed0ca45722ca9b07/pptp.sh",
			TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/c20873b922cb06224d3a9b2be9a38676dec4365a/openvpn.sh",
			TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/9bb5c1818428bd9b2605669611bf9d445897b51d/wireguard.sh",
			TypeTestSuccess: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
			TypeTestFail:    "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
			TypeShadowsocks: "https://gist.githubusercontent.com/my0419/44f0e8aa1b70fa887432ad6d1d376830/raw/efb5bf933038ef9bb5291f7a506ee0ac2267fc0f/shadowsocks.sh",
			TypeSocksFive:   "https://gist.githubusercontent.com/my0419/5ddf74cb80eed50d19ae799d37fcaecc/raw/60439e05f9251b9e1e7c18c2331360b4793c10a2/socks5.sh",
			TypeOwncloud:    "https://gist.githubusercontent.com/my0419/e56665dc7a359ad276adcec76521384b/raw/05de6091905f24a6fcf230c4d0a8f97fb23fed4d/owncloud.sh",
			TypeNextcloud:   "https://gist.githubusercontent.com/my0419/f8a4d1b4b6d6b91095f5f74400862609/raw/71c2bcda17786fd9121dd3deb476fa0f743844a7/nextcloud.sh",
			TypeTorBridge:   "https://gist.githubusercontent.com/my0419/5b8641dee6af47b8510afeed4a6c765e/raw/d24875986af26778ed8739e1619e168f20ec7f62/torbridge.sh",
		},
		"debian11": {
			TypeL2TP:        "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/5386506d045b5dd73e83b88f58e2ce4045708679/l2tp.sh",
			TypePPTP:        "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/c4df04e15b9a5b2edfd4cdad1aef32a71cea9e28/pptp.sh",
			TypeOpenVPN:     "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/116255c26aa565f7e00361606f196c3694257ba5/openvpn.sh",
			TypeWireGuard:   "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/81ca35106f7499ad7f556b95cb6693a347a604a9/wireguard.sh",
			TypeTestSuccess: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
			TypeTestFail:    "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
			TypeShadowsocks: "https://gist.githubusercontent.com/my0419/44f0e8aa1b70fa887432ad6d1d376830/raw/d9342b8ee8648c355240d4859696abe3a9e6d2b9/shadowsocks.sh",
			TypeSocksFive:   "https://gist.githubusercontent.com/my0419/5ddf74cb80eed50d19ae799d37fcaecc/raw/61d65fd4a3875b32d1debe1d8038275d43423462/socks5.sh",
			TypeOwncloud:    "https://gist.githubusercontent.com/my0419/e56665dc7a359ad276adcec76521384b/raw/6c8ac174a5bd558dbbb2bb5a2cacd525ab2cdd4f/owncloud.sh",
			TypeNextcloud:   "https://gist.githubusercontent.com/my0419/f8a4d1b4b6d6b91095f5f74400862609/raw/0aec7afd87cdb1e6f72b84f5213b6b2987d45e0d/nextcloud.sh",
			TypeTorBridge:   "https://gist.githubusercontent.com/my0419/5b8641dee6af47b8510afeed4a6c765e/raw/b59c9a0dd6c292b8c7de055ec6f570f78b71fa69/torbridge.sh",
		},
	}
}

func (t Type) script() string {
	types := listOfTypes()
	return types[t.os][t.alias]
}

func (t Type) valid() bool {
	return t.script() != ""
}

func createType(alias string, os string) (*Type, error) {
	typeItem := &Type{
		alias: alias,
		os:    os,
	}
	if false == typeItem.valid() {
		return nil, errors.New("This type is not supported")
	}
	return typeItem, nil
}
