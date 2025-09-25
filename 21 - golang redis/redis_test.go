package main

import (
	"context"
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
