package redis

import (
	"fmt"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/gomodule/redigo/redis"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	c, err := redis.Dial("tcp", ":6379")
	testutil.AssertNil(t, err)
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(5)

	_, err = c.Do("LPUSH", "test.prepare", "a", "b", "c", "d", "e")
	testutil.AssertNil(t, err)

	go func() {
		for {
			reply, err := redis.String(c.Do("RPOPLPUSH", "test.prepare", "test.doing"))
			if err == nil {
				fmt.Println("Done with", reply)
				_, err := c.Do("LREM", "test.doing", 1, reply)
				testutil.AssertNil(t, err)
				wg.Done()
			} else {
				break
			}
		}
	}()

	wg.Wait()
}
