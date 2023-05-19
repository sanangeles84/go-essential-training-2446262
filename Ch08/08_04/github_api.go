// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"` // from public_repos in JSON
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	/* TODO:
	Call the github API for a given login
	e.g. https://api.github.com/users/tebeka

	And return User struct
	*/

	resp, err := http.Get("https://api.github.com/users/" + url.PathEscape(login))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	var usr User
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&usr); err != nil {
		return nil, err
	}

	return &usr, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
