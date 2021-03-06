package main

import (
	"golang.org/x/net/context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"localhost:2379"}
)

type ETCD struct {
	Key string
	Create_revision int32
	Mod_revision int32
	Version int32
	Value string
}


func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 设置 key1 的值为 value1
	key1, value1 := "key1", `value1`
	if resp, err := cli.Put(context.TODO(), key1, value1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// 设置 key1 的值为 value2, 并返回前一个值
	value2 := "value2"
	if resp, err := cli.Put(context.TODO(), key1, value2, clientv3.WithPrevKV()); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	if resp, err := cli.Get(context.TODO(), key1); err != nil {
		log.Println(err)
	} else {
		log.Println(resp.Header)
		log.Println(resp.Count)



		log.Println(resp.Kvs[0],
			string(resp.Kvs[0].Key),
			string(resp.Kvs[0].Value),
			resp.Kvs[0].CreateRevision,
			resp.Kvs[0].Lease,
			resp.Kvs[0].ModRevision,
			resp.Kvs[0].Version,
			resp.Kvs[0].Size())
		log.Println(resp.More)
		

	}
}
