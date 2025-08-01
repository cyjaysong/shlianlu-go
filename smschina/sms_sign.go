package smschina

import (
	"github.com/cyjaysong/shlianlu-go"
)

// CreateSignReq 创建短信签名
type CreateSignReq struct {
	Content        string `json:"content"`        // 短信签名内容,不能包含表情和符号
	Type           int    `json:"type"`           // 签名来源 1.本公司（签名与本账户认证主体有关联）2.他公司（签名与本公司认证主体无关联）
	Remark         string `json:"remark"`         // 详细描述签名使用场景:营业执照,公司简称,APP名称,公众号,商标名称,软件著作权(软著),小程序
	CreditCodeUrl  string `json:"creditCodeUrl"`  // 媒体文件url链接,英文逗号分割,最多5张图片,上传相关资质：营业执照,授权书,软著,小程序相关截图等证明材料(建议营业执照放首位), type值为2时必填
	IdCardFront    string `json:"idCardFront"`    // 传媒体文件url链接,上传经办人身份证正面图片,type值为2时必填
	IdCardBack     string `json:"idCardBack"`     // 传媒体文件url链接,上传经办人身份证反面图片,type值为2时必填
	Company        string `json:"company"`        // 公司名称,type值为2时必填
	LegalPerson    string `json:"legalPerson"`    // 法人姓名,type值为2时必填
	CreditCode     string `json:"creditCode"`     // 信用代码,type值为2时必填
	CreditUserName string `json:"creditUserName"` // 经办人姓名,type值为2时必填
	IdCard         string `json:"idCard"`         // 经办人身份证号码,type值为2时必填
	Phone          string `json:"phone"`          // 经办人手机号,type值为2时必填
}

func (req *CreateSignReq) Do(cli *shlianlu.Client) (res *CreateSignRes, err error) {
	res = &CreateSignRes{}
	if err = cli.Pose("/sms/product/sign/create", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type CreateSignRes struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	SignId    int64  `json:"SignId"`
}

// GetSignListReq 查询短信签名列表
type GetSignListReq struct{}

func (req *GetSignListReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[[]GetSignItem], err error) {
	res = &shlianlu.BaseRes[[]GetSignItem]{}
	if err = cli.Pose("/sms/product/sign/get", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetSignItem struct {
	SignId        int64  `json:"signId"`
	UserId        int64  `json:"userId"`
	UserName      string `json:"userName"`
	ProductId     string `json:"productId"` // 产品类型
	SignName      string `json:"signName"`
	Content       string `json:"content"`
	Status        int    `json:"status"`
	Ctime         string `json:"ctime"`
	SignFlag      int    `json:"signFlag"`
	SignPostCount int    `json:"signPostCount"`
	SignSendCount int    `json:"signSendCount"`
	IsDefault     int    `json:"isDefault"`
	RefuseReason  string `json:"refuseReason"`
	ExtCode       string `json:"extCode"`
	Remark        string `json:"remark"`
	Ext1          string `json:"ext1"`
}

// DeleteSignReq 删除短信签名
type DeleteSignReq struct {
	SignId int64 `json:"SignId"` // 短信签名Id
}

func (req *DeleteSignReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[struct{}], err error) {
	res = &shlianlu.BaseRes[struct{}]{}
	if err = cli.Pose("/sms/product/sign/delete", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

// SignStateNoticeReq 短信签名状态通知
type SignStateNoticeReq struct {
	ProductId    string `json:"productId"`    // 产品类型
	Status       int    `json:"status"`       // 审核状态，1为通过，3为驳回
	Id           int64  `json:"id"`           // 签名Id
	Title        string `json:"title"`        // 固定值：签名审核
	Content      string `json:"content"`      // 推送信息：您的签名：id审核成功/审核失败
	Type         string `json:"type"`         // 固定值：sign
	CTime        string `json:"cTime"`        // 创建时间
	RefuseReason string `json:"refuseReason"` // 审核拒绝原因描述
}
