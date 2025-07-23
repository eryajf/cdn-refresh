package alicdn

import (
	"cnb.cool/znb/cdn-refresh/pkg/tools"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v8/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var CdnClient *cdn20180510.Client

func InitClient(ak, sk string) error {
	var err error
	config := openapi.Config{AccessKeyId: tea.String(ak), AccessKeySecret: tea.String(sk)}
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	CdnClient, err = cdn20180510.NewClient(&config)
	if err != nil {
		return err
	}
	return nil
}

func Refresh(ak, sk, zoneName, rtype string, urls []string) error {
	err := InitClient(ak, sk)
	if err != nil {
		return err
	}
	refreshDcdnObjectCachesRequest := &cdn20180510.RefreshObjectCachesRequest{
		ObjectType: tea.String(tools.AliGetRefreshType(rtype)),
		Force:      tea.Bool(false),
		ObjectPath: tea.String(tools.AliGetUrls(urls)),
	}
	_, err = CdnClient.RefreshObjectCachesWithOptions(refreshDcdnObjectCachesRequest, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	return nil
}
