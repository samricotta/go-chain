package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	se:    "blockchain-cli",
	Short: "Blockchain CLI is a tool for interacting with the blockchain",
	Long: `A longer description that spans multiple lines and likely contains 
examples and usage of using your application.`,
}

// execute adds all child commands to the root command and sets flags appropriately.
// this is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// define your flags and configuration settings.
	// persistent flags defined here,
	// will be global for application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockchain-cli.yaml)")

	// Cobra supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	Execute()
}
