package espn

import "fmt"

// League is an ESPN fantasy league.
type League struct {
	GameType GameType
	ID       string
	Year     int

	CurrentMatchupPeriod int
	ScoringPeriodID      int
	FirstScoringPeriod   int
	CurrentWeek          int
	Settings             LeagueSettingsJSON
	Members              map[string]*LeagueMemberJSON
	Teams                map[int64]*Team
	Schedule             [][]Matchup

	Players map[int64]Player

	client *espnClient
}

// Team is information about a team within an ESPN fantasy league.
type Team struct {
	ID           int64
	Abbreviation string
	Name         string
	Owners       []string
	Schedule     []*Team
	Scores       []float64
	Outcomes     []GameOutcome
}

// Matchup is a single ESPN fantasy game.
type Matchup struct {
	ID        int64
	HomeTeam  *Team
	HomeScore float64
	AwayTeam  *Team
	AwayScore float64
	Winner    string
}

// Player is information about a player (heh)
type Player struct {
	ID       int64
	Position string
	FullName string
	Team     string
	PctOwned float64
}

// NewPublicLeague creates and initializes a public ESPN league.
func NewPublicLeague(gameType GameType, leagueID string, year int) (League, error) {
	return newLeagueInternal(gameType, leagueID, year, newPublicClient(gameType, leagueID, year))
}

// NewPrivateLeague creates and initializes a private ESPN league (using a user's espn_s2 and SWID cookie values).
func NewPrivateLeague(gameType GameType, leagueID string, year int, espnS2 string, swid string) (League, error) {
	return newLeagueInternal(gameType, leagueID, year, newPrivateClient(gameType, leagueID, year, espnS2, swid))
}

func newLeagueInternal(gameType GameType, leagueID string, year int, client *espnClient) (League, error) {
	league := League{
		GameType: gameType,
		ID:       leagueID,
		Year:     year,

		client: client,
	}

	err := league.RefreshData()

	if err != nil {
		return league, err
	}

	return league, nil
}

