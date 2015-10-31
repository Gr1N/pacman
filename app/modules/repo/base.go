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
	repos() ([]*Repo, error)
	repoDeps(string)
}

type Repo struct {
	Name        string
	Description string
	Private     bool
	Fork        bool
	RepoURL     string
	Homepage    string
}

func newWalker(serviceName, serviceAccessToken string) walker {
	return map[string]func(string) walker{
		"github": newGitHub,
	}[serviceName](serviceAccessToken)
}

func HandleUpdate(userID int64, serviceName string) error {
	service, err := models.GetUserService(userID, serviceName)
	if err != nil {
		return err
	}

	walker := newWalker(service.Name, service.AccessToken)

	repos, err := walker.repos()
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

func HandleUpdateDeps(userID int64, serviceName, repoName string) {

}
