/*
Copyright © 2022 David Cohen dave@tutorialinux.com
*/
package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"go.mod/cmd"
)

func main() {
	// Open the bolt db
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd.Execute(db)
}

// EnsureBucketExists() creates a boltdb bucket, if necessary
func EnsureBucketExists(tx *bolt.Tx, bucketName string) error {
	_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		return fmt.Errorf("create bucket %s: %s", bucketName, err)
	}
	return nil
}
