package watcher

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcd *clientv3.Client
var err error

func init(){
	etcd, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		fmt.Print("Error")
	}
}

func WacthKey(keyName string){
	watchChan := etcd.Watch(context.Background(), keyName, clientv3.WithPrefix())

	for response := range watchChan {
		for _, event := range response.Events{
			switch event.Type {
				case clientv3.EventTypePut:
					fmt.Printf("Key %s modified. New Value: %s\n", event.Kv.Key, event.Kv.Value)
				case clientv3.EventTypeDelete:
					fmt.Printf("Key %s modified. New Value: %s\n", event.Kv.Key, event.Kv.Value)
			}
		}
	}
}