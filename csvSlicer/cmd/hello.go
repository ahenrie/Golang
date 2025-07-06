package cmd


import (
	"github.com/spf13/cobra"
	"fmt"
)

var HelloCommand = &cobra.Command{
	Use:	"say hello",
	Run: func(cmd *cobra.Command) {
		fmt.Println("Hello, there!")
	},
}

func init() {
	rootCmd.HelloCommand()
}
