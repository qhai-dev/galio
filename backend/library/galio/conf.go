package galio

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	etcd "go.etcd.io/etcd/client/v3"
)

// type Option func(*options)

// type options struct {
// 	ctx context.Context

// 	prefix string
// }

func NewConfLoad() {
	ctx := context.Background()

	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{"14.103.230.114:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		fmt.Println(err, "client init")
	}

	res, err := client.Get(ctx, "rpc.service.account")
	if err != nil {
		fmt.Println(err, "client get")
	}

	v := viper.New()
	v.SetConfigType("yaml")

	for _, kv := range res.Kvs {
		v.ReadConfig(bytes.NewBuffer(kv.Value))
	}

	fmt.Printf("kind: %s", v.Get("kind"))
	fmt.Printf("apiVersion: %s", v.Get("apiVersion"))
	fmt.Printf("environments: %s", v.Get("spec.environments"))

}
