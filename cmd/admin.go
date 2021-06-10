package cmd

import (
	"os"
	"os/signal"
	"syscall"

	admin "github.com/Hamster601/flashSale/application/http/admin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Seckill admin server.",
	Long:  `Seckill admin server.`,
	Run: func(cmd *cobra.Command, args []string) {
		onExit := make(chan error)
		go func() {
			if err := admin.Run(); err != nil {
				logrus.Error(err)
				onExit <- err
			}
			close(onExit)
		}()
		onSignal := make(chan os.Signal)
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-onSignal:
			logrus.Info("exit by signal ", sig)
			admin.Exit()
		case err := <-onExit:
			logrus.Info("exit by error ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
