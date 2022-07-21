package consts

import "time"

const (
	_ int = iota
	CHI
	PENG
	GANG
	WIN
)
const PlayMahjongTimeout = 30 * time.Second

var OpCodeData = map[int]string{
	CHI:  "吃",
	PENG: "碰",
	GANG: "杠",
	WIN:  "胡",
}
