package service_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/livekit/livekit-server/pkg/service"
	"github.com/stretchr/testify/require"
	"testing"
)

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func TestIsValidDomain(t *testing.T) {
	list := map[string]bool{
		"turn.myhost.com": true,
		"turn.google.com": true,
		"https://host.com": false,
		"turn://host.com": false,
	}
	for key, result := range list {
		service.IsValidDomain(key)
		require.Equal(t, service.IsValidDomain(key), result)
	}
}