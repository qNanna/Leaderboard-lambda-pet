package infrastructure

import (
	databaseTypes "slanj/common/database/types"
	types "slanj/common/steam/types"
)

type IScoreRepository interface {
	GetScores() []databaseTypes.GetScoreResponse
}

type ISteamRepository interface {
	GetSteamPlayers(ids []string) []types.SteamPlayer
}
