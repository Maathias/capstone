package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func Connect(service, addr, pass string) *redis.Client {
	
	var dbn int
	
	if service == "edge" {
		dbn = 1
	} else if service == "api" {
		dbn = 2
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       dbn,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to redis")
		fmt.Println(err)
		return nil
	}

	fmt.Println("Connected to redis")
	return rdb
}

func Set(key string, value interface{}) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("Error saving to redis")
		fmt.Println("key: ", key, "value: ", value)
	}
}

func Get( key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting from redis")
		fmt.Println("key: ", key)
		return ""
	}
	return val
}

func GeoAdd( key string, longitude float64, latitude float64, member string) {
	err := rdb.GeoAdd(ctx, key, &redis.GeoLocation{
		Name:      member,
		Longitude: longitude,
		Latitude:  latitude,
	}).Err()
	if err != nil {
		fmt.Println("Error adding geo location to redis")
		fmt.Println("key: ", key, "longitude: ", longitude, "latitude: ", latitude, "member: ", member)
		fmt.Println(err)
	}
}

func GeoDist( key string, member1 string, member2 string) float64 {
	dist, err := rdb.GeoDist(ctx, key, member1, member2, "km").Result()
	if err != nil {
		fmt.Println("Error getting distance from redis")
		fmt.Println("key: ", key, "member1: ", member1, "member2: ", member2)
		return 0
	}
	return dist
}

func GeoRadius( key string, longitude float64, latitude float64, radius float64) []redis.GeoLocation {
	loc := &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        "km",
		WithCoord:   false,
		WithDist:    false,
		WithGeoHash: false,
		Count:       0,
		Sort:        "ASC",
		Store:       "",
		StoreDist:   "",
	}
	members, err := rdb.GeoRadius(ctx, key, longitude, latitude, loc).Result()
	if err != nil {
		fmt.Println("Error getting radius from redis")
		fmt.Println("key: ", key, "longitude: ", longitude, "latitude: ", latitude, "radius: ", radius)
		return nil
	}
	return members
}

func SortedAdd( key string, score int64, member string) {
	err := rdb.ZAdd(ctx, key, redis.Z{
		Score:  float64(score),
		Member: member,
	}).Err()
	if err != nil {
		fmt.Println("Error adding to sorted set")
		fmt.Println("key: ", key, "score: ", score, "member: ", member)
	}
}

func SortedRange( key string, start int64, stop int64) []string {
	vals, err := rdb.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		fmt.Println("Error getting sorted set range")
		fmt.Println("key: ", key, "start: ", start, "stop: ", stop)
		return nil
	}
	return vals
}