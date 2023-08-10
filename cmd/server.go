package cmd

import (
	"database/sql"
	"log"

	"github.com/9bany/task/api"
	db "github.com/9bany/task/db/sqlc"
	httpclient "github.com/9bany/task/http_client"
	"github.com/9bany/task/util"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startServerCmd)
}

func startServer(cmd *cobra.Command, args []string) {
	var err error
	config := util.LoadConfig()

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database:", err)
	}

	store := db.New(conn)

	iframelyClient := httpclient.New(config.IframeURL)

	server, err := api.NewServer(config, store, iframelyClient)
	if err != nil {
		log.Fatalln("can not start server", err)
	}
	server.Start(config.ServerAddress)
}

var startServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Long:  ``,
	Run:   startServer,
}
