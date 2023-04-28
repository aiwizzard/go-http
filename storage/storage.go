package storage

import "github.com/aiwizzard/go-http/types"

type Storage interface {
    Get(int) *types.User
}

