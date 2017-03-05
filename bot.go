package main

import (
  "flag"
  "fmt"
  "strings"
  "github.com/bwmarrin/discordgo"
)

var (
  Token string
  BotID string
)

func init() { flag.StringVar(&Token, "t", "", "Bot Token")
  flag.Parse()
}

func main() {
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("Error creating Discord session, ", err)
    return
  }

  u, err := dg.User("@me")
  if err != nil {
    fmt.Println("Error obtaining bot details, ", err)
    return
  }
  BotID = u.ID

  dg.AddHandler(messageHandler)

  err = dg.Open()
  if err != nil {
    fmt.Println("Error opening connection, ", err)
    return
  }

  fmt.Println("ZigBot is now Online! Press Ctrl-C to exit.")
  <-make(chan struct{})
  return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == BotID {
    return
  }
  if strings.HasPrefix(m.Content, "$") {
    cmd_data := strings.SplitN(strings.TrimPrefix(m.Content, "$"), " ", 2)
    fmt.Println("CMD: ", cmd_data)
    switch cmd_data[0] {
      case "help":
        _, _ = s.ChannelMessageSend(m.ChannelID, "Available commands:\n``` - $help\n - $calc expression\n - $dice\n - $magic8ball\n-\n-\n```")
      case "calc":
        resp := calc(cmd_data[1]) 
        _, _ = s.ChannelMessageSend(m.ChannelID, resp)
      case "dice":
        resp := dice()
        _, _ = s.ChannelMessageSend(m.ChannelID, resp)
     case "magic8ball":
        resp := magic8Ball()
        _, _ = s.ChannelMessageSend(m.ChannelID, resp)
      default:
        _, _ = s.ChannelMessageSend(m.ChannelID, "Unknow command. Use $help for more information!")
    }
  }
}
