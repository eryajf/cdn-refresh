package tencentcdn

import (
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Refresh(AccessKey, SecretKey, rtype string, urls []string) error {
	credential := common.NewCredential(AccessKey, SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)

	switch rtype {
	case "url":
		return refreshUrlCache(client, urls)
	case "path":
		return refreshDirCache(client, urls)
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
