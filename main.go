package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		ret    *clientv3.GetResponse
		err    error
	)
	config = clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}}
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	if ret, err = client.Get(context.TODO(), "a"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Kvs)
}
