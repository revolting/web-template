package db

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	//"errors"
	"fmt"

	"log"

	"github.com/boltdb/bolt"
	"github.com/nu7hatch/gouuid"
)

const dbPath = "./db/leaves.db"

type Profile struct {
	Uid			string
	Name		string
	Phone		string
}

func init() {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Profile"))
		if (err != nil) {
			log.Fatal(err)
		}
		return nil
	})

	if (err != nil) {
		log.Fatal(err)
	}
}

func GetProfile(phone string) (*Profile, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var profile *Profile

	err = db.View(func(tx *bolt.Tx) error {
		p := tx.Bucket([]byte("Profile"))
		acct := p.Get([]byte(phone))

		err = json.Unmarshal(acct, &profile)
		if (err != nil) {
			return err
		}
		return nil
	})

	if (err != nil) {
		return nil, err
	}

	return profile, nil
}

func UpdateProfile(uid uuid.UUID, name string, phone string) (*Profile, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	id := hex.EncodeToString(uid[:])
	hash := md5.Sum([]byte(phone))
	phoneHash := hex.EncodeToString(hash[:])
	profile := &Profile{Uid: id, Name: name, Phone: phoneHash}

	encoded, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		p := tx.Bucket([]byte("Profile"))

		return p.Put([]byte(profile.Phone), encoded)
	})

	if (err != nil) {
		return nil, err
	}
	println("returning profile ", profile)
	return profile, err
}
