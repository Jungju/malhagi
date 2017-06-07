package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

//Sentence ...
//bee generate model sentence -fields="id:int,created_at:datetime,verb:string,korean:string,past:string"
type Sentence struct {
	CreatedAt time.Time
	Verb      string
	Korean    string
	Past      string
}

// AddSentence insert a new Sentence into database and returns
// last inserted Id on success.
func AddSentence(m *Sentence) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sentences"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		sentenceBytes, err := json.Marshal(m)
		if err != nil {
			return nil
		}

		return b.Put([]byte(m.Verb), sentenceBytes)
	})
}

// GetAllSentence retrieves all Sentence matches certain condition. Returns empty list if
// no records exist
func GetAllSentence() ([]Sentence, error) {
	sentences := []Sentence{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sentences"))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			sentence := &Sentence{}
			if err := json.Unmarshal(v, sentence); err != nil {
				return err
			}
			sentences = append(sentences, *sentence)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return sentences, nil
}

// UpdateSentenceById ...
func UpdateSentenceById(m *Sentence) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sentences"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		bs, err := json.Marshal(&m)
		if err != nil {
			return err
		}

		return b.Put([]byte(m.Verb), bs)
	})
}

// DeleteSentence deletes Sentence by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSentence(verb string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sentences"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return b.Delete([]byte(verb))
	})
}
