// 1. Define exchange ratio
// 2. Provide exchange function
// 3. Provide money and virtual currency exchange functions
package def

const (
	Copper int64 = 1
	Silver       = Copper * 100
	Gold         = Silver * 100
)

// 将三种硬币全部兑换成铜币
func CoinMerge(gold, silver, copper int64) int64 {
	return gold*Gold + silver*Silver + copper*Copper
}

// 将铜币兑换成三种硬币的组合
func CoinSplit(c int64) (gold, silver, copper int64) {
	gold = c / Gold
	c -= gold * Gold
	silver = c / Silver
	c -= silver * Silver
	copper = c
	return
}
