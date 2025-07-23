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

func Refresh(r tools.RefreshReq) error {
	err := InitClient(r.Ak, r.Sk)
	if err != nil {
		return err
	}
	refreshDcdnObjectCachesRequest := &dcdn20180115.RefreshDcdnObjectCachesRequest{
		ObjectType: tea.String(tools.AliGetRefreshType(r.Rtype)),
		Force:      tea.Bool(false),
		ObjectPath: tea.String(tools.AliGetUrls(r.Urls)),
	}
	_, err = DcdnClient.RefreshDcdnObjectCachesWithOptions(refreshDcdnObjectCachesRequest, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	return nil
}
