package def

const (
	// 使用Forest
	P0 int64 = 1 << iota
	// 使用充值
	P1
	// 创建自己的任务
	P2
	// 查看自己的任务
	P3
	// 修改自己的任务
	P4
	// 创建所有人的任务
	P5
	// 查看所有人的任务
	P6
	// 修改所有人的任务
	P7
	// 创建注册码
	P8
	// 查看注册码
	P9
	// 修改注册码
	P10
	// 创建用户
	P11
	// 查看用户
	P12
	// 修改用户
	P13
	// 查看自己的记录
	P14
	// 修改自己的记录
	P15
	// 查看所有人的记录
	P16
	// 修改所有人的记录
	P17
	// 查看自己的交易记录
	P18
	// 修改自己的交易记录
	P19
	// 查看所有人的交易记录
	P20
	// 修改所有人的交易记录
	P21
	// 查看自己的分析
	P22
	// 查看所有人的分析
	P23
	// 使用ToDo
	P24
)

func CombinePermission(p ...int64) int64 {
	var res int64 = 0
	for _, v := range p {
		res |= v
	}
	return res
}

func CheckPermission(userPermission int64, permission ...int64) bool {
	for _, v := range permission {
		if (userPermission & v) == 0 {
			return false
		}
	}
	return true
}

func Admin() int64 {
	return CombinePermission(P5, P6, P7, P8, P9, P10, P11, P12, P13, P16, P17, P20, P21, P23)
}

func Normal() int64 {
	return CombinePermission(P0, P1, P2, P3, P4, P14, P15, P18, P19, P22)
}
