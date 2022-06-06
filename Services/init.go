package services

import config "github.com/Zenovore/ocbc-practice-day4/Config"

var (
	userInterface UserServiceInterface
)

func init() {
	config.InitDB()
	userInterface = NewService()
}
