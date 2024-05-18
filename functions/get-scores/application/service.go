package application

import (
	database "slanj/common/database/repository"
	steam "slanj/common/steam/repository"
	infrastructure "slanj/functions/get-scores/infrastructure"

	entity "slanj/functions/get-scores/domain"
)

var databaseRepository infrastructure.IScoreRepository = &database.ScoreRepository{}
var steamRepository infrastructure.ISteamRepository = &steam.SteamRepository{}

type Service struct{}

func (s *Service) GetScores() []entity.Score {
	scores := databaseRepository.GetScores()

	var userIds []string
	for _, user := range scores {
		userIds = append(userIds, user.UserID)
	}

	players := steamRepository.GetSteamPlayers(userIds)

	var result []entity.Score
	for _, score := range scores {
		for _, player := range players {
			if score.UserID == player.SteamID {
				result = append(result, entity.Score{UserID: score.UserID, Name: player.PersonaName, Score: score.Score})
			}
		}
	}

	return result
}
