package servers

import (
	"dc2/internal/user"
	"dc2/internal/youke"
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
