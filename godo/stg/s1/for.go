package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)


//STRUCTS
type conf struct {
    Projectname string `yaml:"projectname"`
    Members int64 `yaml:"members"`
    Size string `yaml:"size"`
    Image string `yaml:"image"`
    Location string `yaml:"location"`
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


func main() {
	//Build config struct
    var configFile conf
    configFile.getConf()

    fmt.Println(configFile.Members)
}