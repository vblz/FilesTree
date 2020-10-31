package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"os"
)

var bucketName = []byte("fileInfos")

func Write(dbPath string, data map[string]os.FileInfo) error {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return fmt.Errorf("open error: %w", err)
	}

	defer db.Close()

	return db.Update(func(tx *bbolt.Tx) error {
		// just to clear, doesn't matter if it doesn't exist
		_ = tx.DeleteBucket(bucketName)

		b, err := tx.CreateBucket(bucketName)
		if err != nil {
			return err
		}

		for k, v := range data {
			fileInfo := CopyToFileInfo(v)
			bytes, err := json.Marshal(fileInfo)
			if err != nil {
				return err
			}

			err = b.Put([]byte(k), bytes)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func Read(dbPath string) (map[string]os.FileInfo, error) {
	db, err := bbolt.Open(dbPath, 600, nil)
	if err != nil {
		return nil, fmt.Errorf("open error: %w", err)
	}

	defer db.Close()

	result := make(map[string]os.FileInfo)

	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)

		if b == nil {
			return fmt.Errorf("wrong database, empty bucket")
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var fileInfo FileInfo
			err := json.Unmarshal(v, &fileInfo)
			if err != nil {
				return fmt.Errorf("deserializetion error: %w", err)
			}

			result[string(k)] = fileInfo
		}

		return nil
	})

	return result, err
}
