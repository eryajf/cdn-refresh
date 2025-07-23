package tools

type RefreshReq struct {
	Ak, Sk, ZoneName, Rtype string
	Urls                    []string
}
