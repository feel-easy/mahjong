package game

import (
	"github.com/feel-easy/mahjong/card"
	"github.com/feel-easy/mahjong/consts"
	"github.com/feel-easy/mahjong/tile"
	"github.com/feel-easy/mahjong/win"
)

type Game struct {
	players *PlayerIterator
	deck    *Deck
	pile    *Pile
}

func (g *Game) Players() *PlayerIterator {
	return g.players
}

func (g *Game) Deck() *Deck {
	return g.deck
}

func (g *Game) Pile() *Pile {
	return g.pile
}

func (g *Game) Next() *playerController {
	player := g.Players().Next()
	g.pile.SetCurrentPlayer(player)
	return player
}

func New(players []Player) *Game {
	return &Game{
		players: newPlayerIterator(players),
		deck:    NewDeck(),
		pile:    NewPile(),
	}
}

func (g *Game) GetPlayerTiles(id int) string {
	tiles := g.players.GetPlayerController(id).Hand()
	return tile.ToTileString(tiles)
}

func (g *Game) DealStartingTiles() {
	g.players.ForEach(func(player *playerController) {
		hand := g.deck.Draw(13)
		player.AddTiles(hand)
	})
}

func (g *Game) Current() *playerController {
	return g.players.Current()
}

func (g Game) ExtractState(player *playerController) State {
	playerSequence := make([]*playerController, 0)
	playerShowCards := make(map[string][]*ShowCard)
	specialPrivileges := make(map[int][]int)
	canWin := make([]*playerController, 0)
	originallyPlayer := g.pile.originallyPlayer
	topTile := g.pile.Top()
	g.players.ForEach(func(player *playerController) {
		playerSequence = append(playerSequence, player)
		playerShowCards[player.Name()] = player.GetShowCard()
		if _, ok := g.pile.SayNoPlayer()[player.ID()]; !ok &&
			topTile > 0 && g.pile.lastPlayer.ID() != player.ID() {
			if win.CanWin(append(player.Hand(), g.pile.Top()), player.GetShowCardTiles()) {
				canWin = append(canWin, player)
			}
			if card.CanGang(player.Hand(), topTile) {
				specialPrivileges[player.ID()] = append(specialPrivileges[player.ID()], consts.GANG)
			}
			if card.CanPeng(player.Hand(), topTile) {
				specialPrivileges[player.ID()] = append(specialPrivileges[player.ID()], consts.PENG)
			}
			if originallyPlayer.ID() == player.ID() &&
				card.CanChi(player.Hand(), topTile) {
				specialPrivileges[player.ID()] = append(specialPrivileges[player.ID()], consts.CHI)
			}
		}
	})
	return State{
		LastPlayer:        g.pile.lastPlayer,
		OriginallyPlayer:  originallyPlayer,
		LastPlayedTile:    g.pile.Top(),
		PlayedTiles:       g.pile.Tiles(),
		CurrentPlayerHand: player.Tiles(),
		CurrentPlayer:     g.Current(),
		PlayerSequence:    playerSequence,
		PlayerShowCards:   playerShowCards,
		SpecialPrivileges: specialPrivileges,
		CanWin:            canWin,
	}
}
