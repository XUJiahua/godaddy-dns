package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/xujiahua/godaddy-dns/pkg/domain"
	"os"
	"path/filepath"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "key",
		Prompt:   &survey.Input{Message: "API Key:"},
		Validate: survey.Required,
	},
	{
		Name:     "secret",
		Prompt:   &survey.Input{Message: "API Secret:"},
		Validate: survey.Required,
	},
}

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure godaddy secret",
	Run: func(cmd *cobra.Command, args []string) {
		var answers domain.Secret

		err := survey.Ask(qs, &answers)
		handleErr(err)

		// Find home directory.
		home, err := homedir.Dir()
		handleErr(err)

		dir := filepath.Join(home, ".godaddy")

		err = os.MkdirAll(dir, 0700)
		handleErr(err)

		data, err := json.Marshal(answers)
		handleErr(err)

		file := filepath.Join(dir, "secret.json")

		err = os.WriteFile(file, data, 0600)
		handleErr(err)

		fmt.Println("OK")
	},
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
