package game

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feel-easy/mahjong/tile"
	"github.com/feel-easy/mahjong/ting"
)

type State struct {
	LastPlayer        *playerController
	OriginallyPlayer  *playerController
	CurrentPlayer     *playerController
	LastPlayedTile    int
	PlayedTiles       []int
	CurrentPlayerHand []int
	PlayerSequence    []*playerController
	PlayerShowCards   map[string][]*ShowCard
	SpecialPrivileges map[int][]int
	CanWin            []*playerController
}

func (s State) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("playedTiles:%s", tile.ToTileString(s.PlayedTiles)))
	var playerStatuses []string

	for _, player := range s.PlayerSequence {
		playerStatus := fmt.Sprintf("%s:", player.Name())
		if canTing, _ := ting.CanTing(player.Hand(), player.GetShowCardTiles()); canTing {
			playerStatus += "(听)"
		}
		if showCards, ok := s.PlayerShowCards[player.Name()]; ok && len(showCards) > 0 {
			for _, showCard := range showCards {
				playerStatus += fmt.Sprintf("%s ", showCard.String())
			}
		}
		playerStatuses = append(playerStatuses, playerStatus)
	}
	drew := s.CurrentPlayerHand[len(s.CurrentPlayerHand)-1]
	sort.Ints(s.CurrentPlayerHand)
	if s.LastPlayer != nil {
		lines = append(lines, fmt.Sprintf("ShowCards:\n%s ", strings.Join(playerStatuses, "\n")))
		lines = append(lines, fmt.Sprintf("%s played: %s", s.LastPlayer.Name(), tile.Tile(s.LastPlayedTile).String()))
	}
	lines = append(lines, fmt.Sprintf("Your drew: %s ", tile.Tile(drew)))
	lines = append(lines, fmt.Sprintf("Your hand: %s \n", tile.ToTileString(s.CurrentPlayerHand)))
	return strings.Join(lines, "\n")
}
