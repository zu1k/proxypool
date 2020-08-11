package proxy

import (
	"fmt"
	"testing"
)

func TestParseVmessLink(t *testing.T) {
	//link := "vmess://YXV0bzo5MGQ0YTM3NC1kYWU4LTExZWEtODY3Zi01NjAwMDJlY2MzNDlAZml2ZWRlbWFuZHMubWw6NDQz?remarks=%E6%97%A5%E6%9C%AC5%EF%BC%9A%E7%94%B5%E6%8A%A5%E9%A2%91%E9%81%93%EF%BC%9A"

	link := "vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIkBTU1JTVUItVjI5LeS7mOi0ueaOqOiNkDp0LmNuL0VHSkl5cmwiLA0KICAiYWRkIjogImJ1cmdlcmtpbmdnb29kLm1sIiwNCiAgInBvcnQiOiAiNDQzIiwNCiAgImlkIjogIjRiNWJhMzkwLWRhZjQtMTFlYS04ODEwLTU2MDAwMmVjYzk2NyIsDQogICJhaWQiOiAiNDYiLA0KICAibmV0IjogIndzIiwNCiAgInR5cGUiOiAibm9uZSIsDQogICJob3N0IjogImJ1cmdlcmtpbmdnb29kLm1sIiwNCiAgInBhdGgiOiAiL0M2bGt2UzNLLyIsDQogICJ0bHMiOiAidGxzIg0KfQ=="
	fmt.Println(ParseVmessLink(link))
}
