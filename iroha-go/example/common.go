package example

import (
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

const (
	ToriiAddress    = "localhost:50051"
	DomainId        = "test"
	AdminAccountId  = "admin@test"
	UserAccountId   = "test@test"
	AdminPrivateKey = "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70"
	AssetId         = "coolcoin#test"
	UserPrivateKey  = "7e00405ece477bb6dd9b03a78eee4e708afc2f5bcdce399573a5958942f4a390"
)

func connect() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		ToriiAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
