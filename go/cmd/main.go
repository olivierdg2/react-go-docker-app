package main

import (
	"https://github.com/olivierdg2/react-go-docker-app/go/pkg/cmd/server"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	kv = clientv3.NewKV(cli)
	if err != nil {
		fmt.Printf("%v", err)
	}
	handleRequests()
	defer cli.Close()
}
