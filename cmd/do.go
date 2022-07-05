package cmd

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as done",
	Long:  `Changes status for a given task from "todo" to "done".`,
	Run: func(cmd *cobra.Command, args []string) {
		itemArgStr, err := strconv.Atoi(args[0])
		if err != nil || itemArgStr == 0 {
			panic("Error: invalid Task ID.")
		}
		itemIdx := itemArgStr - 1

		db := initDB()
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte("tasks"))
			c := b.Cursor()

			// Dynamically number tasks with status "todo"
			i := 0
			for k, v := c.First(); k != nil; k, v = c.Next() {
				if string(v) == "todo" {
					fmt.Println(string(k), "has itemIdx", itemIdx, "and i=", i)
					// Have we arrived at the correct item?
					if i == itemIdx {
						fmt.Printf("You have completed the '%s' task.\n", string(k))
						// Set the task to done
						err := b.Put(k, []byte("done"))
						if err != nil {
							fmt.Println("Error setting task to done.")
							return err
						}
						return nil
					} else {
						// Increment visible counter only for "todo" tasks
						i++
					}
				}
			}

			// // Initial attempt (optimized a bit too early)
			// // If we weren't dynamically numbering tasks, this would be the way to go.
			// c.First()
			// // Stop just before the item we actually want
			// for i := 0; i < itemIdx-1; i++ {
			// 	c.Next()
			// }
			// // Grab the item and actually allocate memory
			// k, _ := c.Next()
			// fmt.Printf("You have completed the '%s' task.", k)

			return nil
		})
		if err != nil {
			fmt.Println("ERROR: something went wrong while trying to view the DB: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
