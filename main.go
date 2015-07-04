package main

import (
    "fmt"

    "github.com/nlopes/slack"
)

func main() {
    api := slack.New("TOKEN")
    params := slack.PostMessageParameters{}
    channelId, timestamp, err := api.PostMessage("#gnoia", "Eder viadinho by golang.", params)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    fmt.Printf("Message successfully sent to channel %s at %s", channelId, timestamp)
}