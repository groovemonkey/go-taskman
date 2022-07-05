package cmd

import (
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task",
	Long:  `Adds a task to the database, with a "todo" status.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		argString := strings.Join(args, " ")

		// Add string
		err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			err := b.Put([]byte(argString), []byte("todo"))
			return err
		})
		if err != nil {
			fmt.Println("ERROR: problem adding task to database: ", err, "; tried to add ", argString)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
