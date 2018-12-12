/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var token string

func main() {
	fmt.Println("Starting contributors server...")
	if _, err := os.Stat("token"); os.IsNotExist(err) {
		fmt.Printf("Can't find token file: %s", err)
		os.Exit(1)
	}
	t, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Printf("Can't read token file: %s", err)
		os.Exit(1)
	}
	token = string(t)
	if token == "yourtoken" {
		fmt.Println("CHANGE YOUR GITHUB TOKEN IN backend/svc-contributors/token!!!")
		os.Exit(1)
	}

	fmt.Print(token)

	http.HandleFunc("/contributors", contributorsHandler())
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func contributorsHandler() func(http.ResponseWriter, *http.Request) {
	client, ctx := NewClient()
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request!")
		page := 1
		emptyResponse := false
		numContribs := 0
		if n, writeErr := fmt.Fprintf(w, "["); n <= 0 || writeErr != nil {
			fmt.Printf("error writing response %s", writeErr)
		}
		for ; emptyResponse == false; page++ {
			users, _, err := client.Repositories.ListCollaborators(ctx, "kubernetes", "minikube", &github.ListCollaboratorsOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: 100,
				},
			})
			emptyResponse = len(users) == 0
			numContribs += len(users)
			if err != nil {
				fmt.Printf("query error: %s\n", err)
				if n, writeErr := fmt.Fprintf(w, "{'error':'%s'}", err); n <= 0 || writeErr != nil {
					fmt.Printf("error returning error response: %s, %s", writeErr, err)
				}
				return
			}

			for idx, user := range users {
				if idx == 0 && page > 1 {
					if n, writeErr := fmt.Fprint(w, ","); n <= 0 || writeErr != nil {
						fmt.Printf("error writing response: %s", writeErr)
					}
				}
				if n, writeErr := fmt.Fprintf(w, `{"avatar": "%s", "login": "%s"}`, user.GetAvatarURL(), user.GetLogin()); n <= 0 || writeErr != nil {
					fmt.Printf("error writing response: %s", writeErr)
				}
				if idx < len(users)-1 {
					if n, writeErr := fmt.Fprint(w, ",\n"); n <= 0 || writeErr != nil {
						fmt.Printf("error writing response: %s", writeErr)
					}
				}
			}
		}
		if n, writeErr := fmt.Fprintf(w, "]"); n <= 0 || writeErr != nil {
			fmt.Printf("error writing response %s", writeErr)
		}
		fmt.Printf("found %d contributors!\n", numContribs)
	}
}

func NewClient() (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, ctx
}
