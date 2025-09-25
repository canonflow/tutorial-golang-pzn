package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)

	/*
		=== RUN   TestConnection
		--- PASS: TestConnection (0.00s)
		PASS
	*/
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()

	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)

	/*
		=== RUN	  TestPing
		--- PASS: TestPing (0.01s)
		PASS
	*/
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Nathan Garzya", time.Second*5)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Nathan Garzya", result)

	time.Sleep(time.Second * 5)
	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)

	/*
		=== RUN   TestString
		--- PASS: TestString (5.02s)
		PASS
	*/
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "nathan")
	client.RPush(ctx, "names", "garzya")
	client.RPush(ctx, "names", "santoso")

	assert.Equal(t, "nathan", client.LPop(ctx, "names").Val())
	assert.Equal(t, "garzya", client.LPop(ctx, "names").Val())
	assert.Equal(t, "santoso", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")

	/*
		=== RUN   TestList
		--- PASS: TestList (0.02s)
		PASS
	*/
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "nathan")
	client.SAdd(ctx, "students", "nathan")
	client.SAdd(ctx, "students", "garzya")
	client.SAdd(ctx, "students", "garzya")
	client.SAdd(ctx, "students", "santoso")
	client.SAdd(ctx, "students", "santoso")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"nathan", "garzya", "santoso"}, client.SMembers(ctx, "students").Val())

	/*
		=== RUN   TestSet
		--- PASS: TestSet (0.02s)
		PASS
	*/
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "nathan"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "garzya"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "santoso"})

	assert.Equal(t, []string{"garzya", "santoso", "nathan"}, client.ZRange(ctx, "scores", 0, 2).Val())
	assert.Equal(t, "nathan", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "santoso", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "garzya", client.ZPopMax(ctx, "scores").Val()[0].Member)

	/*
		=== RUN   TestSortedSet
		--- PASS: TestSortedSet (0.02s)
		PASS
	*/
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "nathan")
	client.HSet(ctx, "user:1", "email", "nathan@example.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "nathan", user["name"])
	assert.Equal(t, "nathan@example.com", user["email"])

	client.Del(ctx, "user:1")

	/*
		=== RUN   TestHash
		--- PASS: TestHash (0.02s)
		PASS
	*/
}

func TestGeoPoint(t *testing.T) {
	// TODO: Menambahkan Geo Point
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 106.822702,
		Latitude:  -6.177590,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 106.820889,
		Latitude:  -6.174964,
	})

	// TODO: Mencari Geo Point
	assert.Equal(t, 0.3543, client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())
	fmt.Println(client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.821825,
		Latitude:   -6.175105,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)

	/*
		=== RUN   TestGeoPoint
		--- PASS: TestGeoPoint (0.02s)
		PASS
	*/
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "nathan", "garzya", "santoso")
	client.PFAdd(ctx, "visitors", "nathan", "canon", "flow")
	client.PFAdd(ctx, "visitors", "canon", "flow", "joko")

	assert.Equal(t, int64(6), client.PFCount(ctx, "visitors").Val())

	/*
		=== RUN   TestHyperLogLog
		--- PASS: TestHyperLogLog (0.02s)
		PASS
	*/
}

func TestPipeline(t *testing.T) {
	client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Nathan", time.Second*5)
		pipeliner.SetEx(ctx, "address", "Indonesia", time.Second*5)
		return nil
	})

	assert.Equal(t, "Nathan", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())

	/*
		=== RUN   TestPipeline
		--- PASS: TestPipeline (0.02s)
		PASS
	*/
}

func TestTransaction(t *testing.T) {
	client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Nathan", time.Second*5)
		pipeliner.SetEx(ctx, "address", "Surabaya", time.Second*5)
		return nil
	})

	assert.Equal(t, "Nathan", client.Get(ctx, "name").Val())
	assert.Equal(t, "Surabaya", client.Get(ctx, "address").Val())

	/*
		=== RUN   TestTransaction
		--- PASS: TestTransaction (0.02s)
		PASS
	*/
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "Nathan",
				"address": "Indonesia",
			},
		})
	}

	/*
		=== RUN   TestPublishStream
		--- PASS: TestPublishStream (0.02s)
		PASS
	*/
}

func TestCreateConsumer(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")

	/*
		=== RUN   TestCreateConsumer
		--- PASS: TestCreateConsumer (0.02s)
		PASS
	*/
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    time.Second * 5,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)

			/*
				map[address:Indonesia name:Nathan]
				map[address:Indonesia name:Nathan]
			*/
		}
	}

	/*
		=== RUN   TestGetStream
		--- PASS: TestGetStream (0.01s)
		PASS
	*/
}

func TestSubscribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()

	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)

		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}

	/*
		=== RUN   TestSubscribePubSub
		Hello 0
		Hello 1
		Hello 2
		Hello 3
		Hello 4
		Hello 5
		Hello 6
		Hello 7
		Hello 8
		Hello 9
		--- PASS: TestSubscribePubSub (3.19s)
		PASS
	*/
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i))
	}

	/*
		=== RUN   TestPublishPubSub
		--- PASS: TestPublishPubSub (0.02s)
		PASS
	*/
}
