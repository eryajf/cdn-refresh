package aliesa

import (
	"cnb.cool/znb/cdn-refresh/pkg/tools"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	esa20240910 "github.com/alibabacloud-go/esa-20240910/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

var EsaClient *esa20240910.Client

func InitClient(ak, sk string) error {
	var err error
	config := openapi.Config{AccessKeyId: tea.String(ak), AccessKeySecret: tea.String(sk)}
	config.Endpoint = tea.String("esa.cn-hangzhou.aliyuncs.com")
	EsaClient, err = esa20240910.NewClient(&config)
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
	siteID, err := getSiteID(r.ZoneName)
	if err != nil {
		return err
	}
	var content *esa20240910.PurgeCachesRequestContent
	if r.Rtype == "url" {
		content = &esa20240910.PurgeCachesRequestContent{
			PurgeAll: tea.Bool(false),
			Files:    tools.StringSliceToInterfaceSlice(tea.StringSlice(r.Urls)),
		}
	} else {
		content = &esa20240910.PurgeCachesRequestContent{
			PurgeAll:    tea.Bool(false),
			Directories: tea.StringSlice(r.Urls),
		}
	}
	purgeCachesRequest := &esa20240910.PurgeCachesRequest{
		Content: content,
		SiteId:  siteID,
		Type:    tea.String(tools.AliGetRefreshType(r.Rtype)),
		Force:   tea.Bool(true),
	}
	_, err = EsaClient.PurgeCachesWithOptions(purgeCachesRequest, &util.RuntimeOptions{})
	if err != nil {
		return err
	}
	return nil
}

func getSiteID(siteName string) (*int64, error) {
	listSitesRequest := &esa20240910.ListSitesRequest{
		SiteName: tea.String(siteName),
	}
	response, err := EsaClient.ListSitesWithOptions(listSitesRequest, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return response.Body.Sites[0].SiteId, nil
}
