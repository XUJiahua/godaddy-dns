package cmd

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/xujiahua/godaddy-dns/pkg/domain"
	"os"
	"path/filepath"
)

var fqdn string

var type_ string

var value string

// ensureCmd represents the ensure command
var ensureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "ensure a dns record exist",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		handleErr(err)

		dir := filepath.Join(home, ".godaddy")
		file := filepath.Join(dir, "secret.json")
		data, err := os.ReadFile(file)
		handleErr(err)

		var secret domain.Secret
		err = json.Unmarshal(data, &secret)
		handleErr(err)

		client := domain.New(secret)
		err = client.EnsureRecordSimplified(fqdn, type_, value)
		handleErr(err)
	},
}

func init() {
	rootCmd.AddCommand(ensureCmd)

	ensureCmd.PersistentFlags().StringVar(&fqdn, "fqdn", "", "Fully Qualified Domain Name, e.g., youtube.com, www.youtube.com")
	ensureCmd.PersistentFlags().StringVar(&type_, "type", "A", "A, CNAME")
	ensureCmd.PersistentFlags().StringVar(&value, "value", "", "IP address or CNAME FQDN")
}
