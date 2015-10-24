package repo

import (
	"errors"
	"net/http"
	"strings"

	"github.com/franela/goreq"

	"github.com/Gr1N/pacman/app/models"
)

var (
	ErrAccessTokenInvalid   = errors.New("AccessToken invalid")
	ErrReposResponseInvalid = errors.New("Got invalid repos response")
)

type Walker struct {
	Endpoint Endpoint
}

type Endpoint struct {
	RepoURL string
}

type Repo struct {
	Name        string
	Description string
	Private     bool
	Fork        bool
	URL         string
	Homepage    string
}

// FIXME: Supports only GitHub repos response
type repoJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Fork        bool   `json:"fork"`
	URL         string `json:"html_url"`
	Homepage    string `json:"homepage"`
}

func HandleUpdate(userID int64, serviceName string) error {
	service, err := models.GetUserService(userID, serviceName)
	if err != nil {
		return err
	}

	Walker := NewWalker(service.Name)

	repos, err := Walker.Repos(service.AccessToken)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		// FIXME: Handle DB error
		// TODO: Update repo
		models.CreateUserRepo(service.ID, repo.Name, repo.Description,
			repo.Private, repo.Fork, repo.URL, repo.Homepage)
	}

	return nil
}

func NewWalker(serviceName string) *Walker {
	return map[string]func() *Walker{
		"github": newGitHubWalker,
	}[serviceName]()
}

func (w Walker) Repos(accessToken string) ([]*Repo, error) {
	authorization := strings.Join([]string{
		"token",
		accessToken,
	}, " ")
	resp, err := goreq.Request{
		Method:      "GET",
		Uri:         w.Endpoint.RepoURL,
		ContentType: "application/json",
		Accept:      "application/json",
	}.WithHeader("Authorization", authorization).Do()
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return nil, ErrAccessTokenInvalid
	}

	var reposJ []*repoJSON
	if err := resp.Body.FromJsonTo(&reposJ); err != nil {
		return nil, ErrReposResponseInvalid
	}

	repos := make([]*Repo, len(reposJ))
	for i := range reposJ {
		repoJ := reposJ[i]
		repos[i] = &Repo{
			Name:        repoJ.Name,
			Description: repoJ.Description,
			Private:     repoJ.Private,
			Fork:        repoJ.Fork,
			URL:         repoJ.URL,
			Homepage:    repoJ.Homepage,
		}
	}

	return repos, nil
}
