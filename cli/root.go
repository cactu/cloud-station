package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	version bool
)

var RootCmd = cobra.Command{
	Use:   "cloud-station-cli",
	Long:  "cloud-station-cli 文件中转服务",
	Short: "cloud-station-cli ...",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("0.0.1")
			return nil
		}
		return errors.New("no flags find")
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&version, "cli version", "v", false, "the client version")
}
