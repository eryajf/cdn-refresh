package alidcdn

import (
	"cnb.cool/znb/cdn-refresh/pkg/tools"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dcdn20180115 "github.com/alibabacloud-go/dcdn-20180115/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var DcdnClient *dcdn20180115.Client

func InitClient(ak, sk string) error {
	var err error
	config := openapi.Config{AccessKeyId: tea.String(ak), AccessKeySecret: tea.String(sk)}
	config.Endpoint = tea.String("dcdn.aliyuncs.com")
	DcdnClient, err = dcdn20180115.NewClient(&config)
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
	refreshDcdnObjectCachesRequest := &dcdn20180115.RefreshDcdnObjectCachesRequest{
		ObjectType: tea.String(tools.AliGetRefreshType(rtype)),
		Force:      tea.Bool(false),
		ObjectPath: tea.String(tools.AliGetUrls(urls)),
	}
	_, err = DcdnClient.RefreshDcdnObjectCachesWithOptions(refreshDcdnObjectCachesRequest, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	return nil
}
