package infrastructure

type IAddScoreRepository interface {
	SaveScore(userId string, score int) bool
}
