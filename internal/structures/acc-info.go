package structures

import "time"

type CoinlistAccs struct {
	Id           int
	Password     string
	CID          int64
	StatusVerify string
	Email        string
	CreatedAt    time.Time
}
