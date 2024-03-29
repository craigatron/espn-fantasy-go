package espn

// CumulativeScoreJSON is raw cumulative score data for a player from the ESPN API.
type CumulativeScoreJSON struct {
	Losses      int         `json:"losses"`
	ScoreByStat interface{} `json:"scoreByStat"`
	StatBySlot  interface{} `json:"statBySlot"`
	Ties        int         `json:"ties"`
	Wins        int         `json:"wins"`
}

// LeagueSettingsJSON is the raw settings JSON from the ESPN API.
type LeagueSettingsJSON struct {
	AcquisitionSettings struct {
		AcquisitionBudget            int      `json:"acquisitionBudget"`
		AcquisitionLimit             int      `json:"acquisitionLimit"`
		AcquisitionType              string   `json:"acquisitionType"`
		IsUsingAcquisitionBudget     bool     `json:"isUsingAcquisitionBudget"`
		MatchupAcquisitionLimit      float64  `json:"matchupAcquisitionLimit"`
		MatchupLimitPerScoringPeriod bool     `json:"matchupLimitPerScoringPeriod"`
		MinimumBid                   int      `json:"minimumBid"`
		WaiverHours                  int      `json:"waiverHours"`
		WaiverOrderReset             bool     `json:"waiverOrderReset"`
		WaiverProcessDays            []string `json:"waiverProcessDays"`
		WaiverProcessHour            int      `json:"waiverProcessHour"`
	} `json:"acquisitionSettings"`
	DraftSettings struct {
		AuctionBudget     int    `json:"auctionBudget"`
		IsTradingEnabled  bool   `json:"isTradingEnabled"`
		KeeperCount       int    `json:"keeperCount"`
		KeeperCountFuture int    `json:"keeperCountFuture"`
		KeeperOrderType   string `json:"keeperOrderType"`
		LeagueSubType     string `json:"leagueSubType"`
		OrderType         string `json:"orderType"`
		PickOrder         []int  `json:"pickOrder"`
		TimePerSelection  int    `json:"timePerSelection"`
		Type              string `json:"type"`
	} `json:"draftSettings"`
	FinanceSettings struct {
		EntryFee           float64 `json:"entryFee"`
		MiscFee            float64 `json:"miscFee"`
		PerLoss            float64 `json:"perLoss"`
		PerTrade           float64 `json:"perTrade"`
		PlayerAcquisition  float64 `json:"playerAcquisition"`
		PlayerDrop         float64 `json:"playerDrop"`
		PlayerMoveToActive float64 `json:"playerMoveToActive"`
		PlayerMoveToIR     float64 `json:"playerMoveToIR"`
	} `json:"financeSettings"`
	IsCustomizable  bool   `json:"isCustomizable"`
	IsPublic        bool   `json:"isPublic"`
	Name            string `json:"name"`
	RestrictionType string `json:"restrictionType"`
	RosterSettings  struct {
		IsBenchUnlimited       bool   `json:"isBenchUnlimited"`
		IsUsingUndroppableList bool   `json:"isUsingUndroppableList"`
		LineupLocktimeType     string `json:"lineupLocktimeType"`
		LineupSlotCounts       struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num2  int `json:"2"`
			Num3  int `json:"3"`
			Num4  int `json:"4"`
			Num5  int `json:"5"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num10 int `json:"10"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num18 int `json:"18"`
			Num19 int `json:"19"`
			Num20 int `json:"20"`
			Num21 int `json:"21"`
			Num22 int `json:"22"`
			Num23 int `json:"23"`
			Num24 int `json:"24"`
		} `json:"lineupSlotCounts"`
		LineupSlotStatLimits struct {
		} `json:"lineupSlotStatLimits"`
		MoveLimit      int `json:"moveLimit"`
		PositionLimits struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num2  int `json:"2"`
			Num3  int `json:"3"`
			Num4  int `json:"4"`
			Num5  int `json:"5"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num10 int `json:"10"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
		} `json:"positionLimits"`
		RosterLocktimeType string `json:"rosterLocktimeType"`
		UniverseIds        []int  `json:"universeIds"`
	} `json:"rosterSettings"`
	ScheduleSettings struct {
		Divisions []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"divisions"`
		MatchupPeriodCount  int `json:"matchupPeriodCount"`
		MatchupPeriodLength int `json:"matchupPeriodLength"`
		MatchupPeriods      struct {
			Num1  []int `json:"1"`
			Num2  []int `json:"2"`
			Num3  []int `json:"3"`
			Num4  []int `json:"4"`
			Num5  []int `json:"5"`
			Num6  []int `json:"6"`
			Num7  []int `json:"7"`
			Num8  []int `json:"8"`
			Num9  []int `json:"9"`
			Num10 []int `json:"10"`
			Num11 []int `json:"11"`
			Num12 []int `json:"12"`
			Num13 []int `json:"13"`
			Num14 []int `json:"14"`
			Num15 []int `json:"15"`
			Num16 []int `json:"16"`
			Num17 []int `json:"17"`
		} `json:"matchupPeriods"`
		PeriodTypeID               int    `json:"periodTypeId"`
		PlayoffMatchupPeriodLength int    `json:"playoffMatchupPeriodLength"`
		PlayoffSeedingRule         string `json:"playoffSeedingRule"`
		PlayoffSeedingRuleBy       int    `json:"playoffSeedingRuleBy"`
		PlayoffTeamCount           int    `json:"playoffTeamCount"`
	} `json:"scheduleSettings"`
	ScoringSettings struct {
		AllowOutOfPositionScoring bool   `json:"allowOutOfPositionScoring"`
		HomeTeamBonus             int    `json:"homeTeamBonus"`
		MatchupTieRule            string `json:"matchupTieRule"`
		MatchupTieRuleBy          int    `json:"matchupTieRuleBy"`
		PlayerRankType            string `json:"playerRankType"`
		PlayoffHomeTeamBonus      int    `json:"playoffHomeTeamBonus"`
		PlayoffMatchupTieRule     string `json:"playoffMatchupTieRule"`
		PlayoffMatchupTieRuleBy   int    `json:"playoffMatchupTieRuleBy"`
		ScoringItems              []struct {
			IsReverseItem   bool    `json:"isReverseItem"`
			LeagueRanking   float64 `json:"leagueRanking"`
			LeagueTotal     float64 `json:"leagueTotal"`
			Points          float64 `json:"points"`
			PointsOverrides struct {
				Num16 float64 `json:"16"`
			} `json:"pointsOverrides,omitempty"`
			StatID int `json:"statId"`
		} `json:"scoringItems"`
		ScoringType string `json:"scoringType"`
	} `json:"scoringSettings"`
	Size          int `json:"size"`
	TradeSettings struct {
		AllowOutOfUniverse bool  `json:"allowOutOfUniverse"`
		DeadlineDate       int64 `json:"deadlineDate"`
		Max                int   `json:"max"`
		RevisionHours      int   `json:"revisionHours"`
		VetoVotesRequired  int   `json:"vetoVotesRequired"`
	} `json:"tradeSettings"`
}

// LeagueMemberJSON is the raw member JSON from the ESPN API.
type LeagueMemberJSON struct {
	DisplayName          string `json:"displayName"`
	FirstName            string `json:"firstName"`
	ID                   string `json:"id"`
	LastName             string `json:"lastName"`
	NotificationSettings []struct {
		Enabled bool   `json:"enabled"`
		ID      string `json:"id"`
		Type    string `json:"type"`
	} `json:"notificationSettings"`
}

// LeagueMatchupJSON is the raw matchup JSON from the ESPN API.
type LeagueMatchupJSON struct {
	Away struct {
		CumulativeScore        CumulativeScoreJSON `json:"cumulativeScore"`
		GamesPlayed            int                 `json:"gamesPlayed"`
		RosterForMatchupPeriod struct {
			Entries []interface{} `json:"entries"`
		} `json:"rosterForMatchupPeriod"`
		RosterForMatchupPeriodDelayed struct {
			Entries []interface{} `json:"entries"`
		} `json:"rosterForMatchupPeriodDelayed"`
		TeamID                   int64   `json:"teamId"`
		TotalPoints              float64 `json:"totalPoints"`
		TotalProjectedPointsLive float64 `json:"totalProjectedPointsLive"`
	} `json:"away"`
	Home struct {
		CumulativeScore        CumulativeScoreJSON `json:"cumulativeScore"`
		GamesPlayed            int                 `json:"gamesPlayed"`
		RosterForMatchupPeriod struct {
			Entries []interface{} `json:"entries"`
		} `json:"rosterForMatchupPeriod"`
		RosterForMatchupPeriodDelayed struct {
			Entries []interface{} `json:"entries"`
		} `json:"rosterForMatchupPeriodDelayed"`
		TeamID                   int64   `json:"teamId"`
		TotalPoints              float64 `json:"totalPoints"`
		TotalProjectedPointsLive float64 `json:"totalProjectedPointsLive"`
	} `json:"home"`
	ID              int64  `json:"id"`
	MatchupPeriodID int    `json:"matchupPeriodId"`
	Winner          string `json:"winner"`
}

// LeagueInfoResponseJSON is the raw response from the ESPN API.
type LeagueInfoResponseJSON struct {
	DraftDetail struct {
		Drafted    bool `json:"drafted"`
		InProgress bool `json:"inProgress"`
	} `json:"draftDetail"`
	GameID          int                 `json:"gameId"`
	ID              int64               `json:"id"`
	Members         []LeagueMemberJSON  `json:"members"`
	Schedule        []LeagueMatchupJSON `json:"schedule"`
	ScoringPeriodID int                 `json:"scoringPeriodId"`
	SeasonID        int64               `json:"seasonId"`
	SegmentID       int64               `json:"segmentId"`
	Settings        LeagueSettingsJSON  `json:"settings"`
	Status          struct {
		CreatedAsLeagueType      int   `json:"createdAsLeagueType"`
		CurrentLeagueType        int   `json:"currentLeagueType"`
		CurrentMatchupPeriod     int   `json:"currentMatchupPeriod"`
		FinalScoringPeriod       int   `json:"finalScoringPeriod"`
		FirstScoringPeriod       int   `json:"firstScoringPeriod"`
		IsActive                 bool  `json:"isActive"`
		IsExpired                bool  `json:"isExpired"`
		IsFull                   bool  `json:"isFull"`
		IsPlayoffMatchupEdited   bool  `json:"isPlayoffMatchupEdited"`
		IsToBeDeleted            bool  `json:"isToBeDeleted"`
		IsViewable               bool  `json:"isViewable"`
		IsWaiverOrderEdited      bool  `json:"isWaiverOrderEdited"`
		LatestScoringPeriod      int   `json:"latestScoringPeriod"`
		PreviousSeasons          []int `json:"previousSeasons"`
		TeamsJoined              int   `json:"teamsJoined"`
		TransactionScoringPeriod int   `json:"transactionScoringPeriod"`
		WaiverLastExecutionDate  int64 `json:"waiverLastExecutionDate"`
		WaiverProcessStatus      struct {
		} `json:"waiverProcessStatus"`
	} `json:"status"`
	Teams []struct {
		Abbrev                string   `json:"abbrev"`
		CurrentProjectedRank  int      `json:"currentProjectedRank"`
		DivisionID            int      `json:"divisionId"`
		DraftDayProjectedRank int      `json:"draftDayProjectedRank"`
		ID                    int64    `json:"id"`
		IsActive              bool     `json:"isActive"`
		Location              string   `json:"location"`
		Logo                  string   `json:"logo"`
		LogoType              string   `json:"logoType"`
		Nickname              string   `json:"nickname"`
		Owners                []string `json:"owners"`
		PlayoffSeed           int      `json:"playoffSeed"`
		Points                float64  `json:"points"`
		PointsAdjusted        float64  `json:"pointsAdjusted"`
		PointsDelta           float64  `json:"pointsDelta"`
		PrimaryOwner          string   `json:"primaryOwner"`
		RankCalculatedFinal   int      `json:"rankCalculatedFinal"`
		RankFinal             int      `json:"rankFinal"`
		Record                struct {
			Away struct {
				GamesBack     float64 `json:"gamesBack"`
				Losses        int     `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst float64 `json:"pointsAgainst"`
				PointsFor     float64 `json:"pointsFor"`
				StreakLength  int     `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int     `json:"ties"`
				Wins          int     `json:"wins"`
			} `json:"away"`
			Division struct {
				GamesBack     float64 `json:"gamesBack"`
				Losses        int     `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst float64 `json:"pointsAgainst"`
				PointsFor     float64 `json:"pointsFor"`
				StreakLength  int     `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int     `json:"ties"`
				Wins          int     `json:"wins"`
			} `json:"division"`
			Home struct {
				GamesBack     float64 `json:"gamesBack"`
				Losses        int     `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst float64 `json:"pointsAgainst"`
				PointsFor     float64 `json:"pointsFor"`
				StreakLength  int     `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int     `json:"ties"`
				Wins          int     `json:"wins"`
			} `json:"home"`
			Overall struct {
				GamesBack     float64 `json:"gamesBack"`
				Losses        int     `json:"losses"`
				Percentage    float64 `json:"percentage"`
				PointsAgainst float64 `json:"pointsAgainst"`
				PointsFor     float64 `json:"pointsFor"`
				StreakLength  int     `json:"streakLength"`
				StreakType    string  `json:"streakType"`
				Ties          int     `json:"ties"`
				Wins          int     `json:"wins"`
			} `json:"overall"`
		} `json:"record"`
		Roster struct {
			AppliedStatTotal float64 `json:"appliedStatTotal"`
			Entries          []struct {
				AcquisitionDate       int64       `json:"acquisitionDate"`
				AcquisitionType       string      `json:"acquisitionType"`
				InjuryStatus          string      `json:"injuryStatus"`
				LineupSlotID          int         `json:"lineupSlotId"`
				PendingTransactionIds interface{} `json:"pendingTransactionIds"`
				PlayerID              int         `json:"playerId"`
				PlayerPoolEntry       struct {
					AppliedStatTotal  float64 `json:"appliedStatTotal"`
					ID                int64   `json:"id"`
					KeeperValue       int     `json:"keeperValue"`
					KeeperValueFuture int     `json:"keeperValueFuture"`
					LineupLocked      bool    `json:"lineupLocked"`
					OnTeamID          int     `json:"onTeamId"`
					Player            struct {
						Active               bool `json:"active"`
						DefaultPositionID    int  `json:"defaultPositionId"`
						DraftRanksByRankType struct {
							Standard struct {
								AuctionValue int    `json:"auctionValue"`
								Published    bool   `json:"published"`
								Rank         int    `json:"rank"`
								RankSourceID int    `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int    `json:"slotId"`
							} `json:"STANDARD"`
							Ppr struct {
								AuctionValue int    `json:"auctionValue"`
								Published    bool   `json:"published"`
								Rank         int    `json:"rank"`
								RankSourceID int    `json:"rankSourceId"`
								RankType     string `json:"rankType"`
								SlotID       int    `json:"slotId"`
							} `json:"PPR"`
						} `json:"draftRanksByRankType"`
						Droppable     bool   `json:"droppable"`
						EligibleSlots []int  `json:"eligibleSlots"`
						FirstName     string `json:"firstName"`
						FullName      string `json:"fullName"`
						ID            int64  `json:"id"`
						Injured       bool   `json:"injured"`
						InjuryStatus  string `json:"injuryStatus"`
						LastName      string `json:"lastName"`
						LastNewsDate  int64  `json:"lastNewsDate"`
						Ownership     struct {
							AuctionValueAverage  float64 `json:"auctionValueAverage"`
							AverageDraftPosition float64 `json:"averageDraftPosition"`
							PercentChange        float64 `json:"percentChange"`
							PercentOwned         float64 `json:"percentOwned"`
							PercentStarted       float64 `json:"percentStarted"`
						} `json:"ownership"`
						ProTeamID int `json:"proTeamId"`
						Rankings  struct {
							Num0 []struct {
								AuctionValue int     `json:"auctionValue"`
								Published    bool    `json:"published"`
								Rank         int     `json:"rank"`
								RankSourceID int     `json:"rankSourceId"`
								RankType     string  `json:"rankType"`
								SlotID       int     `json:"slotId"`
								AverageRank  float64 `json:"averageRank,omitempty"`
							} `json:"0"`
						} `json:"rankings"`
						SeasonOutlook string `json:"seasonOutlook"`
						Stats         []struct {
							AppliedAverage  float64                `json:"appliedAverage,omitempty"`
							AppliedStats    map[string]interface{} `json:"appliedStats,omitempty"`
							AppliedTotal    float64                `json:"appliedTotal"`
							ExternalID      string                 `json:"externalId"`
							ID              string                 `json:"id"`
							ProTeamID       int                    `json:"proTeamId"`
							ScoringPeriodID int                    `json:"scoringPeriodId"`
							SeasonID        int64                  `json:"seasonId"`
							StatSourceID    int64                  `json:"statSourceId"`
							StatSplitTypeID int64                  `json:"statSplitTypeId"`
							Stats           map[string]interface{} `json:"stats,omitempty"`
						} `json:"stats"`
						UniverseID int `json:"universeId"`
					} `json:"player"`
					RosterLocked bool   `json:"rosterLocked"`
					Status       string `json:"status"`
					TradeLocked  bool   `json:"tradeLocked"`
				} `json:"playerPoolEntry"`
				Status string `json:"status"`
			} `json:"entries"`
			TradeReservedEntries int `json:"tradeReservedEntries"`
		} `json:"roster"`
		TradeBlock struct {
		} `json:"tradeBlock"`
		TransactionCounter struct {
			AcquisitionBudgetSpent   int64 `json:"acquisitionBudgetSpent"`
			Acquisitions             int64 `json:"acquisitions"`
			Drops                    int64 `json:"drops"`
			MatchupAcquisitionTotals struct {
			} `json:"matchupAcquisitionTotals"`
			Misc         int64   `json:"misc"`
			MoveToActive int64   `json:"moveToActive"`
			MoveToIR     int64   `json:"moveToIR"`
			Paid         float64 `json:"paid"`
			TeamCharges  float64 `json:"teamCharges"`
			Trades       int64   `json:"trades"`
		} `json:"transactionCounter"`
		WaiverRank    int `json:"waiverRank"`
		DraftStrategy struct {
			FutureKeeperPlayerIds []interface{} `json:"futureKeeperPlayerIds"`
			KeeperPlayerIds       []interface{} `json:"keeperPlayerIds"`
		} `json:"draftStrategy,omitempty"`
	} `json:"teams"`
}

