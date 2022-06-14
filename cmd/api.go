/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/kristiansantos/learning/src/config/environments"
	"github.com/kristiansantos/learning/src/config/server"
	"github.com/kristiansantos/learning/src/shared/providers/logger"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var (
	port        int
	addr        string
	environment string
	brandName   string
	version     string
	apiCmd      = &cobra.Command{
		Use:   "api",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := godotenv.Load("./src/config/environments/.env." + environment); err != nil {
				panic(err)
			}

			os.Getenv("ENV")

			cfg, err := environments.ReadEnvironments(environment, brandName, version)

			if err != nil {
				panic(err)
			}

			svr := server.New(addr, port)
			log := logger.New()
			svr.Run(cfg, log)

			chanExit := make(chan os.Signal, 1)
			signal.Notify(chanExit, os.Interrupt)
			<-chanExit
		},
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)

	// Get start server options
	apiCmd.PersistentFlags().IntVarP(&port, "port", "p", 3000, "The -p option specified port HTTP server")
	apiCmd.PersistentFlags().StringVarP(&addr, "address", "a", "127.0.0.1", "The -b option binds specified IP, by default it is localhost")
	apiCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "development", "The -e option specified the environment")
	apiCmd.PersistentFlags().StringVarP(&brandName, "brand", "b", os.Getenv("BRAND"), "The -b option specified brand to deploy")
	apiCmd.PersistentFlags().StringVarP(&version, "version", "v", os.Getenv("VERSION"), "The -v option specified version to deploy")
}
