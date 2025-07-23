package tencentcdn

import (
	"cnb.cool/znb/cdn-refresh/pkg/tools"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Refresh(r tools.RefreshReq) error {
	credential := common.NewCredential(r.Ak, r.Sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)

	switch r.Rtype {
	case "url":
		return refreshUrlCache(client, r.Urls)
	case "path":
		return refreshDirCache(client, r.Urls)
	default:
		return nil
	}
}

// refreshUrlCache refreshes the cache of the specified urls
func refreshUrlCache(client *cdn.Client, urls []string) error {
	request := cdn.NewPurgeUrlsCacheRequest()
	request.Urls = common.StringPtrs(urls)
	request.UrlEncode = common.BoolPtr(true)
	_, err := client.PurgeUrlsCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

// refreshDirCache refreshes the cache of the specified directories
func refreshDirCache(client *cdn.Client, urls []string) error {
	request := cdn.NewPurgePathCacheRequest()
	request.Paths = common.StringPtrs(urls)
	request.FlushType = common.StringPtr("flush")
	request.UrlEncode = common.BoolPtr(true)
	_, err := client.PurgePathCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
