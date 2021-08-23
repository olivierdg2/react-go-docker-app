package main

import (
	"fmt"
	"time"

	"github.com/olivierdg2/react-go-docker-app/go/pkg/cmd/server"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	server.kv = clientv3.NewKV(cli)
	if err != nil {
		fmt.Printf("%v", err)
	}
	server.HandleRequests()
	defer cli.Close()
}
