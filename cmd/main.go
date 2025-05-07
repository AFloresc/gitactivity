package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitactivity/domain"
	"io"
	"net/http"
	"os"
)

const URLbase = "https://api.github.com/users/"

var EventType = []string{"CreateEvent", "DeleteEvent", "ForkEvent", "GollumEvent", "IssueCommentEvent", "IssuesEvent", "MemberEvent", "PublicEvent", "PullRequestEvent", "PullRequestReviewEvent", "PullRequestReviewCommentEvent", "PullRequestReviewCommentEvent", "PullRequestReviewThreadEvent", "PushEvent", "ReleaseEvent", "SponsorshipEvent", "WatchEvent"}

func main() {
	argWithProg := os.Args

	if len(argWithProg) < 2 {
		fmt.Println("Plesa add Github username as argument")
		return
	}

	//var username string = argWithProg[1]
	response, err := getData(argWithProg[1])
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

	mapRepos := createReposMap(result)
	for key, repo := range mapRepos {
		pushCounter := 0
		pullCounter := 0
		//commitsCounter := 0
		issueEventCounter := 0
		for _, event := range result {
			if key == event.Repo.ID {
				if event.Type == "PushEvent" {
					pushCounter += 1
					//TODO get commits
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
		fmt.Printf("- Pulled%d for repository %s \n", pullCounter, repo)
		fmt.Printf("- Issues %d for repository %s \n", issueEventCounter, repo)
	}

}

func getData(username string) (*http.Response, error) {
	response, err := http.Get(URLbase + username + "/events")
	if err != nil {
		return response, err
	}
	return response, nil
}

func createReposMap(events []domain.Event) map[int]string {
	mapRepos := make(map[int]string)

	for _, val := range events {
		_, found := mapRepos[val.Repo.ID]
		if !found {
			mapRepos[val.Repo.ID] = val.Repo.Name
		}
	}
	return mapRepos
}
