package repo

import (
	"errors"

	"github.com/Gr1N/pacman/app/models"
)

var (
	ErrAccessTokenInvalid   = errors.New("AccessToken invalid")
	ErrReposResponseInvalid = errors.New("Got invalid repos response")
)

type walker interface {
	repos(accessToken string) ([]*Repo, error)
}

type Repo struct {
	Name        string
	Description string
	Private     bool
	Fork        bool
	RepoURL     string
	Homepage    string
}

func newWalker(serviceName string) walker {
	return map[string]func() walker{
		"github": newGitHub,
	}[serviceName]()
}

func HandleUpdate(userID int64, serviceName string) error {
	service, err := models.GetUserService(userID, serviceName)
	if err != nil {
		return err
	}

	walker := newWalker(service.Name)

	repos, err := walker.repos(service.AccessToken)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		// FIXME: Handle DB error
		// TODO: Update repo
		models.CreateUserRepo(service.ID, repo.Name, repo.Description,
			repo.Private, repo.Fork, repo.RepoURL, repo.Homepage)
	}

	return nil
}
