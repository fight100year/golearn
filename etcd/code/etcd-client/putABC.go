package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://172.168.10.116:2379", "http://172.168.10.119:2379", "http://172.168.10.135:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Put(ctx, "abc", "1122")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
