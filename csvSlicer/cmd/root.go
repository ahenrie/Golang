var rootCmd = &cobra.Command{
  Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
