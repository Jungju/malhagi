package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

//Game ...
//bee generate model game -fields="id:int,created_at:datetime,ended:bool,point:int"
type Game struct {
	Id        int
	CreatedAt time.Time
	Ended     bool
	Point     int
}

// AddGame insert a new Game into database and returns
// last inserted Id on success.
func AddGame(g *Game) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("games"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		gameBytes, err := json.Marshal(g)
		if err != nil {
			return nil
		}

		return b.Put(itob(g.Id), gameBytes)
	})
}

// GetGameById retrieves Game by Id. Returns error if
// Id doesn't exist
func GetGameById(id int) (*Game, error) {
	game := &Game{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sentences"))
		if b == nil {
			return nil
		}
		bytes := b.Get(itob(id))

		return json.Unmarshal(bytes, game)
	})
	return game, err
}

// GetAllGame retrieves all Game matches certain condition. Returns empty list if
// no records exist
func GetAllGame() ([]Game, error) {
	games := []Game{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			game := &Game{}
			if err := json.Unmarshal(v, game); err != nil {
				return err
			}
			games = append(games, *game)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return games, nil
}

// UpdateGameById ...
func UpdateGameById(m *Game) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sentences"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		bytes, err := json.Marshal(&m)
		if err != nil {
			return err
		}

		return b.Put(itob(m.Id), bytes)
	})
}
