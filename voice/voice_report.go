package voice

import "github.com/cyjaysong/shlianlu-go"

// GetReportListReq 拉取报告列表
type GetReportListReq struct {
	TaskId   string `json:"TaskId"`   //任务ID
	PageNo   int64  `json:"pageNo"`   //第几页,从1开始
	PageSize int64  `json:"pageSize"` //每页显示多少条,默认10条
}

func (req *GetReportListReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[[]GetReportItem], err error) {
	res = &shlianlu.BaseRes[[]GetReportItem]{}
	if err = cli.Pose("/sms/voice/report", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetReportItem struct {
	SequenceId string `json:"sequenceId"`    // 序列号
	Phone      string `json:"phone"`         // 号码
	Content    string `json:"content"`       // 内容
	Status     int    `json:"status,string"` // 发送状态,0未知,1发送成功,2发送失败
	RespTime   string `json:"respTime"`      // 接收时间,状态未知为空
	RespCode   string `json:"respCode"`      // 状态码,DELIVERED为发送成功,其他为发送失败,对照表见短信状态码
	//CodeDesc   string `json:"codeDesc"`
	Hold int64  `json:"hold,string"` // 计费时长(秒)
	Tag  string `json:"tag"`
}

type ReportNoticeReq struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	TaskId     string `json:"taskId"`
	SequenceId string `json:"sequenceId"`
	Phone      string `json:"phone"`    // 号码
	Resptime   string `json:"resptime"` // 接收时间,状态未知为空
	RespCode   string `json:"respCode"` // 状态码,DELIVERED为发送成功,其他为发送失败,对照表见短信状态码
	//CodeDesc   string `json:"codeDesc"`
	Hold int64  `json:"hold,string"` // 计费时长(秒)
	Tag  string `json:"tag"`
}
