package dam

import "encoding/json"

type Inc struct {
	Name string // 表名
	Val  int    // 自增值
}

type User struct {
	Uuid       int    `json:"uuid"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	Permission int64  `json:"permission"`
	Plant      int    `json:"plant"`
	RmbIn      int64  `json:"rmb_in"`  // 累计充值金额
	RmbOut     int64  `json:"rmb_out"` // 累计充值金额
	Coin       int64  `json:"coin"`    // 当前硬币数量
	Deleted    int64  `json:"deleted"` // 删除的时间，0代表帐号状态正常
}

type Work struct {
	Wuid int    `json:"wuid"`
	Uuid int    `json:"uuid"`
	Name string `json:"name"`
	Type int `json:"type"`
	// i_type 的计量基准（例如多少分钟或多少次）
	// i_type 为 自动计算收益 时代表自动执行的时间间隔
	Unit int `json:"unit"`
	// 硬币数量
	Coin int64 `json:"coin"`
	// 关联的自动执行的任务的wuid
	Associate int `json:"associate"`
	// 删除的时间，0代表帐号状态正常
	Deleted int64 `json:"deleted"`
}

type Record struct {
	Urid int `json:"urid"`
	Uuid int `json:"uuid"`
	Wuid int `json:"wuid"`
	// 开始时间戳
	TimeStart int64 `json:"time_start"`
	// 结束时间戳
	TimeEnd int64 `json:"time_end"`
	// 获得硬币数量
	Coin int64 `json:"coin"`
	// 备注标签
	Tag []string `json:"tag"`
}

func (r * Record)Transfer() *RecordDam {
	tag, _ := json.Marshal(r.Tag)
	record := RecordDam{
		Urid:      r.Urid,
		Uuid:      r.Uuid,
		Wuid:      r.Wuid,
		TimeStart: r.TimeStart,
		TimeEnd:   r.TimeEnd,
		Coin:      r.Coin,
		Tag:       string(tag),
	}
	return &record
}

type RecordDam struct {
	Urid int
	Uuid int
	Wuid int
	TimeStart int64
	TimeEnd int64
	Coin int64
	Tag string
}

func (r * RecordDam)Transfer() *Record {
	record := Record{
		Urid:      r.Urid,
		Uuid:      r.Uuid,
		Wuid:      r.Wuid,
		TimeStart: r.TimeStart,
		TimeEnd:   r.TimeEnd,
		Coin:      r.Coin,
		Tag:       make([]string, 0),
	}
	_ = json.Unmarshal([]byte(r.Tag), &record.Tag)
	return &record
}

type RecordWithWork struct {
	Record Record `json:"record"`
	Work   Work   `json:"work"`
}

type Transaction struct {
	Utid int `json:"utid"`
	Uuid int `json:"uuid"`
	Wuid int `json:"wuid"`
	// 此条记录的时间戳
	UnixTime int64 `json:"unix_time"`
	// 人民币数量
	Rmb int64 `json:"rmb"`
	// 获得的硬币数量
	Coin int64 `json:"coin"`
}

type PermissionCode struct {
	Pcid        int    `json:"pcid"`
	Code        string `json:"code"`
	Permission  int64  `json:"permission"`
	Deadline    int64  `json:"deadline"`
	GeneratedBy int    `json:"generated_by"`
	UsedBy      int    `json:"used_by"`
	Deleted     int64  `json:"deleted"`
}
