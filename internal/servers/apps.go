package servers

import (
	"ddd_demo2/internal/user"
	"ddd_demo2/internal/youke"
)

type Apps struct {
	UserApp  user.UserAppInterface
	YoukeApp youke.YoukeAppInterface
}

func NewApps(repos *Repos) *Apps {
	return &Apps{
		UserApp:  user.NewUserApp(repos.UserRepo, repos.AuthRepo, repos.BillRepo),
		YoukeApp: youke.NewYoukeApp(repos.YoukeRepo),
	}
}
