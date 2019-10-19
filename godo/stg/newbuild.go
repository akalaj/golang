package main

import (
	"fmt"
	"context"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//VARIABLES
const (
	pat = "83621c2023b2d28b1934b34d3179e6f455911b8f046f029ffbc2088cc6c4a430"
)

//STRUCTS
type conf struct {
	Projectname string `yaml:"projectname"`
	Members     int  `yaml:"members"`
	Size        string `yaml:"size"`
	Image       string `yaml:"image"`
	Location    string `yaml:"location"`
}

type TokenSource struct {
	AccessToken string
}

//FUNCTIONS
func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	//Build config struct
	var configFile conf
	configFile.getConf()

	//build auth object
	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	//determine Droplet config
	var dropletNames []string

	for i := 1; i < configFile.Members+1; i++ {
		dropletNames = append(dropletNames, fmt.Sprintf("%s-%v", configFile.Projectname, i))
	}

	mykey := godo.DropletCreateSSHKey{ID: 25339235}
	mykeys := []godo.DropletCreateSSHKey{mykey}

	createRequest := &godo.DropletMultiCreateRequest{
		Names:  dropletNames,
		Region: configFile.Location,
		Size:   configFile.Size,
		Image: godo.DropletCreateImage{
			Slug: configFile.Image,
		},
		SSHKeys: mykeys,
	}

	ctx := context.TODO()

	client.Droplets.CreateMultiple(ctx, createRequest)

}
