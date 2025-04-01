package smschina

import (
	"github.com/cyjaysong/shlianlu"
)

// CreateTemplateReq 创建短信模版
type CreateTemplateReq struct {
	SignId       int64  `json:"SignId"`       //签名Id
	Content      string `json:"content"`      //模板内容(无需携带签名)，营销短信模板内容格式：内容 + 退订文案，常见退订文案有：拒收请回复R。
	TemplateName string `json:"TemplateName"` //短信模板名称
}

func (req *CreateTemplateReq) Do(cli *shlianlu.Client) (res *CreateTemplateRes, err error) {
	res = &CreateTemplateRes{}
	if err = cli.Pose("/sms/product/template/create", "1.1.0", req, res); err != nil {
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

// DeleteTemplateReq 删除短信模版
type DeleteTemplateReq struct {
	TemplateId int64 `json:"TemplateId"` //短信模板id
}

func (req *DeleteTemplateReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[struct{}], err error) {
	res = &shlianlu.BaseRes[struct{}]{}
	if err = cli.Pose("/sms/product/template/delete", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

// UpdateTemplateReq 编辑短信模版
type UpdateTemplateReq struct {
	TemplateId   int64  `json:"TemplateId"`   //短信模板id
	SignId       int64  `json:"SignId"`       //签名Id
	Content      string `json:"content"`      //模板内容(无需携带签名)，营销短信模板内容格式：内容 + 退订文案，常见退订文案有：拒收请回复R。
	TemplateName string `json:"TemplateName"` //短信模板名称
}

func (req *UpdateTemplateReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[struct{}], err error) {
	res = &shlianlu.BaseRes[struct{}]{}
	if err = cli.Pose("/sms/product/template/update", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

// GetTemplateOneReq 查询短信模版单个
type GetTemplateOneReq struct {
	TemplateId int64 `json:"TemplateId"` //短信模板id
}

func (req *GetTemplateOneReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[GetTemplateItem], err error) {
	res = &shlianlu.BaseRes[GetTemplateItem]{}
	if err = cli.Pose("/sms/product/template/getById", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

// GetTemplateListReq 查询短信模版列表
type GetTemplateListReq struct{}

func (req *GetTemplateListReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[[]GetTemplateItem], err error) {
	res = &shlianlu.BaseRes[[]GetTemplateItem]{}
	if err = cli.Pose("/sms/product/template/get", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetTemplateItem struct {
	User         string `json:"user"`
	UpmsSign     string `json:"upmsSign"`
	Id           int64  `json:"id"`           //模板Id
	SignId       int64  `json:"signId"`       //短信签名Id
	SignContent  string `json:"signContent"`  //短信签名内容
	SuffixSignId int64  `json:"suffixSignId"` //后缀签名Id
	SuffixSign   string `json:"suffixSign"`   //后缀签名内容
	ProductId    int64  `json:"productId"`    //产品类型
	TemplateName string `json:"templateName"` //模板名称
	Content      string `json:"content"`      //模版内容
	Status       int    `json:"status"`       //审核状态，1审核通过，2待审核，3审核驳回
	Ctime        string `json:"ctime"`        //创建时间
	UserId       int64  `json:"userId"`       // 用户ID
	UserName     string `json:"userName"`     // 用户名称
	RefuseReason string `json:"refuseReason"`
	ChannelCmc   string `json:"channelCmc"`
	ChannelCuc   string `json:"channelCuc"`
	ChannelCtc   string `json:"channelCtc"`
	Access       string `json:"access"`
	Type         string `json:"type"`
	Remark       string `json:"remark"`
	LimitDay     string `json:"limitDay"`
	PostCount    int    `json:"postCount"`
	SendCount    int    `json:"sendCount"`
}

// TemplateStateNoticeReq 短信模版状态通知
type TemplateStateNoticeReq struct {
	Status    int    `json:"status"`    // 审核状态，1为通过，3为驳回
	ProductId int64  `json:"productId"` // 产品类型
	Id        int64  `json:"id"`        // 模板id
	Title     string `json:"title"`     // 固定值：模板审核
	Content   string `json:"content"`   // 推送信息：您的模板：id审核成功/审核失败
	Type      string `json:"type"`      // 固定值：template
	CTime     string `json:"cTime"`     // 创建时间
}
