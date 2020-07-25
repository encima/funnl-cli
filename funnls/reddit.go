package funnls

import (
	"fmt"
	"log"

	"github.com/jzelinskie/geddit"
	"github.com/spf13/viper"
)

// Reddit struct
type Reddit struct {
	Posts []string
}

// RdAuth Authenticate User for Reddit
func RdAuth() *geddit.OAuthSession {
	o, err := geddit.NewOAuthSession(
		viper.GetString("funnls.reddit.auth.client_id"),
		viper.GetString("funnls.reddit.auth.client_secret"),
		"Funnl Module for Reddit Saved Posts, comments etc retrieval",
		"http://localhost:5000",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create new auth token for confidential clients (personal scripts/apps).
	err = o.LoginAuth(viper.GetString("funnls.reddit.auth.username"), viper.GetString("funnls.reddit.auth.pwd"))
	if err != nil {
		log.Fatal(err)
	}
	return o
}

// RdGetPosts Retrieve a list of the user's posts
func RdGetPosts(o *geddit.OAuthSession) []*geddit.Submission {
	l := geddit.ListingOptions{Count: 400}
	p, err := o.MySavedLinks(l)
	if err != nil {
		fmt.Println(err)
	}
	return p
}
