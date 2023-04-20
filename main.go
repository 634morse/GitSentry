package main

// Import all of the required library packages
import (
	"github.com/634morse/GitSentry/repo"
	"flag"
	"fmt"
)

func helpfile() {
	fmt.Println(`
Usage: GitSentry is a tool written in Go, and is used to enumerate GitHub repositories.

Commands:
- repo    Enumerate repositories
- rights  Enumerate user or organization rights
- etc     (blah blah blah)`)
}

func main() {
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

	// calling repo-download packages based on user or org
	if RepoType == "OrgRepo" {
		repo.org_repo(Org, GitUser, token, GitUser, OutFile)
	}

}
