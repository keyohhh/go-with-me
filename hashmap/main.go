package main

import (
	"fmt"
)

type HashMap struct {
	Bucket []keyAndValue
	Size   int
}

type keyAndValue struct {
	Key   string
	Value string
}

func newHashMap() *HashMap {
	return &HashMap{
		Bucket: make([]keyAndValue, 16),
		Size:   0,
	}
}

func (h *HashMap) resize() {
	loadFactor := 0.80
	capacity := len(h.Bucket)
	expand := 0

	for _, r := range h.Bucket {
		if r.Key != "" {
			expand++
		}
	}

	size := float64(capacity) * loadFactor
	if float64(expand) > size {
		newCapacity := capacity * 2
		newBucket := make([]keyAndValue, newCapacity)

		for _, kv := range h.Bucket {
			if kv.Key != "" {
				index := h.hash(kv.Key) % int(newCapacity)
				newBucket[index] = kv
			}
		}

		h.Bucket = newBucket
	}
}

func (h *HashMap) hash(key string) int {
	hashCode := 0
	const primeNum = 31
	for _, r := range key {
		hashCode = primeNum*hashCode + int(r)
	}
	return hashCode
}

func (h *HashMap) set(key string, value string) {
	h.resize()
	index := h.hash(key) % len(h.Bucket)
	if h.Bucket[index].Key == key {
		h.Bucket[index].Value = value
		return
	}
	if h.Bucket[index].Key == "" {
		h.Bucket[index] = keyAndValue{Key: key, Value: value}
		h.Size++
		return
	}

	for i := 1; i < len(h.Bucket); i++ {
		newIndex := (index + i) % len(h.Bucket)
		if h.Bucket[newIndex].Key == "" {
			h.Bucket[newIndex] = keyAndValue{Key: key, Value: value}
			h.Size++
			return
		} else if h.Bucket[newIndex].Key == key {
			h.Bucket[newIndex].Value = value
			return
		}
	}
	panic("Hashmap is full!")
}

func (h *HashMap) get(key string) (interface{}, error) {
	index := h.hash(key) % len(h.Bucket)

	if h.Bucket[index].Key == key {
		return h.Bucket[index].Value, nil
	}
	return nil, fmt.Errorf("%s not found", key)
}

func (h *HashMap) has(key string) bool {
	index := h.hash(key) % len(h.Bucket)

	if h.Bucket[index].Key == key {
		return true
	}
	return false
}

func (h *HashMap) remove(key string) bool {
	index := h.hash(key) % len(h.Bucket)

	if key == h.Bucket[index].Key {
		h.Bucket[index] = keyAndValue{}
		h.Size--
		return true
	}
	return false
}

func (h *HashMap) length() int {
	fmt.Printf("Hashmap has a size of %d\n", h.Size)
	return h.Size
}

func (h *HashMap) clear() {
	h.Bucket = []keyAndValue{}
	h.Size = 0
}

func (h *HashMap) keys() []string {
	var key []string
	for _, v := range h.Bucket {
		if v.Key != "" {
			key = append(key, v.Key)
		}
	}
	return key
}

func (h *HashMap) values() []string {
	var value []string
	for _, v := range h.Bucket {
		if v.Value != "" {
			value = append(value, v.Value)
		}
	}
	return value
}

func (h *HashMap) print() {
	for i, v := range h.Bucket {
		fmt.Printf("%d %q\n", i, v)
	}
}

func main() {

}
