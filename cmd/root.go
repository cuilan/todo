package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// 配置文件
var cfgFile string

// 数据文件
var dataFile string

var RootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "Todo application for CLI.",
	Long: `Help you record to-do items and work more efficiently.
You're welcome`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which if defined here will be global for your application

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	// 数据文件路径
	RootCmd.PersistentFlags().StringVar(&dataFile,
		"datafile",
		home+string(os.PathSeparator)+".todo-cli.json",
		"data file to store todos")

	// 配置文件路径
	RootCmd.PersistentFlags().StringVar(&cfgFile,
		"config",
		home+string(os.PathSeparator)+".todo-cli.yaml",
		"config file")
}

// 初始化配置文件
func initConfig() {
	fmt.Println("cfg", cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".todo-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("todo-cli")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
