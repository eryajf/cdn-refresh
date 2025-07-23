package qiniucdn

import (
	"fmt"

	"cnb.cool/znb/cdn-refresh/pkg/tools"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/cdn"
)

func Refresh(r tools.RefreshReq) error {
	mac := auth.New(r.Ak, r.Sk)
	cdnManager := cdn.NewCdnManager(mac)
	switch r.Rtype {
	case "url":
		if len(r.Urls) > 100 {
			return fmt.Errorf("urls length must less than 100")
		}
		_, err := cdnManager.RefreshUrls(r.Urls)
		if err != nil {
			return err
		}
	case "path":
		if len(r.Urls) > 10 {
			return fmt.Errorf("urls length must less than 10")
		}
		_, err := cdnManager.RefreshDirs(r.Urls)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
