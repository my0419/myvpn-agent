package installer

import "errors"

const (
	StagePre  = "pre"
	StagePost = "post"
)

const (
	TypeL2TP        = "l2tp"
	TypeIKEV2       = "ikev2"
	TypePPTP        = "pptp"
	TypeOpenVPN     = "openvpn"
	TypeWireGuard   = "wireguard"
	TypeShadowsocks = "shadowsocks"
	TypeSocksFive   = "socks5"
	OpenConnect     = "openconnect"
	TypeOwncloud    = "owncloud"
	TypeNextcloud   = "nextcloud"
	TypeTorBridge   = "torbridge"

	TypeTestSuccess = "testSuccess"
	TypeTestFail    = "testFail"
)

type Type struct {
	alias string
}

func listOfPreScripts() map[string]map[string]string {
	return map[string]map[string]string{
		TypeL2TP: {
			StagePre: "https://gist.githubusercontent.com/my0419/560071251f2427a9a19862d8a04edb94/raw/5386506d045b5dd73e83b88f58e2ce4045708679/l2tp.sh",
		},
		TypeIKEV2: {
			StagePre: "https://gist.githubusercontent.com/my0419/4bdd4dbfaa4d488ddac1c5fe268bf3a8/raw/75c50f71af01ef6a77797c3e231d8a020a87b72a/ikev2.sh",
		},
		TypePPTP: {
			StagePre: "https://gist.githubusercontent.com/my0419/db77a7bdb466b9df01ffa3f96f4b3f37/raw/c4df04e15b9a5b2edfd4cdad1aef32a71cea9e28/pptp.sh",
		},
		TypeOpenVPN: {
			StagePre: "https://gist.githubusercontent.com/my0419/73ba68f383b5772030078ec871456b06/raw/116255c26aa565f7e00361606f196c3694257ba5/openvpn.sh",
		},
		TypeWireGuard: {
			StagePre: "https://gist.githubusercontent.com/my0419/dd0111d60375dc756c19a70e0907e32b/raw/81ca35106f7499ad7f556b95cb6693a347a604a9/wireguard.sh",
		},
		TypeTestSuccess: {
			StagePre:  "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
			StagePost: "https://gist.githubusercontent.com/my0419/4b5eeaaa98f4b5ae7eeee7b2f5b5cb9a/raw/e730a11c7d5d2caf9ea63ee672f9a195250d06be/test-success.sh",
		},
		TypeTestFail: {
			StagePre: "https://gist.githubusercontent.com/my0419/b4868d16abdd5e43ee7b58c71b079529/raw/a71ee1dbba6d066282db85587ecbf757f4703281/test-fail.sh",
		},
		TypeShadowsocks: {
			StagePre: "https://gist.githubusercontent.com/my0419/44f0e8aa1b70fa887432ad6d1d376830/raw/d9342b8ee8648c355240d4859696abe3a9e6d2b9/shadowsocks.sh",
		},
		OpenConnect: {
			StagePre: "https://gist.githubusercontent.com/my0419/424b2774b24a21bbf520950d779431f5/raw/f4b54d322919bf8263715523c37d68d8078a4971/openconnect.sh",
		},
		TypeSocksFive: {
			StagePre: "https://gist.githubusercontent.com/my0419/5ddf74cb80eed50d19ae799d37fcaecc/raw/61d65fd4a3875b32d1debe1d8038275d43423462/socks5.sh",
		},
		TypeOwncloud: {
			StagePre:  "https://gist.githubusercontent.com/my0419/e56665dc7a359ad276adcec76521384b/raw/6c8ac174a5bd558dbbb2bb5a2cacd525ab2cdd4f/owncloud.sh",
			StagePost: "https://gist.githubusercontent.com/my0419/46f08c7c462a7577afe8a987bf4c67be/raw/7ece5f20332de13d5ce1d907f884dc2b0a23dd2a/owncloud:post.sh",
		},
		TypeNextcloud: {
			StagePre:  "https://gist.githubusercontent.com/my0419/f8a4d1b4b6d6b91095f5f74400862609/raw/868e728d779ea8d1d0c7714366cbda3fbe267289/nextcloud.sh",
			StagePost: "https://gist.githubusercontent.com/my0419/0583ad72751f2411d390dd580776ddbd/raw/05c5fdb84ced0afba3af71b9a021f5d8b7fd7535/nextcloud:post.sh",
		},
		TypeTorBridge: {
			StagePre: "https://gist.githubusercontent.com/my0419/5b8641dee6af47b8510afeed4a6c765e/raw/b59c9a0dd6c292b8c7de055ec6f570f78b71fa69/torbridge.sh",
		},
	}
}

func (t Type) script(stage string) string {
	types := listOfPreScripts()
	if val, ok := types[t.alias][stage]; ok {
		return val
	}
	return ""
}

func (t Type) valid() bool {
	return t.script(StagePre) != ""
}

func createType(alias string) (*Type, error) {
	typeItem := &Type{
		alias: alias,
	}
	if false == typeItem.valid() {
		return nil, errors.New("type is not supported")
	}
	return typeItem, nil
}
