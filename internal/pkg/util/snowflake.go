package util

import (
	"github.com/bwmarrin/snowflake"
)

// MakeSnowflakeID snowflake算法生成一个int64类型的id
// 符号位(1bit) + 毫秒级时间(41bit) + 机器id(10bit) + 序列(12bit)
// 0 00000000000000000000000000000000000000000 0000000000 000000000000
// 1 + 41 + 10 + 12 = 64
func MakeSnowflakeID() (int64, error) {
	node, err := snowflake.NewNode(0) // 10bit，机器id，0~1023
	if err != nil {
		return 0, err
	}
	id := node.Generate()
	return id.Int64(), nil
}
