package cmd

import (
	"database/sql"
	"log"

	db "github.com/9bany/task/db/sqlc"
	"github.com/9bany/task/util"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var newKey string

var keysCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create new key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(newKey) == 0 {
			log.Println("command: task key create --key='<key>'")
		}
		config := util.LoadConfig()
		conn, err := sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			log.Fatal("Can not connect to database:", err)
		}

		store := db.New(conn)
		key, err := store.CreateKey(cmd.Context(), newKey)
		if err != nil {
			log.Println("can not insert your ket into db")
		}
		log.Println("key insertd: ", key.Key)
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(keysCreateCmd)
	keysCreateCmd.PersistentFlags().StringVar(&newKey, "key", "", "key you want insert to db")
}

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "start server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("key functions")
	},
}
