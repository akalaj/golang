package main

//import "github.com/digitalocean/godo"
import "fmt"



func main() {

	type DropletCreateSSHKey struct {
    	ID          int
    	Fingerprint string
	}

	key := make([]DropletCreateSSHKey, 1)
	mySSHKey = append(key, {id:123})

	//mySSHKey := godo.[]DropletCreateSSHKey{ID:25339235}

	
	fmt.Println(mySSHKey)
}