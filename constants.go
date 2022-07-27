package espn

const espnBaseURL = "https://fantasy.espn.com/apis/v3/games"

// GameType is the type of the ESPN fantasy league (currently only FFL is supported)
type GameType int32

const (
	// GameTypeNfl represents the fantasy football league type
	GameTypeNfl GameType = iota
)

func (g GameType) String() string {
	switch g {
	case GameTypeNfl:
		return "ffl"
	}
	return ""
}

// GameOutcome represents the outcome of a particular matchup
type GameOutcome int32

const (
	// GameOutcomeUndecided represents a matchup that has not yet finished
	GameOutcomeUndecided GameOutcome = iota
	// GameOutcomeWin represents a matchup that the given team won
	GameOutcomeWin
	// GameOutcomeLoss represents a matchup that the given team lost
	GameOutcomeLoss
)

var activityMap = map[int]string{
	178: "FA ADDED",
	179: "DROPPED",
	180: "WAIVER ADDED",
	181: "DROPPED",
	239: "DROPPED",
	244: "TRADED",
}

var positionMap = map[int]string{
	0:  "QB",
	1:  "TQB",
	2:  "RB",
	3:  "RB/WR",
	4:  "WR",
	5:  "WR/TE",
	6:  "TE",
	7:  "OP",
	8:  "DT",
	9:  "DE",
	10: "LB",
	11: "DL",
	12: "CB",
	13: "S",
	14: "DB",
	15: "DP",
	16: "D/ST",
	17: "K",
	18: "P",
	19: "HC",
	20: "BE",
	21: "IR",
	22: "",
	23: "RB/WR/TE",
	24: "ER",
	25: "Rookie",
}

var proTeamMap = map[int]string{
	0:  "None",
	1:  "ATL",
	2:  "BUF",
	3:  "CHI",
	4:  "CIN",
	5:  "CLE",
	6:  "DAL",
	7:  "DEN",
	8:  "DET",
	9:  "GB",
	10: "TEN",
	11: "IND",
	12: "KC",
	13: "OAK",
	14: "LAR",
	15: "MIA",
	16: "MIN",
	17: "NE",
	18: "NO",
	19: "NYG",
	20: "NYJ",
	21: "PHI",
	22: "ARI",
	23: "PIT",
	24: "LAC",
	25: "SF",
	26: "SEA",
	27: "TB",
	28: "WSH",
	29: "CAR",
	30: "JAX",
	33: "BAL",
	34: "HOU",
}
