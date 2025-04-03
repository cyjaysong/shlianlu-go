package voice

import "github.com/cyjaysong/shlianlu"

// SendTemplateVoiceReq 发送模版语音
type SendTemplateVoiceReq struct {
	TemplateId       int64    `json:"TemplateId"`         //模版id，请登录联麓客户端点击（语音服务）进入短信模板页面创建模板获取
	PhoneNumberSet   []string `json:"PhoneNumberSet"`     //接收短信的手机号码数组,上限为10000,每个元素一个手机号码
	TemplateParamSet []string `json:"TemplateParamSet"`   //提交模板变量内容，根据平台创建模板变量数提交相应变量
	TaskTime         int64    `json:"TaskTime,omitempty"` //定时任务,设置任务按照预定的时间发送,UNIX时间戳
	Tag              string   `json:"Tag"`                //自定义标签,当请求传入此参数时则拉取/推送报告时也会携带此参数
}

func (req *SendTemplateVoiceReq) Do(cli *shlianlu.Client) (res *SendVoiceRes, err error) {
	res = &SendVoiceRes{}
	if err = cli.Pose("/sms/voice/send", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

// SendPersonalVoiceReq 发送个性语音
type SendPersonalVoiceReq struct {
	TemplateId      int64      `json:"TemplateId"`         //模版id，请登录联麓客户端点击（语音服务）进入短信模板页面创建模板获取
	ContextParamSet [][]string `json:"ContextParamSet"`    //提交号码和变量元素,号码放在变量元素之前,每个号码和变量元素代表需要形成的一个整体内容
	TaskTime        int64      `json:"TaskTime,omitempty"` //定时任务,设置任务按照预定的时间发送,UNIX时间戳
	Tag             string     `json:"Tag"`                //自定义标签,当请求传入此参数时则拉取/推送报告时也会携带此参数
}

func (req *SendPersonalVoiceReq) Do(cli *shlianlu.Client) (res *SendVoiceRes, err error) {
	res = &SendVoiceRes{}
	if err = cli.Pose("/sms/voice/personal/send", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type SendVoiceRes struct {
	Message   string `json:"message"`   // 返回信息
	Status    string `json:"status"`    // 接口请求状态码
	Timestamp int64  `json:"timestamp"` // UNIX 时间戳
	Count     int    `json:"count"`     // 返回数量
	Tag       string `json:"tag"`       // 自定义标签
	TaskId    string `json:"taskId"`    // 唯一请求 ID,每次请求都会返回,定位问题时需要提供该次请求的 RequestId,
}
