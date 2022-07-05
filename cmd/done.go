package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "lists tasks that have been done",
	Long:  `Sub-command of 'list' -- lists only tasks that have been done.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		err := db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				if string(v) == "done" {
					fmt.Printf("DONE: %s\n", k)
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
	listCmd.AddCommand(doneCmd)
}
