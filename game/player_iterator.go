package game

type PlayerIterator struct {
	players map[int]*playerController
	cycler  *Cycler
}

func (i *PlayerIterator) GetPlayerController(id int) *playerController {
	return i.players[id]
}

func newPlayerIterator(players []Player) *PlayerIterator {
	var playerIDs []int
	playerMap := make(map[int]*playerController, len(players))
	for _, player := range players {
		playerID := player.PlayerID()
		playerIDs = append(playerIDs, playerID)
		playerMap[playerID] = newPlayerController(player)
	}
	return &PlayerIterator{
		players: playerMap,
		cycler:  NewCycler(playerIDs),
	}
}

func (i *PlayerIterator) Current() *playerController {
	return i.players[i.cycler.Current()]
}

func (i *PlayerIterator) ForEach(function func(player *playerController)) {
	for range i.players {
		function(i.Current())
		i.Next()
	}
}

func (i *PlayerIterator) Next() *playerController {
	return i.players[i.cycler.Next()]
}
