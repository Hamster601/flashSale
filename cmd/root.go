package cmd

import (
	"fmt"
	"os"

	"github.com/Hamster601/flashSale/application/infrastructures/cluster"
	log "github.com/Hamster601/flashSale/application/infrastructures/logger"
	"github.com/Hamster601/flashSale/application/infrastructures/pkg/etcd"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flashSale",
	Short: "flashSale server.",
	Long:  `flashSale server.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logrus.Info(cfgFile)
}

func init() {
	cobra.OnInitialize(initConfig)

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&cfgFile, "config", "c", "./config/conf.toml", "config file (default is $HOME/.conf.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".conf" (without extension).
		viper.AddConfigPath(home)
		// 设置配置文件名
		viper.SetConfigName(".conf")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
	if err := etcd.Init(); err != nil {
		panic(err)
	}
	log.InitLogger()
	if err := cluster.WatchClusterConfig(); err != nil {
		panic(err)
	}
}
