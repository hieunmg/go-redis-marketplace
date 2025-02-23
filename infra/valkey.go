package infra

import (
	"context"
	"encoding/json"
	"go-redis-marketplace/pkg/common"
	"go-redis-marketplace/pkg/config"

	"github.com/valkey-io/valkey-go"
)

var (
	ValkeyClient valkey.Client
)

type ValkeyCache interface {
	Get(ctx context.Context, key string, dst interface{}) (bool, error)
	// Set(ctx context.Context, key string, val interface{}) error
	// Delete(ctx context.Context, key string) error
	// HGet(ctx context.Context, key, field string, dst interface{}) (bool, error)
	// HMGet(ctx context.Context, key string, fields []string) ([]interface{}, error)
	// HGetAll(ctx context.Context, key string) (map[string]string, error)
	// HSet(ctx context.Context, key string, values ...interface{}) error
	// HDel(ctx context.Context, key, field string) error
	// RPush(ctx context.Context, key string, val interface{}) error
	// LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	// Publish(ctx context.Context, topic string, payload interface{}) error
	// ZPopMinOrAddOne(ctx context.Context, key string, score float64, member interface{}) (bool, string, error)
	// ZRemOne(ctx context.Context, key string, member interface{}) error
	// ExecPipeLine(ctx context.Context, cmds *[]RedisCmd) error
}

type ValkeyCacheImplement struct {
	client valkey.Client
}

type Valkey int

const (
	// DELETE represents delete operation
	VALKEY_DELETE Valkey = iota
	VALKEY_HSETONE
	VALEY_RPUSH
)

// ValkeyPayload is a abstract interface for payload type
type ValkeyPayload interface {
	Payload()
}

// ValkeyDeletePayload is the payload type of delete method
type ValkeyDeletePayload struct {
	ValkeyPayload
	Key string
}

type ValkeyHsetOnePayload struct {
	ValkeyPayload
	Key   string
	Field string
	Val   interface{}
}

type ValkeyRpushPayload struct {
	ValkeyPayload
	Key string
	Val interface{}
}

func NewValkeyClient(conf *config.Config) (valkey.Client, error) {
	ValkeyClient, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: common.GetServerAddrs(conf.Valkey.Addrs),
		Password:    conf.Valkey.Password,
	})
	if err != nil {
		return nil, err
	}

	return ValkeyClient, nil
}

func NewValkeyCacheImplement(valkeyClient valkey.Client) ValkeyCache {
	return &ValkeyCacheImplement{client: valkeyClient}
}

func (vc *ValkeyCacheImplement) Get(ctx context.Context, key string, dst interface{}) (bool, error) {
	val, err := vc.client.Do(ctx, vc.client.B().Get().Key(key).Build()).ToString()
	if err == valkey.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		json.Unmarshal([]byte(val), dst)
	}
	return true, nil
}
