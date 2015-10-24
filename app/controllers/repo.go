package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/modules/jsonapi"
	"github.com/Gr1N/pacman/app/modules/repo"
)

type Repo struct {
	TokenAuthenticated
}

type repoItemAttrs struct {
	Name        string `json:"name"`
	Description string `json:"desciption"`
	Private     bool   `json:"private"`
	Fork        bool   `json:"fork"`
	URL         string `json:"url"`
	Homepage    string `json:"homepage"`
	Created     int64  `json:"created"`
}

func (c Repo) GetAll(service string) revel.Result {
	user := c.getUser()
	// FIXME: Handle error
	repos, _ := models.GetUserReposByService(user.ID, service)

	items := make([]*jsonapi.Item, len(repos))
	for i := range items {
		items[i] = c.item(repos[i])
	}

	return c.RenderJSONOk(items)
}

func (c Repo) UpdateAll(service string) revel.Result {
	user := c.getUser()
	// FIXME: Handle error
	_ = repo.HandleUpdate(user.ID, service)

	return c.RenderNoContent()
}

func (c Repo) item(repo *models.Repo) *jsonapi.Item {
	return &jsonapi.Item{
		Type: "repos",
		ID:   repo.ID,
		Attributes: repoItemAttrs{
			Name:        repo.Name,
			Description: repo.Description,
			Private:     repo.Private,
			Fork:        repo.Fork,
			URL:         repo.URL,
			Homepage:    repo.Homepage,
			Created:     repo.CreatedAt.Unix(),
		},
		Links: jsonapi.ItemLinks{
			Self: "TBD",
		},
	}
}
