package smschina

import "shlianlu"

// SendNormalSmsReq 发送普通短信
type SendNormalSmsReq struct {
	Type           string   `json:"Type"`               //短信类型,本接口固定值1,
	PhoneNumberSet []string `json:"PhoneNumberSet"`     //接收短信的手机号码数组,上限为10000,每个元素一个手机号码,
	SignName       string   `json:"SignName"`           //短信签名以实名认证公司简称或品牌名称命名,请前往联麓客户端点击(通知/营销短信)-短信签名提交审核,审核通过即可使用,未提交审核签名无法使用,
	SessionContext string   `json:"SessionContext"`     //短信内容中无需带签名,提交营销短信时内容结尾需要带:拒收请回复R(短信格式不正确将会被驳回,详情咨询对接商务),
	TaskTime       int64    `json:"TaskTime,omitempty"` //定时短信,设置短信按照预定的时间发送,UNIX时间戳,
	Tag            string   `json:"Tag"`                //自定义标签,当请求传入此参数时则拉取/推送报告时也会携带此参数,
}

func (req *SendNormalSmsReq) Do(cli *shlianlu.Client) (res *SendSmsRes, err error) {
	req.Type, res = "1", &SendSmsRes{}
	if err = cli.Pose("/sms/trade/normal/send", "1.2.0", req, res); err != nil {
		return nil, err
	}
	return
}

// SendPersonalSmsReq 发送个性短信
type SendPersonalSmsReq struct {
	Type              string     `json:"Type"`               //短信类型,本接口固定值2,
	SignName          string     `json:"SignName"`           //短信签名以实名认证公司简称或品牌名称命名,请前往联麓客户端点击(通知/营销短信)-短信签名提交审核,审核通过即可使用,未提交审核签名无法使用,
	SessionContextSet []string   `json:"SessionContextSet"`  //建立一个内容模板,内容相应位置建立需要变量的元素：{%1%}{%2%}排列 ,提交营销短信时内容结尾需要带:拒收请回复R(短信格式不正确将会被驳回,详情咨询对接商务)
	ContextParamSet   [][]string `json:"ContextParamSet"`    //提交号码和变量元素,号码放在变量元素之前,每个号码和变量元素代表需要形成的一个整体内容,
	TaskTime          int64      `json:"TaskTime,omitempty"` //定时短信,设置短信按照预定的时间发送,UNIX时间戳,
	Tag               string     `json:"Tag"`                //自定义标签,当请求传入此参数时则拉取/推送报告时也会携带此参数,
}

func (req *SendPersonalSmsReq) Do(cli *shlianlu.Client) (res *SendSmsRes, err error) {
	req.Type, res = "2", &SendSmsRes{}
	if err = cli.Pose("/sms/trade/personal/send", "1.2.0", req, res); err != nil {
		return nil, err
	}
	return
}

// SendTemplateSmsReq 发送模版短信
type SendTemplateSmsReq struct {
	Type             string   `json:"Type"`               //短信类型,本接口固定值3,
	TemplateId       int64    `json:"TemplateId"`         //模版id，请登录联麓客户端点击(通知/营销短信)进入短信模板页面创建模板获取,
	PhoneNumberSet   []string `json:"PhoneNumberSet"`     //接收短信的手机号码数组,上限为10000,每个元素一个手机号码,
	TemplateParamSet []string `json:"TemplateParamSet"`   //提交模板变量内容，根据平台创建模板变量数提交相应变量
	TaskTime         int64    `json:"TaskTime,omitempty"` //定时短信,设置短信按照预定的时间发送,UNIX时间戳,
	Tag              string   `json:"Tag"`                //自定义标签,当请求传入此参数时则拉取/推送报告时也会携带此参数,
}

func (req *SendTemplateSmsReq) Do(cli *shlianlu.Client) (res *SendSmsRes, err error) {
	req.Type, res = "3", &SendSmsRes{}
	if err = cli.Pose("/sms/trade/template/send", "1.2.0", req, res); err != nil {
		return nil, err
	}
	return
}

type SendSmsRes struct {
	Message   string `json:"message"`   // 返回信息
	Status    string `json:"status"`    // 接口请求状态码
	Timestamp int64  `json:"timestamp"` // UNIX 时间戳
	Count     int    `json:"count"`     // 返回数量
	Tag       string `json:"tag"`       // 自定义标签
	TaskId    string `json:"taskId"`    // 唯一请求 ID,每次请求都会返回,定位问题时需要提供该次请求的 RequestId,
}
