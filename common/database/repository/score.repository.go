package repository

import (
	"slanj/common/database"
	"slanj/common/database/schema"
	"slanj/common/database/types"
)

type ScoreRepository struct{}

func (s *ScoreRepository) SaveScore(id string, score int) bool {
	data := schema.Score{
		Score:  score,
		UserID: id,
	}

	output := database.GetConnection().Create(&data)
	return output.RowsAffected == 1
}

func (s *ScoreRepository) GetScores() []types.GetScoreResponse {
	var scores []types.GetScoreResponse

	queryString := `
		SELECT DISTINCT s.user_id, MAX(score) AS score 
		FROM Scores s
		INNER JOIN Users u ON s.user_id = u.id
		GROUP BY s.user_id
		ORDER BY MAX(s.score) DESC
		LIMIT 10
	`
	database.GetConnection().Raw(queryString).Scan(&scores)

	response := []types.GetScoreResponse{}
	for _, score := range scores {
		data := types.GetScoreResponse{UserID: score.UserID, Score: score.Score}
		response = append(response, data)
	}

	return response
}