// GetLeague gets a full set of league information from the ESPN API.
func (league League) GetLeague() (LeagueInfoResponseJSON, error) {
	res := LeagueInfoResponseJSON{}
	err := league.client.getLeagueInternal([]string{"mTeam", "mRoster", "mMatchup", "mSettings", "mStandings"}, "", "", &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// RefreshData updates the League struct with the latest data from the ESPN API.
func (league *League) RefreshData() error {
	leagueInfo, err := league.GetLeague()
	if err != nil {
		return err
	}

	league.CurrentMatchupPeriod = leagueInfo.Status.CurrentMatchupPeriod
	league.ScoringPeriodID = leagueInfo.ScoringPeriodID
	league.FirstScoringPeriod = leagueInfo.Status.FirstScoringPeriod
	if league.Year < 2018 {
		league.CurrentWeek = leagueInfo.ScoringPeriodID
	} else if leagueInfo.ScoringPeriodID <= leagueInfo.Status.FinalScoringPeriod {
		league.CurrentWeek = leagueInfo.ScoringPeriodID
	} else {
		league.CurrentWeek = leagueInfo.Status.FinalScoringPeriod
	}
	league.Settings = leagueInfo.Settings

	membersMap := make(map[string]*LeagueMemberJSON)
	for _, m := range leagueInfo.Members {
		membersMap[m.ID] = &m
	}
	league.Members = membersMap

	teams := make(map[int64]*Team)
	for _, t := range leagueInfo.Teams {
		team := Team{
			ID:           t.ID,
			Abbreviation: t.Abbrev,
			Name:         fmt.Sprintf("%s %s", t.Location, t.Nickname),
			Owners:       t.Owners,
			Schedule:     make([]*Team, 0),
			Scores:       make([]float64, 0),
			Outcomes:     make([]GameOutcome, 0),
		}
		// TODO: populate roster, other data
		teams[team.ID] = &team
	}

	league.Teams = teams

	scheduleByWeek := make([][]Matchup, 0, leagueInfo.Status.FinalScoringPeriod)
	for i := 1; i <= leagueInfo.Status.FinalScoringPeriod; i++ {
		scheduleByWeek = append(scheduleByWeek, make([]Matchup, 0, len(league.Teams)/2))
	}
	for _, matchup := range leagueInfo.Schedule {
		lm := league.convertMatchupJSON(matchup)
		weekGames := scheduleByWeek[matchup.MatchupPeriodID-1]
		scheduleByWeek[matchup.MatchupPeriodID-1] = append(weekGames, lm)
	}
	league.Schedule = scheduleByWeek

	for _, week := range scheduleByWeek {
		for _, game := range week {
			game.HomeTeam.Schedule = append(game.HomeTeam.Schedule, game.AwayTeam)
			game.HomeTeam.Scores = append(game.HomeTeam.Scores, game.HomeScore)

			if game.AwayTeam != nil {
				game.AwayTeam.Schedule = append(game.AwayTeam.Schedule, game.HomeTeam)
				game.AwayTeam.Scores = append(game.AwayTeam.Scores, game.AwayScore)
			}

			if game.Winner == "HOME" {
				game.HomeTeam.Outcomes = append(game.HomeTeam.Outcomes, GameOutcomeWin)
				game.AwayTeam.Outcomes = append(game.AwayTeam.Outcomes, GameOutcomeLoss)
			} else if game.Winner == "AWAY" {
				game.HomeTeam.Outcomes = append(game.HomeTeam.Outcomes, GameOutcomeLoss)
				game.AwayTeam.Outcomes = append(game.AwayTeam.Outcomes, GameOutcomeWin)
			} else if game.Winner == "UNDECIDED" {
				game.HomeTeam.Outcomes = append(game.HomeTeam.Outcomes, GameOutcomeUndecided)
				// bye weeks seem to be represented as an undecided game with only a home team
				if game.AwayTeam != nil {
					game.AwayTeam.Outcomes = append(game.AwayTeam.Outcomes, GameOutcomeUndecided)
				}
			} else {
				panic(fmt.Errorf("unknown game winner type: %s", game.Winner))
			}
		}
	}

	players, err := league.GetActivePlayers()
	if err != nil {
		return err
	}
	playerMap := make(map[int64]Player)
	for _, p := range players {
		playerMap[p.Player.ID] = Player{
			ID:       p.Player.ID,
			Position: positionMap[p.Player.DefaultPositionID],
			FullName: p.Player.FullName,
			Team:     proTeamMap[p.Player.ProTeamID],
			PctOwned: p.Player.Ownership.PercentOwned,
		}
	}
	league.Players = playerMap

	return nil
}

func (league League) convertMatchupJSON(lm LeagueMatchupJSON) Matchup {
	return Matchup{
		ID:        lm.ID,
		HomeTeam:  league.Teams[lm.Home.TeamID],
		HomeScore: lm.Home.TotalProjectedPointsLive,
		AwayTeam:  league.Teams[lm.Away.TeamID],
		AwayScore: lm.Away.TotalProjectedPointsLive,
		Winner:    lm.Winner,
	}
}

// Scoreboard fetches matchups and scores for the current week from ESPN.
func (league League) Scoreboard() ([]Matchup, error) {
	res := LeagueInfoResponseJSON{}

	filter := fmt.Sprintf("{\"schedule\":{\"filterMatchupPeriodIds\":{\"value\":[%d]}}}", league.CurrentWeek)
	err := league.client.getLeagueInternal([]string{"mScoreboard"}, filter, "", &res)
	if err != nil {
		return nil, err
	}

	matchups := make([]Matchup, 0, len(res.Schedule))
	for _, m := range res.Schedule {
		matchups = append(matchups, league.convertMatchupJSON(m))
	}
	return matchups, nil
}

// RecentAction is an action involving a single player as part of a RecentActivity.
type RecentAction struct {
	Team   int64
	Action string
	Player int64
}

// RecentActivity is a set of player transactions.
type RecentActivity struct {
	Actions   []RecentAction
	Timestamp int64
	ESPNID    string
}

// RecentActivity returns recent player transactions.
func (league League) RecentActivity(count int, offset int) ([]RecentActivity, error) {
	res := ActivityJSON{}
	filter := fmt.Sprintf("{\"topics\":{\"filterType\":{\"value\":[\"ACTIVITY_TRANSACTIONS\"]},\"limit\":%d,\"limitPerMessageSet\":{\"value\":%d},\"offset\":%d,\"sortMessageDate\":{\"sortPriority\":1,\"sortAsc\":false},\"sortFor\":{\"sortPriority\":2,\"sortAsc\":false},\"filterIncludeMessageTypeIds\":{\"value\": [178,180,179,239,181,244]}}}", count, count, offset)

	err := league.client.getLeagueInternal([]string{"kona_league_communication"}, filter, "/communication", &res)

	activity := make([]RecentActivity, 0)
	for _, t := range res.Topics {
		actions := make([]RecentAction, 0)
		for _, m := range t.Messages {
			var team int64
			if m.MessageTypeID == 244 { // TRADED
				team = m.From
			} else if m.MessageTypeID == 239 { // DROPPED
				team = m.For
			} else { // FA ADDED, WAIVER ADDED
				team = m.To
			}
			// TODO: look up player details
			actions = append(actions, RecentAction{
				Team:   team,
				Action: activityMap[m.MessageTypeID],
				Player: m.TargetID,
			})
		}
		activity = append(activity, RecentActivity{
			Actions:   actions,
			Timestamp: t.Date,
			ESPNID:    t.ID,
		})
	}

	return activity, err
}

// GetActivePlayers returns the list of active players for this league.
func (league League) GetActivePlayers() (PlayersJSON, error) {
	res := PlayersJSON{}

	filter := "{\"filterActive\": {\"value\": true}}"
	err := league.client.getLeagueInternal([]string{"players_wl"}, filter, "/players", &res)

	return res, err
}
