package voice

import "github.com/cyjaysong/shlianlu-go"

// GetBalanceReq 产品余额查询
type GetBalanceReq struct{}

func (req *GetBalanceReq) Do(cli *shlianlu.Client) (res *GetBalanceRes, err error) {
	res = &GetBalanceRes{}
	if err = cli.Pose("/sms/product/balance", "1.1.0", req, res); err != nil {
		return nil, err
	}
	return
}

type GetBalanceRes struct {
	Message   string `json:"message"`   // 消息
	Status    string `json:"status"`    // 状态
	Timestamp int64  `json:"timestamp"` // 时间戳
	Balance   int64  `json:"balance"`   // 当前账户余额(厘)
}
