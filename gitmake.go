package main

import "fmt"
import "strings"
import "os"
import "golang.org/x/oauth2"
import "github.com/google/go-github/github"

func main() {
  name := getName()
  client := getClient()
  createRepo(client, name)
}

func getName() string {
  dir, _ := os.Getwd()
  parts := strings.Split(dir, "/")
  lastIndex := len(parts) - 1
  name := parts[lastIndex]
  return name
}

func getClient() *github.Client {
  // Load oauth
  token := os.Getenv("GIT_TOKEN")
  ts:= oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token},
  )

  // Authenticate
  tc:= oauth2.NewClient(oauth2.NoContext, ts)
  client := github.NewClient(tc)

  return client
}

func createRepo(client * github.Client, name string) bool {
  repo := github.Repository{Name: &name}
  _, _, err := client.Repositories.Create("", &repo)

  if err == nil {
    fmt.Println("Created repo", name)
  } else {
    fmt.Println("There was an error creating", name)
    fmt.Println(err)
  }

  return err == nil
}