// ActivityJSON is the return type of the recent activity API call.
type ActivityJSON struct {
	Topics []struct {
		Author       string `json:"author"`
		CreationInfo struct {
			ClientAddress interface{} `json:"clientAddress"`
			Date          int64       `json:"date"`
			Platform      string      `json:"platform"`
			Source        string      `json:"source"`
		} `json:"creationInfo"`
		Date           int64  `json:"date"`
		DateEdited     int64  `json:"dateEdited"`
		ID             string `json:"id"`
		IsDeleted      bool   `json:"isDeleted"`
		IsEdited       bool   `json:"isEdited"`
		LastUpdateInfo struct {
			ClientAddress interface{} `json:"clientAddress"`
			Date          int64       `json:"date"`
			Platform      string      `json:"platform"`
			Source        string      `json:"source"`
		} `json:"lastUpdateInfo"`
		Messages []struct {
			Author       string `json:"author"`
			CreationInfo struct {
				ClientAddress interface{} `json:"clientAddress"`
				Date          int64       `json:"date"`
				Platform      string      `json:"platform"`
				Source        string      `json:"source"`
			} `json:"creationInfo"`
			Date              int64  `json:"date"`
			For               int64  `json:"for"`
			From              int64  `json:"from"`
			ID                string `json:"id"`
			IsAlternateFormat bool   `json:"isAlternateFormat"`
			IsDeleted         bool   `json:"isDeleted"`
			IsEdited          bool   `json:"isEdited"`
			MessageTypeID     int    `json:"messageTypeId"`
			TargetID          int64  `json:"targetId"`
			To                int64  `json:"to"`
			TopicID           string `json:"topicId"`
		} `json:"messages"`
		TargetID          interface{} `json:"targetId,omitempty"`
		TotalMessageCount int         `json:"totalMessageCount"`
		Type              string      `json:"type"`
		For               int64       `json:"for,omitempty"`
	} `json:"topics"`
}

// PlayerJSON is a player from the active players endpoint.
type PlayerJSON struct {
	Player struct {
		DefaultPositionID int    `json:"defaultPositionId"`
		Droppable         bool   `json:"droppable"`
		EligibleSlots     []int  `json:"eligibleSlots"`
		FirstName         string `json:"firstName"`
		FullName          string `json:"fullName"`
		ID                int64  `json:"id"`
		LastName          string `json:"lastName"`
		Ownership         struct {
			PercentOwned float64 `json:"percentOwned"`
		} `json:"ownership"`
		ProTeamID  int `json:"proTeamId"`
		UniverseID int `json:"universeId"`
	} `json:"player"`
}

// PlayersJSON is the active players endpoint repsonse.
type PlayersJSON []PlayerJSON
