package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
)

func main() {

	temp, err := os.MkdirTemp("", "tmp")
	if err != nil {
		panic(err)
	}
	auth1 := &http.BasicAuth{Username: "oauth2", Password: os.Getenv("GITHUB_TOKEN")}
	auth2 := &http.BasicAuth{Username: "something", Password: os.Getenv("GITHUB_TOKEN")}
	r, err := git.PlainClone(temp, false, &git.CloneOptions{
		URL:  "https://github.com/l-ehlers/go-git-1195",
		Auth: auth1,
	})
	if err != nil {
		panic(err)
	}
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	submodules, err := w.Submodules()
	if err != nil {
		panic(err)
	}
	err = submodules.Update(&git.SubmoduleUpdateOptions{Init: true, Auth: auth2})
	if err != nil {
		panic(err)
	}
}
