package voice

import "github.com/cyjaysong/shlianlu"

// CreateTemplateReq 创建语音模版
type CreateTemplateReq struct {
	TemplateName string `json:"TemplateName"` //短信模板名称
	TemplateType string `json:"TemplateType"` //模板类型(TXT,WAV)
	// SessionContext 模版内容
	// 1、模板如需变量,请在模板文本相应位置插入变量,如:尊敬的{%变量1%},您本次登录平台验证码为:{%变量2%},两分钟内有效;
	// 2、音频文件格式仅支持采用PCM编码的wav格式音频文件,比特率:16位,音频采样率:8000(单声道)/44100(立体声);请将音频文件控制在1分钟以内,大小控制在8MB以内
	SessionContext string `json:"SessionContext"`
}

func (req *CreateTemplateReq) Do(cli *shlianlu.Client) (res *CreateTemplateRes, err error) {
	res = &CreateTemplateRes{}
	if err = cli.Pose("/sms/voice/template/create", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type CreateTemplateRes struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	Timestamp  int64  `json:"timestamp"`
	TemplateId int64  `json:"TemplateId"`
}
