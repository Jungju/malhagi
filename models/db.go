package models

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// //MalhagiDB ...
// type MalhagiDB struct {
// 	DB *bolt.DB
// }

// func (mDB MalhagiDB) ListJawbons() ([]jawbone, error) {
// 	jawbones := []jawbone{}
// 	err := mDB.DB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("jawbones"))
// 		if b == nil {
// 			return nil
// 		}
// 		b.ForEach(func(k, v []byte) error {
// 			jawbone := &jawbone{}
// 			if err := json.Unmarshal(v, jawbone); err != nil {
// 				return err
// 			}
// 			jawbones = append(jawbones, *jawbone)
// 			return nil
// 		})
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jawbones, nil
// }

// func (mDB MalhagiDB) CreateJawbone(j jawbone) error {
// 	return mDB.DB.Update(func(tx *bolt.Tx) error {
// 		jawbones, err := tx.CreateBucketIfNotExists([]byte("jawbones"))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}

// 		jawboneBytes, err := json.Marshal(&j)
// 		if err != nil {
// 			return nil
// 		}

// 		return jawbones.Put([]byte(j.ID), jawboneBytes)
// 	})
// }

// func (mDB MalhagiDB) GetJawbone(token string) (*jawbone, error) {
// 	j := &jawbone{}
// 	err := mDB.DB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("jawbones"))
// 		if b == nil {
// 			return errInvalidToken
// 		}
// 		k := []byte(token)
// 		valueBytes := b.Get(k)
// 		if len(valueBytes) <= 0 {
// 			return errInvalidToken
// 		}
// 		if err := json.Unmarshal(b.Get(k), j); err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	logrus.Debugf("ID:%s", j.ID)
// 	logrus.Debugf("AccessToken:%s", j.Token.AccessToken)
// 	return j, nil
// }
