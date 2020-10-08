/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"strings"

	uc "gogo/app/cli/usecase/messenger/slack"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var slackChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "A brief description of your command",
	Long:  `.`,
	Run:   send,
}

func send(cmd *cobra.Command, args []string) {
	msg := strings.Join(args, " ")
	fmt.Println("begin slack send")

	token := viper.GetString("slack.chat.token")
	channel := viper.GetString("slack.chat.channel")

	slack := uc.NewChat(token, channel)
	slack.Send(msg)
	fmt.Println("finish slack send")
}

func init() {
	slackCmd.AddCommand(slackChatCmd)
}
