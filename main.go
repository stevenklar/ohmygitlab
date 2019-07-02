package main

import (
    "log"
    "fmt"

    "github.com/xanzy/go-gitlab"
    "github.com/spf13/viper"
)

type App struct {
    client *gitlab.Client
}

func main() {
    // configuration
    viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("yaml") 
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    token := viper.GetString("token")
    url := viper.GetString("url")

    // initailize gitlab client
    app := &App{}
    app.client = gitlab.NewClient(nil, token)
    app.client.SetBaseURL(url)

    app.printOpenMergeRequests()
}

func (app *App) printOpenMergeRequests() {
    // list merge requests
    rOpts := &gitlab.ListMergeRequestsOptions{
        State: gitlab.String("opened"),
        Scope: gitlab.String("all"),
    }
    mrs, _, err := app.client.MergeRequests.ListMergeRequests(rOpts)
    if err != nil {
        log.Fatal(err)
    }

    spacer()
    fmt.Printf("\tOpen Merge Requests (%d)\n", len(mrs))
    spacer()

    for _, mr := range mrs {
        fmt.Printf(
            "%s\n\t> %s\n",
            mr.Title,
            mr.WebURL,
        )

        fmt.Printf(
            "\t# %s => %s\n\n",
            mr.SourceBranch,
            mr.TargetBranch,
        )
    }
}

func spacer() {
    fmt.Println("==========================================")
}
