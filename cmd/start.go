package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pepodev/autoauth/internal/utils"
	"github.com/pepodev/xlog"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/pkg/auth"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:        "start",
	Short:      "Start autoauth",
	Long:       message.GetWelcome(),
	SuggestFor: []string{"run"},
	Example:    "autoauth start -d -f config.yml -k my_s3cr3t_k3y",
	Run: func(cmd *cobra.Command, args []string) {
		isDetached, _ := cmd.Flags().GetBool("detach")
		fileName, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")
		viper.SetDefault("key", key)

		if isDetached {
			utils.StartInDetachMode()
		}

		preset := auth.LoadPresetFromPath(utils.GetWorkingDirectory(), fileName)

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		auth.StartAutoLogin(preset)

		xlog.Infof("os %v AutoAuth stopped", <-sigs)
	},
}

func init() {
	startCmd.Flags().BoolP("detach", "d", false, "start autoauth in the detach mode")
	startCmd.Flags().StringP("file", "f", "autoauth.yml", "file of preset you want to start auto auth")
	startCmd.Flags().StringP("key", "k", "", "key to decrypt your data from config file")
}
