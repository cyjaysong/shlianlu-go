package voice

import "github.com/cyjaysong/shlianlu-go"

// GetReplyListReq 拉取回复列表
type GetReplyListReq struct {
	Date     string `json:"Date"`     //回复日期，格式如20230828
	PageNo   int64  `json:"pageNo"`   //第几页,从1开始
	PageSize int64  `json:"pageSize"` //每页显示多少条,最多显示100条
}

func (req *GetReplyListReq) Do(cli *shlianlu.Client) (res *GetReplyListRes, err error) {
	res = &GetReplyListRes{}
	if err = cli.Pose("/sms/voice/reply", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetReplyListRes struct {
	Message   string         `json:"message"`
	Status    string         `json:"status"`
	Timestamp int64          `json:"timestamp"`
	Total     int            `json:"total"`
	Data      []GetReplyItem `json:"data"`
}

type GetReplyItem struct {
	TaskId      string `json:"taskId"`
	DownContent string `json:"downContent"` // 语音播放文本内容 或者 语音文件url
	Phone       string `json:"phone"`
	RespTime    string `json:"respTime"`
	RespContent string `json:"respContent"`
	Tag         string `json:"tag"`
}

type ReplyNoticeReq struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	TaskId   string `json:"taskId"`
	Phone    string `json:"phone"`
	RespTime string `json:"resptime"`
	RespCode string `json:"respCode"`
	Tag      string `json:"tag"`
	Hold     int64  `json:"hold,string"` // 通话持续时间（秒）
}
