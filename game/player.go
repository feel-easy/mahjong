package game

type Player interface {
	PlayerID() int
	NickName() string
	Play(tiles []int, gameState State) (int, error)
	Take(tiles []int, gameState State) (int, []int, error)
}
