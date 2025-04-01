package smschina

import "github.com/cyjaysong/shlianlu"

// GetReplyListReq 拉取回复列表
type GetReplyListReq struct {
	Date     string `json:"Date"`     //回复日期，格式如20230828
	PageNo   int64  `json:"pageNo"`   //第几页,从1开始
	PageSize int64  `json:"pageSize"` //每页显示多少条,最多显示100条
}

func (req *GetReplyListReq) Do(cli *shlianlu.Client) (res *GetReplyListRes, err error) {
	res = &GetReplyListRes{}
	if err = cli.Pose("/sms/trade/reply", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetReplyListRes struct {
	Message   string         `json:"message"`
	Status    string         `json:"status"`
	Timestamp int64          `json:"timestamp,string"`
	Total     int            `json:"total"`
	Data      []GetReplyItem `json:"data"`
}

type GetReplyItem struct {
	TaskId      string `json:"taskId"`
	Phone       string `json:"phone"`
	RespTime    string `json:"respTime"`
	RespContent string `json:"respContent"`
	Tag         string `json:"tag"`
}

type ReplyNoticeReq struct {
	Status      string `json:"status"`
	TaskId      string `json:"taskId"`
	Phone       string `json:"phone"`
	SequenceId  string `json:"sequenceId"`
	ContentDown string `json:"contentDown"`
	ContentUp   string `json:"contentUp"`
	Timestamp   int64  `json:"timestamp"`
}
