package tencenteo

import (
	"cnb.cool/znb/cdn-refresh/pkg/tools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
)

func Refresh(r tools.RefreshReq) error {
	credential := common.NewCredential(r.Ak, r.Sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "teo.tencentcloudapi.com"
	client, _ := teo.NewClient(credential, "", cpf)

	// 获取zoneId
	zoneId, err := getZoneId(client, r.ZoneName)
	if err != nil {
		return err
	}

	request := teo.NewCreatePurgeTaskRequest()
	request.ZoneId = common.StringPtr(zoneId)
	request.Type = common.StringPtr(tools.TcGetRefreshType(r.Rtype))
	request.Method = common.StringPtr("invalidate")
	request.Targets = common.StringPtrs(r.Urls)
	_, err = client.CreatePurgeTask(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func getZoneId(client *teo.Client, zoneName string) (string, error) {
	request := teo.NewDescribeZonesRequest()
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(100)
	request.Filters = []*teo.AdvancedFilter{
		&teo.AdvancedFilter{
			Name:   common.StringPtr("zone-name"),
			Values: common.StringPtrs([]string{zoneName}),
			Fuzzy:  common.BoolPtr(false),
		},
	}
	// 返回的resp是一个DescribeZonesResponse的实例，与请求对象对应
	response, err := client.DescribeZones(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return *response.Response.Zones[0].ZoneId, nil
}
