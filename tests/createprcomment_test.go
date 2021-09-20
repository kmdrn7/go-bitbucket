package tests

import (
	"fmt"
	"github.com/kmdrn7/go-bitbucket"
	"testing"
)

func TestCreatePRComment(t *testing.T) {

	c := setup(t)

	comment, err := c.Repositories.PullRequests.CreateComment(&bitbucket.PullRequestsCommentsOptions{
		PullRequestsOptions: bitbucket.PullRequestsOptions{
			ID:                "1",
			Owner:             owner,
			RepoSlug:          repo,
		},
		CommentContent: "Wow TF 2",
	})
	if err != nil {
		return
	}

	fmt.Println(comment)
}
