package db

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
	"github.com/nu7hatch/gouuid"
)

const dbPath = "../../../db/leaves.db"

type Profile struct {
	Uid			*uuid.UUID
	Name		string
	Phone		string
}

func UpdateProfile(uid *uuid.UUID, name string, phone string) error {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	profile := &Profile{Uid: uid, Name: name, Phone: phone}

	err = db.Update(func(tx *bolt.Tx) error {
		p, err := tx.CreateBucketIfNotExists([]byte("Profile"))
		if (err != nil) {
			return err
		}

		encoded, err := json.Marshal(profile)
		if err != nil {
			return err
		}
		println("saved profile ", profile.Phone)
		return p.Put([]byte(profile.Phone), encoded)
	})
	return err
}
