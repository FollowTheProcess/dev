// Package repo implements dev's interaction with repos (projects), either local on the user's filesystem or remote on e.g. GitHub.
package repo

import (
	"fmt"
	"os"
)

// Repo represents a single software project.
type Repo struct {
	Owner string // The owner i.e GitHub username
	Name  string // The name of the project
	Path  string // The path the repo would have on the local filesystem, whether it currently exists or not
}

// CloneURL returns the URL to pass to git clone in order to clone this repo locally.
func (r Repo) CloneURL() string {
	return fmt.Sprintf("%s.git", r.URL())
}

// URL returns the URL to the repo's homepage.
func (r Repo) URL() string {
	return fmt.Sprintf("https://github.com/%s/%s", r.Owner, r.Name)
}

// IssuesURL returns the URL to the repo's issues page.
func (r Repo) IssuesURL() string {
	return fmt.Sprintf("%s/issues", r.URL())
}

// PullsURL returns the URL to the repo's PR page.
func (r Repo) PullsURL() string {
	return fmt.Sprintf("%s/pulls", r.URL())
}

// ExistsLocal returns whether or not the repo exists on the user's local filesystem, that is;
// does the Path exist.
func (r Repo) ExistsLocal() bool {
	_, err := os.Stat(r.Path)
	return err == nil
}
