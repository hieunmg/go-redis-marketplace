package cmd

import (
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat server",
	Run: func(cmd *cobra.Command, args []string) {
		// server, err := wire.InitializeChatServer("chat")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
