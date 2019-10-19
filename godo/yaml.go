package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type conf struct {
    Hits int64 `yaml:"hits"`
    Time int64 `yaml:"time"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("person.yml")
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
    var c conf
    c.getConf()

    fmt.Println(c.Firstname)
}

members: 2
size: "s-1vcpu-2gb"
image: "centos-7-x64"
location: "sfo2"