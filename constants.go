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
