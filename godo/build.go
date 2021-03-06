package main

import (
	"context"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	pat = ""
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	dropletNames := []string{"my-godo-droplet1"}
	mykey := godo.DropletCreateSSHKey{ID: 25339235}
	mykeys := []godo.DropletCreateSSHKey{mykey}

	createRequest := &godo.DropletMultiCreateRequest{
		Names:   dropletNames,
		Region: "sfo2",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "centos-7-x64",
		},
		SSHKeys: mykeys,
	}

	ctx := context.TODO()

	client.Droplets.CreateMultiple(ctx, createRequest)

}
