// Copyright 2020 21888
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth 登录
package auth

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/21888/microapp"
)

const (
	apiCode2Session   = "/api/apps/jscode2session"
	apiCode2SessionV2 = "/api/apps/v2/jscode2session"
)

/*
ApiCode2SessionV2Res
V2接口返回的数据
*/
type ApiCode2SessionV2Res struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		SessionKey      string `json:"session_key"`
		Openid          string `json:"openid"`
		AnonymousOpenid string `json:"anonymous_openid"`
		Unionid         string `json:"unionid"`
	} `json:"data"`
}

/*
apiCode2SessionV2Req 请求的数据
*/
type apiCode2SessionV2Req struct {
	Appid         string `json:"appid"`
	Secret        string `json:"secret"`
	AnonymousCode string `json:"anonymous_code"`
	Code          string `json:"code"`
}

/*
code2Session

通过login接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session

GET https://developer.toutiao.com/api/apps/jscode2session
*/

func Code2Session(ctx *microapp.MicroApp, params url.Values) (resp []byte, err error) {
	params.Add("appid", ctx.Config.AppId)
	params.Add("secret", ctx.Config.AppSecret)
	return ctx.Client.HTTPGet(apiCode2Session + "?" + params.Encode())
}

func Code2SessionV2(ctx *microapp.MicroApp, code string, anonymousCode string) (resp *ApiCode2SessionV2Res, err error) {
	params := apiCode2SessionV2Req{
		Appid:         ctx.Config.AppId,
		Secret:        ctx.Config.AppSecret,
		AnonymousCode: anonymousCode,
		Code:          code,
	}
	marshal, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	if get, err := ctx.Client.HTTPPost(apiCode2SessionV2, bytes.NewReader(marshal), "application/json"); err != nil {
		return nil, err
	} else {
		apiCode2SessionV2Res := ApiCode2SessionV2Res{}
		if err = json.Unmarshal(get, &apiCode2SessionV2Res); err != nil {
			return nil, err
		}
		return &apiCode2SessionV2Res, nil
	}

}
