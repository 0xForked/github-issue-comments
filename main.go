package main

import (
	"context"
	"fmt"
	"github.com/0xForked/github-issue-comment/cfg"
	"github.com/google/go-github/v53/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func main() {
	// LOAD ENVIRONMENT VARIABLES
	viper.SetConfigFile(".env")
	cfg.LoadEnv()
	// INIT GITHUB CONNECTION
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Instance.GitHubAccessToken})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	comment, _, err := client.Issues.CreateComment(
		context.Background(),
		cfg.Instance.GitHubOwnerName,
		cfg.Instance.GitHubRepoName,
		cfg.Instance.GitHubIssueNumber,
		&github.IssueComment{
			Body: github.String("this new error message @aasumitro please fix it"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}
