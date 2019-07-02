package main

import (
    "log"
    "fmt"

    "github.com/xanzy/go-gitlab"
    "github.com/spf13/viper"
    "github.com/fatih/color"
)

type App struct {
    client *gitlab.Client
}

func main() {
    // configuration
    viper.AddConfigPath(".")
    viper.AddConfigPath("$HOME/.ohmygitlab")
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
        WIP: gitlab.String("no"),
    }
    mrs, _, err := app.client.MergeRequests.ListMergeRequests(rOpts)
    if err != nil {
        log.Fatal(err)
    }

    spacer()
    headline := color.New(color.FgYellow, color.Bold)
    headline.Printf("Open Merge Requests (%d)\n", len(mrs))
    spacer()

    subheadline := color.New(color.FgGreen)
    author := color.New(color.FgCyan)
    info := color.New(color.FgWhite)
    for _, mr := range mrs {
        subheadline.Printf(
            "\n%s\n",
            mr.Title,
        )

        author.Printf(
            "\tby %s\n",
            mr.Author.Username,
        )

        info.Printf(
            "\tMerge %s into %s\n",
            mr.SourceBranch,
            mr.TargetBranch,
        )

        info.Printf(
            "\t%s\n",
            mr.WebURL,
        )
    }
}

func spacer() {
    color.White("==========================================")
}
