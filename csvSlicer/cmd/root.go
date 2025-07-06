package cmd

import (
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
    Use:   "csvSlicer",
    Short: "csvSlcer trims large csv files",
    Long: "csvSlcer trims large csv files",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing csvSlicer '%is'\n", err)
        os.Exit(1)
    }
}
