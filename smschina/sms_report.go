package smschina

import "github.com/cyjaysong/shlianlu-go"

// GetReportListReq 拉取报告列表
type GetReportListReq struct {
	TaskId   string `json:"TaskId"`   //任务ID
	PageNo   int64  `json:"pageNo"`   //第几页,从1开始
	PageSize int64  `json:"pageSize"` //每页显示多少条,默认10条
}

func (req *GetReportListReq) Do(cli *shlianlu.Client) (res *shlianlu.BaseRes[[]GetReportItem], err error) {
	res = &shlianlu.BaseRes[[]GetReportItem]{}
	if err = cli.Pose("/sms/trade/report", "1.1d.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetReportItem struct {
	SequenceId string `json:"sequenceId"`
	Phone      string `json:"phone"`
	Content    string `json:"content"`       // 发送短信内容
	Status     int    `json:"status,string"` // 短信发送状态,0未知,1发送成功,2发送失败
	RespTime   string `json:"respTime"`      // 短信接受毫秒时间戳,状态未知为空
	RespCode   string `json:"respCode"`      // 短信状态码,DELIVRD为发送成功,其他为发送失败,对照表见短信状态码
	Tag        string `json:"tag"`
}

type ReportNoticeReq struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	TaskId     string `json:"taskId"`
	SequenceId string `json:"sequenceId"`
	Phone      string `json:"phone"`
	RespTime   string `json:"resptime"` // 短信接受毫秒时间戳,状态未知为空
	RespCode   string `json:"respCode"`
	CodeDesc   string `json:"codeDesc"`
	Tag        string `json:"tag"`
}
