package repo

import (
	"net/http"
	"strings"

	"github.com/franela/goreq"
)

const (
	gitHubRepoURL = "https://api.github.com/user/repos"
)

type gitHub struct {
	AccessToken string
}

func newGitHub(accessToken string) walker {
	return &gitHub{
		AccessToken: accessToken,
	}
}

type gitHubRepoJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Fork        bool   `json:"fork"`
	RepoURL     string `json:"html_url"`
	Homepage    string `json:"homepage"`
}

func (w *gitHub) repos() ([]*Repo, error) {
	authorization := strings.Join([]string{
		"token",
		w.AccessToken,
	}, " ")
	resp, err := goreq.Request{
		Method:      "GET",
		Uri:         gitHubRepoURL,
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

	var reposJ []*gitHubRepoJSON
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
			RepoURL:     repoJ.RepoURL,
			Homepage:    repoJ.Homepage,
		}
	}

	return repos, nil
}

func (w *gitHub) repoDeps(name string) {

}
