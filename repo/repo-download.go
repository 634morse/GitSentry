package repo

// Import all of the required library packages
import (
	"context"
	"flag"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	// Importing Go-git
	// Better for interacting with git
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"

	// Importing go-github
	// Package github provides a client for using the GitHub API.
	"github.com/google/go-github/v39/github"
)

func user_repo() {
	// flags & arguments
	var RepoType string
	var GitUser string
	var RepoName string
	var token string
	var OutFile string
	var Org string

	flag.StringVar(&RepoType, "RepoType", "", "OrgRepo/UserRepo")
	flag.StringVar(&GitUser, "GitUser", "", "Your git username")
	flag.StringVar(&RepoName, "RepoName", "", "The Name of the repo")
	flag.StringVar(&token, "token", "", "Your Github Token")
	flag.StringVar(&OutFile, "OutFile", "", "Location to download the repo to")
	flag.StringVar(&Org, "Org", "", "The Name of yor github Organization")

	flag.Parse()
	// flags & arguments

	if RepoType == "OrgRepo" {

	}

	// creating context and token source vars
	ctx := context.Background()
	tokensource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, tokensource)
	client := github.NewClient(tc)

	//https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	//
	opts := &github.RepositoryListByOrgOptions{Type: "all"}
	repos, _, err := client.Repositories.ListByOrg(ctx, Org, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}

	//set gitclone options
	options := &git.CloneOptions{
		URL:          "https://github.com/" + GitUser + "/" + RepoName + ".git",
		Auth:         &http.BasicAuth{Username: GitUser, Password: token},
		SingleBranch: false,
	}
	Dir := OutFile + "\\" + RepoName

	//uses go-gits PlainClone function to clone the repository into the given dir, just as a normal git clone does
	//https://pkg.go.dev/github.com/go-git/go-git/v5#section-readme
	_, err := git.PlainClone(Dir, false, options)
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		os.Exit(1)
	}
}

func org_repo() {
	// flags & arguments
	var RepoType string
	var GitUser string
	var RepoName string
	var token string
	var OutFile string
	var Org string

	flag.StringVar(&RepoType, "RepoType", "", "OrgRepo/UserRepo")
	flag.StringVar(&GitUser, "GitUser", "", "Your git username")
	flag.StringVar(&RepoName, "RepoName", "", "The Name of the repo")
	flag.StringVar(&token, "token", "", "Your Github Token")
	flag.StringVar(&OutFile, "OutFile", "", "Location to download the repo to")
	flag.StringVar(&Org, "Org", "", "The Name of yor github Organization")

	flag.Parse()
	// flags & arguments

	// creating context and token source vars
	ctx := context.Background()
	tokensource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, tokensource)
	client := github.NewClient(tc)

	//https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	//
	opts := &github.RepositoryListByOrgOptions{Type: "all"}
	repos, _, err := client.Repositories.ListByOrg(ctx, Org, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}
