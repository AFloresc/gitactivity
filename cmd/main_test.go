package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitactivity/domain"
	"io"
	"testing"
)

func TestIt(t *testing.T) {
	//var username string = argWithProg[1]
	response, err := getData("kamranahmedse")
	if err != nil {
		fmt.Println(err.Error())
	}

	var result []domain.Event
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(responseData), &result)
	if err != nil {
		fmt.Println(errors.New("error parsing JSON"))
		return
	}

	fmt.Printf("- There are %d repositories \n", len(result))
	mapRepos := createReposMap(result)
	for key, repo := range mapRepos {
		pushCounter := 0
		pullCounter := 0
		commitsCounter := 0
		issueEventCounter := 0
		for _, event := range result {
			if key == event.Repo.ID {
				if event.Type == "PushEvent" {
					pushCounter += 1
					//TODO get commits
					commitsCounter = commitsCounter + event.Payload.Size
				}
				if event.Type == "PullEvent" {
					pullCounter += 1
				}
				if event.Type == "IssuesEvent" {
					issueEventCounter += 1
				}
			}
		}
		fmt.Printf("- Pushed %d for repository %s \n", pushCounter, repo)
		fmt.Printf("- Has %d commits for repository %s \n", commitsCounter, repo)
		fmt.Printf("- Issues %d for repository %s \n", issueEventCounter, repo)
	}
}
