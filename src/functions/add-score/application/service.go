package application

import (
	database "slanj/src/common/database/repository"
	infrastructure "slanj/src/functions/add-score/infrastructure"
)

var repository infrastructure.IAddScoreRepository = &database.ScoreRepository{}

type Service struct{}

func (s *Service) SaveScore(userId string, score int) bool {
	return repository.SaveScore(userId, score)
}
