package main

import (
	"encoding/json"
	"errors"
	"hash/fnv"
	"log"
	"net/http"
	"strconv"

	"github.com/rrangith/golrucache/lrucache"
)

type shards []*lrucache.LRUCache

// Entry holds a key and a val, they need to be exported for json decoding to work
type Entry struct {
	Key interface{}
	Val interface{}
}

var numShards = 3
var shardCap = 5

// Fowler–Noll–Vo hash function
func getShard(key interface{}) int {
	h := fnv.New32a()

	switch v := key.(type) { //only support strings and ints for now
	case string:
		h.Write([]byte(v))
	case int:
		h.Write(intToBytes(v))
	}

	return int(h.Sum32() & uint32(numShards-1))
}

// helper function to quickly turn an int into a byte slice
func intToBytes(i int) []byte {
	var il = strconv.IntSize / 8
	b := make([]byte, il)
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
	b[3] = byte(i >> 24)
	if il == 8 {
		b[4] = byte(i >> 32)
		b[5] = byte(i >> 40)
		b[6] = byte(i >> 48)
		b[7] = byte(i >> 56)
	}
	return b
}

func (s *shards) get(key interface{}) interface{} {
	return (*s)[getShard(key)].Get(key)
}

func getHandler(s *shards) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Must be a get request", http.StatusMethodNotAllowed)
			return
		}

		key := r.URL.Path[5:]

		if key == "" {
			http.Error(w, "Must pass in a key", http.StatusBadRequest)
		}

		val := s.get(key)

		e := Entry{
			Key: key,
			Val: val,
		}

		json.NewEncoder(w).Encode(e)
	}
}

func (s *shards) set(e Entry) error {
	return (*s)[getShard(e.Key)].Set(e.Key, e.Val)
}

func setHandler(s *shards) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Must be a post request", http.StatusMethodNotAllowed)
			return
		}

		var e Entry

		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.set(e)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(e)
	}
}

func (s *shards) remove(e Entry) error {
	return (*s)[getShard(e.Key)].Remove(e.Key)
}

func removeHandler(s *shards) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Must be a delete request", http.StatusMethodNotAllowed)
			return
		}

		var e Entry // only key will be passed in, val will be nil

		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.remove(e)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(e)
	}
}

func main() {
	var shardCollection shards
	for i := 0; i < numShards; i++ {
		lru, err := lrucache.New(shardCap)

		if err != nil {
			errors.New("Something went wrong")
		}

		shardCollection = append(shardCollection, lru)
	}

	http.HandleFunc("/get/", getHandler(&shardCollection))
	http.HandleFunc("/set", setHandler(&shardCollection))
	http.HandleFunc("/remove", removeHandler(&shardCollection))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
