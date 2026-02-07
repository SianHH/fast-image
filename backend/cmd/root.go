package cmd

import (
	"fast-image/bootstrap"
	"fast-image/global"
	"fast-image/pkg/signal"
	"fast-image/pkg/utils"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().String("basepath", "", "--basepath /workspace/")
	RootCmd.PersistentFlags().String("log-level", "info", "--log-level debug|info|warn|error|fatal")
	RootCmd.PersistentFlags().Bool("dev", false, "--dev")
}

func persistentFlags(cmd *cobra.Command) (basepath string, mode bool, loglevel string) {
	basepath, _ = cmd.Flags().GetString("basepath")
	mode, _ = cmd.Flags().GetBool("dev")
	loglevel, _ = cmd.Flags().GetString("log-level")
	return basepath, mode, loglevel
}

var RootCmd = cobra.Command{
	Use:   "",
	Short: "./fast-image",
	Run: func(cmd *cobra.Command, args []string) {
		basepath, mode, loglevel := persistentFlags(cmd)
		global.SetBasePath(basepath)
		global.SetLoggerLevel(loglevel)
		global.SetAppMode(utils.TrinaryOperation(mode, global.APP_MODE_DEV, global.APP_MODE_PROD))

		bootstrap.InitLogger()
		bootstrap.InitConfig()
		bootstrap.InitJwt()
		bootstrap.InitBadgerDB()
		bootstrap.InitRouter()
		bootstrap.InitServer()
		bootstrap.InitTodo()
		bootstrap.InitTask()

		<-signal.Free()
		bootstrap.Release()
	},
}
