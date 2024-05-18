package infrastructure

import (
	databaseTypes "slanj/src/common/database/types"
	types "slanj/src/common/steam/types"
)

type IScoreRepository interface {
	GetScores() []databaseTypes.GetScoreResponse
}

type ISteamRepository interface {
	GetSteamPlayers(ids []string) []types.SteamPlayer
}
