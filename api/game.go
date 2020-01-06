package api

import (
	"encoding/json"
	"sort"
)

var client Client

// Game contains information relevant to an NFL game.
// Contains game information, player stats, team stats,
// etc...
type Game struct {
	GameID       string `json:"id"`
	Home         Team   `json:"home"`
	Away         Team   `json:"away"`
	Down         int    `json:"down"`
	ToGo         int    `json:"togo"`
	Clock        string `json:"clock"`
	Pos          string `json:"posteam"`
	Redzone      bool   `json:"redzone"`
	Yl           string `json:"yl"`
	Qtr          string `json:"qtr"`
	ScoreSummary map[string]struct {
		Type string `json:"type"`
		Desc string `json:"desc"`
		Qtr  int    `json:"qtr"`
		Team string `json:"team"`
	} `json:"scrsummary"`
}

// Team represents an NFL team playing within an
// NFL game. Contains team stats and information.
type Team struct {
	Abbr  string `json:"abbr"`
	Score struct {
		Q1    int `json:"1"`
		Q2    int `json:"2"`
		Q3    int `json:"3"`
		Q4    int `json:"4"`
		OT    int `json:"5"`
		Final int `json:"t"`
	} `json:"score"`
	Timeouts int `json:"to"`
	Stats    struct {
		Passing   map[string]Player `json:"passing"`
		Rushing   map[string]Player `json:"rushing"`
		Receiving map[string]Player `json:"receiving"`
		Defense   map[string]Player `json:"defense:`
		Team      struct {
			FirstDowns int    `json:"totfd"`
			TotalYards int    `json:"totyds"`
			PassYards  int    `json:"pyds"`
			RushYards  int    `json:"ryds"`
			Turnovers  int    `json:"trnovr"`
			Top        string `json:"top"`
		} `json:"team"`
	} `json:"stats"`
}

// Player represents an NFL player playing in an
// NFL game. Contains player stats and information.
type Player struct {
	Name      string `json:"name"`
	Att       int    `json:"att"`
	Comp      int    `json:"cmp"`
	Rec       int    `json:"rec"`
	Yards     int    `json:"yds"`
	Tds       int    `json:"tds"`
	IntsO     int    `json:"ints"`
	TwoPtAtt  int    `json:"twopta"`
	TwoPtConv int    `json:"twoptm"`
	Tkl       int    `json:"tkl"`
	Ast       int    `json:"ast"`
	Sacks     int    `json:"sck"`
	IntsD     int    `json:"int"`
	Ffum      int    `json:"ffum"`
}

// GetGame retrieves a specific NFL game that is in progress,
// or has ended.
//
// The ID for the game is structured like so:
//
// 2020 (year) +  01 (month) + 04 (day) + 01 (game number 00-16)
// = 2020010401
func GetGame(id string) *Game {
	apiURL := "http://www.nfl.com/liveupdate/game-center/" + id + "/" + id + "_gtd.json"

	data, err := client.Get(apiURL)
	if err != nil {
		panic(err)
	}

	rawGame := make(map[string]interface{})
	if err := json.Unmarshal(data, &rawGame); err != nil {
		panic(err)
	}

	rawGame, _ = rawGame[id].(map[string]interface{})
	rawGame["id"] = id
	gameData, _ := json.Marshal(rawGame)

	game := &Game{}
	if err := json.Unmarshal(gameData, game); err != nil {
		panic(err)
	}

	return game
}

// Games retrieves a general list of all upcoming/ongoing games.
// Games does not provide indepth information/stats and is mostly
// used to retrieve Game IDs
func Games() []*Game {
	apiURL := "http://www.nfl.com/liveupdate/scores/scores.json"

	data, err := client.Get(apiURL)
	if err != nil {
		panic(err)
	}

	rawGames := make(map[string]interface{})
	if err := json.Unmarshal(data, &rawGames); err != nil {
		panic(err)
	}

	games := []*Game{}
	for k, v := range rawGames {
		v, _ := v.(map[string]interface{})
		v["id"] = k
		gameData, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		game := &Game{}
		if err := json.Unmarshal(gameData, game); err != nil {
			panic(err)
		}
		games = append(games, game)
	}

	sort.Slice(games[:], func(i, j int) bool {
		if []rune(games[i].Home.Abbr)[0] == []rune(games[j].Home.Abbr)[0] {
			return []rune(games[i].Home.Abbr)[1] < []rune(games[j].Home.Abbr)[1]
		}
		return []rune(games[i].Home.Abbr)[0] < []rune(games[j].Home.Abbr)[0]
	})

	return games
}

// Update updates a specific Game instance with up to date
// stats and information.
func (g *Game) Update() {
	apiURL := "http://www.nfl.com/liveupdate/game-center/" + g.GameID + "/" + g.GameID + "_gtd.json"

	data, err := client.Get(apiURL)
	if err != nil {
		panic(err)
	}

	rawGame := make(map[string]interface{})
	json.Unmarshal(data, &rawGame)

	if rawGame[g.GameID] == nil {
		return
	}

	rawGame = rawGame[g.GameID].(map[string]interface{})
	gameData, err := json.Marshal(rawGame)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(gameData, g)
}
