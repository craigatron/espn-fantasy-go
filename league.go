package espn

import "fmt"

type League struct {
	GameType GameType
	LeagueId string
	Year     int

	CurrentMatchupPeriod int
	ScoringPeriodId      int
	FirstScoringPeriod   int
	CurrentWeek          int
	Settings             LeagueSettingsJson
	Members              map[string]*LeagueMemberJson
	Teams                map[int]*Team
	Schedule             [][]Matchup

	client *espnClient
}

type Team struct {
	Id           int
	Abbreviation string
	Name         string
	Owners       []string
	Schedule     []*Team
	Scores       []float64
	Outcomes     []GameOutcome
}

type Matchup struct {
	MatchupId int
	HomeTeam  *Team
	HomeScore float64
	AwayTeam  *Team
	AwayScore float64
	Winner    string
}

func NewPublicLeague(gameType GameType, leagueId string, year int) (League, error) {
	return newLeagueInternal(gameType, leagueId, year, newPublicClient(gameType, leagueId, year))
}

func NewPrivateLeague(gameType GameType, leagueId string, year int, espnS2 string, swid string) (League, error) {
	return newLeagueInternal(gameType, leagueId, year, newPrivateClient(gameType, leagueId, year, espnS2, swid))
}

func newLeagueInternal(gameType GameType, leagueId string, year int, client *espnClient) (League, error) {
	league := League{
		GameType: gameType,
		LeagueId: leagueId,
		Year:     year,

		client: client,
	}

	err := league.RefreshData()

	if err != nil {
		return league, err
	}

	return league, nil
}

func (league League) GetLeague() (LeagueInfoResponseJson, error) {
	res := LeagueInfoResponseJson{}
	err := league.client.getLeagueInternal([]string{"mTeam", "mRoster", "mMatchup", "mSettings", "mStandings"}, "", &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (league *League) RefreshData() error {
	leagueInfo, err := league.GetLeague()
	if err != nil {
		return err
	}

	league.CurrentMatchupPeriod = leagueInfo.Status.CurrentMatchupPeriod
	league.ScoringPeriodId = leagueInfo.ScoringPeriodID
	league.FirstScoringPeriod = leagueInfo.Status.FirstScoringPeriod
	if league.Year < 2018 {
		league.CurrentWeek = leagueInfo.ScoringPeriodID
	} else if leagueInfo.ScoringPeriodID <= leagueInfo.Status.FinalScoringPeriod {
		league.CurrentWeek = leagueInfo.ScoringPeriodID
	} else {
		league.CurrentWeek = leagueInfo.Status.FinalScoringPeriod
	}
	league.Settings = leagueInfo.Settings

	membersMap := make(map[string]*LeagueMemberJson)
	for _, m := range leagueInfo.Members {
		membersMap[m.ID] = &m
	}
	league.Members = membersMap

	teams := make(map[int]*Team)
	for _, t := range leagueInfo.Teams {
		team := Team{
			Id:           t.ID,
			Abbreviation: t.Abbrev,
			Name:         fmt.Sprintf("%s %s", t.Location, t.Nickname),
			Owners:       t.Owners,
			Schedule:     make([]*Team, 0),
			Scores:       make([]float64, 0),
			Outcomes:     make([]GameOutcome, 0),
		}
		// TODO: populate roster, other data
		teams[team.Id] = &team
	}

	league.Teams = teams

	scheduleByWeek := make([][]Matchup, 0, leagueInfo.Status.FinalScoringPeriod)
	for i := 1; i <= leagueInfo.Status.FinalScoringPeriod; i++ {
		scheduleByWeek = append(scheduleByWeek, make([]Matchup, 0, len(league.Teams)/2))
	}
	for _, matchup := range leagueInfo.Schedule {
		lm := league.convertMatchupJson(matchup)
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
				game.HomeTeam.Outcomes = append(game.HomeTeam.Outcomes, GameOutocomeUndecided)
				// bye weeks seem to be represented as an undecided game with only a home team
				if game.AwayTeam != nil {
					game.AwayTeam.Outcomes = append(game.AwayTeam.Outcomes, GameOutocomeUndecided)
				}
			} else {
				panic(fmt.Errorf("Unknown game winner type: %s", game.Winner))
			}
		}
	}

	return nil
}

func (league League) convertMatchupJson(lm LeagueMatchupJson) Matchup {
	return Matchup{
		MatchupId: lm.ID,
		HomeTeam:  league.Teams[lm.Home.TeamID],
		HomeScore: lm.Home.TotalPoints,
		AwayTeam:  league.Teams[lm.Away.TeamID],
		AwayScore: lm.Away.TotalPoints,
		Winner:    lm.Winner,
	}
}

func (league League) Matchups(week int) []Matchup {
	return league.Schedule[week]
}

func (league League) CurrentWeekMatchups() []Matchup {
	return league.Schedule[league.CurrentWeek-1]
}

func (league League) Scoreboard() ([]Matchup, error) {
	res := LeagueInfoResponseJson{}

	filter := fmt.Sprintf("{\"schedule\":{\"filterMatchupPeriodIds\":{\"value\":[%d]}}}", league.CurrentWeek)
	err := league.client.getLeagueInternal([]string{"mScoreboard"}, filter, &res)

	if err != nil {
		return nil, err
	}

	matchups := make([]Matchup, 0, len(res.Schedule))
	for _, m := range res.Schedule {
		matchups = append(matchups, league.convertMatchupJson(m))
	}
	return matchups, nil
}
