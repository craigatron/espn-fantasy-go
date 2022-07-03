package espn

const EspnBaseUrl = "https://fantasy.espn.com/apis/v3/games"

type GameType int32

const (
	GameTypeNfl GameType = iota
)

func (g GameType) String() string {
	switch g {
	case GameTypeNfl:
		return "ffl"
	}
	return ""
}

type GameOutcome int32

const (
	GameOutocomeUndecided GameOutcome = iota
	GameOutcomeWin
	GameOutcomeLoss
)
