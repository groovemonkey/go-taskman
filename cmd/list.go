package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks that still have a status of 'todo'",
	Long:  `list all tasks that still have a status of 'todo'. Ignore all "done" tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		err := db.View(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte("tasks"))
			c := b.Cursor()

			// BoltDB items are stored deterministically so we can just use a counter to number them
			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				if string(v) == "todo" {
					fmt.Printf("%d. %s\n", i, k)
					// Increment visible counter
					i++
				}
			}

			return nil
		})
		if err != nil {
			fmt.Println("ERROR: something went wrong while trying to view the DB: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
