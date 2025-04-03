package shlianlu

import (
	"crypto/md5"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	reqclient "github.com/imroc/req/v3"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	reqClient *reqclient.Client
	mchId     string
	appId     string
	key       string
}

func NewClient(mchId, appId, key string, devMode bool) (client *Client) {
	client = &Client{mchId: mchId, appId: appId, key: key}
	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(1)
	client.reqClient.SetBaseURL("https://apis.shlianlu.com").SetUserAgent("")
	if devMode {
		client.reqClient.DevMode()
	}
	return
}

type BaseRes[T any] struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Data      T      `json:"data"`
}

var signIgnoreField = map[string]struct{}{
	"PhoneNumberSet":    {},
	"SessionContext":    {},
	"SessionContextSet": {},
	"ContextParamSet":   {},
	"TemplateParamSet":  {},
	"Signature":         {},
	"PhoneList":         {},
	"phoneSet":          {},
}

func (c *Client) Pose(path, version string, req, res any) (err error) {
	jsonBytes, _ := sonic.Marshal(req)
	jsonNode, _ := sonic.Get(jsonBytes)
	_, _ = jsonNode.Set("AppId", ast.NewString(c.appId))
	_, _ = jsonNode.Set("MchId", ast.NewString(c.mchId))
	_, _ = jsonNode.Set("Version", ast.NewString(version))
	_, _ = jsonNode.Set("SignType", ast.NewString("MD5"))
	_, _ = jsonNode.Set("TimeStamp", ast.NewString(strconv.FormatInt(time.Now().Unix(), 10)))

	//Sign
	_ = jsonNode.SortKeys(false)
	jsonFieldNum, _ := jsonNode.Len()
	signVals := make([]string, 0, jsonFieldNum)
	for i := 0; i < jsonFieldNum; i++ {
		pair := jsonNode.IndexPair(i)
		if _, has := signIgnoreField[pair.Key]; has {
			continue
		} else if val, _ := pair.Value.String(); val != "" {
			signVals = append(signVals, fmt.Sprintf("%s=%s", pair.Key, val))
		}
	}
	signVals = append(signVals, fmt.Sprintf("key=%s", c.key))
	sign := fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(signVals, "&"))))
	_, _ = jsonNode.Set("Signature", ast.NewString(strings.ToUpper(sign)))

	// Post
	jsonBytes, _ = jsonNode.MarshalJSON()
	httpRes, err := c.reqClient.R().SetBodyJsonBytes(jsonBytes).Post(path)
	if err != nil {
		return err
	} else if httpRes.StatusCode != 200 {
		return fmt.Errorf("http error, status code:%d, msg:%s", httpRes.StatusCode, httpRes.Status)
	} else if err = httpRes.Unmarshal(res); err != nil {
		return err
	}

	return
}
