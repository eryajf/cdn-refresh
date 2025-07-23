package doge

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"cnb.cool/znb/cdn-refresh/pkg/tools"
)

type APIResponse struct {
	Code int     `json:"code"`
	Cost float64 `json:"cost"`
	Data APIData `json:"data"`
	Msg  string  `json:"msg"`
	St   int64   `json:"st"`
}

type APIData struct {
	Count    int    `json:"count"`
	TaskID   string `json:"task_id"`
	URLCount int    `json:"url_count"`
}

// refresh 刷新CDN
func Refresh(r tools.RefreshReq) (*APIResponse, error) {
	urlObj, err := json.Marshal(r.Urls)
	if err != nil {
		log.Fatalln(err)
	}
	params := make(map[string]any)
	params["rtype"] = r.Rtype
	params["urls"] = string(urlObj)
	return DogeCloudAPI(r.Ak, r.Sk, "/cdn/refresh/add.json", params)
}

// DogeCloudAPI 云函数API调用
func DogeCloudAPI(AccessKey, SecretKey, apiPath string, data map[string]any) (*APIResponse, error) {
	var response APIResponse

	var body []byte
	var mime string
	var err error

	values := url.Values{}
	for k, v := range data {
		strVal, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("non-json mode requires string values, key %s got %T", k, v)
		}
		values.Set(k, strVal)
	}
	body = []byte(values.Encode())
	mime = "application/x-www-form-urlencoded"

	signStr := apiPath + "\n" + string(body)
	hmacObj := hmac.New(sha1.New, []byte(SecretKey))
	hmacObj.Write([]byte(signStr))
	sign := hex.EncodeToString(hmacObj.Sum(nil))

	req, err := http.NewRequest("POST", "https://api.dogecloud.com"+apiPath, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Add("Content-Type", mime)
	req.Header.Add("Authorization", "TOKEN "+AccessKey+":"+sign)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
