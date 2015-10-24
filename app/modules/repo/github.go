package repo

func newGitHubWalker() *Walker {
	return &Walker{
		Endpoint: Endpoint{
			RepoURL: "https://api.github.com/user/repos",
		},
	}
}
