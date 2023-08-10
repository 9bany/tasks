package cmd

import (
	"context"
	"database/sql"
	"log"

	db "github.com/9bany/task/db/sqlc"
	"github.com/9bany/task/util"
	"github.com/fatih/structs"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var keyString string

var keysCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create new key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(keyString) == 0 {
			log.Println("command: task keys create --key='<key>'")
		}
		config := util.LoadConfig()
		conn, err := sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			log.Fatal("can not connect to database:", err)
		}

		store := db.New(conn)
		key, err := store.CreateKey(cmd.Context(), keyString)
		if err != nil {
			log.Println("can not insert key into db")
		}
		log.Println("key inserted: ", key.Key)
	},
}

var getKeyCmd = &cobra.Command{
	Use:   "get",
	Short: "get key info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(keyString) == 0 {
			log.Println("command: task keys get --key='<key>'")
		}
		config := util.LoadConfig()
		conn, err := sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			log.Fatal("can not connect to database:", err)
		}

		store := db.New(conn)
		key, err := store.GetKey(context.Background(), keyString)
		if err != nil {
			log.Println("can not get your key info")
		}
		log.Println("key info: ", structs.Map(key))
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(keysCreateCmd)
	keysCmd.AddCommand(getKeyCmd)
	keysCreateCmd.PersistentFlags().StringVar(&keyString, "key", "", "")
	getKeyCmd.PersistentFlags().StringVar(&keyString, "key", "", "")
}

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "start server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("key functions")
	},
}
