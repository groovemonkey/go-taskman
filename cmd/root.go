package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "taskman is a small command-line task manager.",
	Long: `A small command-line task manager written in go: supports add, do, list, and cleanup commands.

(Warning: Just a toy)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go.mod.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Open the bolt db, ensure that our tasks bucket exists, and return a pointer to it
func initDB() *bolt.DB {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure tasks bucket exists
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db
}
