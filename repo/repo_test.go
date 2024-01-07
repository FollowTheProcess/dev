package repo_test

import (
	"os"
	"testing"

	"github.com/FollowTheProcess/dev/repo"
	"github.com/FollowTheProcess/test"
)

func TestRepoBasics(t *testing.T) {
	cwd, err := os.Getwd()
	test.Ok(t, err, "couldn't get cwd")

	tests := []struct {
		name        string
		repo        repo.Repo
		url         string
		cloneURL    string
		issuesURL   string
		pullsURL    string
		existsLocal bool
	}{
		{
			name: "basic",
			repo: repo.Repo{
				Owner: "me",
				Name:  "project",
				Path:  "/doesnt/exist",
			},
			url:         "https://github.com/me/project",
			cloneURL:    "https://github.com/me/project.git",
			issuesURL:   "https://github.com/me/project/issues",
			pullsURL:    "https://github.com/me/project/pulls",
			existsLocal: false,
		},
		{
			name: "exists",
			repo: repo.Repo{
				Owner: "someoneelse",
				Name:  "differentproject",
				Path:  cwd, // cwd will always exist
			},
			url:         "https://github.com/someoneelse/differentproject",
			cloneURL:    "https://github.com/someoneelse/differentproject.git",
			issuesURL:   "https://github.com/someoneelse/differentproject/issues",
			pullsURL:    "https://github.com/someoneelse/differentproject/pulls",
			existsLocal: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.Equal(t, tt.repo.URL(), tt.url)
			test.Equal(t, tt.repo.CloneURL(), tt.cloneURL)
			test.Equal(t, tt.repo.IssuesURL(), tt.issuesURL)
			test.Equal(t, tt.repo.PullsURL(), tt.pullsURL)
			test.Equal(t, tt.repo.ExistsLocal(), tt.existsLocal)
		})
	}
}
