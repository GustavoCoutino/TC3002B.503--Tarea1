package hashmap

import (
	"errors"
	"slices"
)

type KeyVal[T any] struct {
	Key   string
	Value T
}

type HashMap[T any] struct {
	BucketSize int
	FilledSize int
	Bucket     [][]KeyVal[T]
	HashFunc   func(string) uint32
}

func (hm *HashMap[T]) Hash(key string) uint32 {
	return hm.HashFunc(key) % uint32(hm.BucketSize)
}

func New[T any](BucketSize int, HashFunc func(string) uint32) *HashMap[T] {
	return &HashMap[T]{
		BucketSize: BucketSize,
		Bucket:     make([][]KeyVal[T], BucketSize),
		FilledSize: 0,
		HashFunc:   HashFunc,
	}
}

func (hm *HashMap[T]) Get(key string) (string, T, error) {
	if hm.BucketSize > 0 {
		index := hm.Hash(key)
		bucket := hm.Bucket[index]
		for _, kv := range bucket {
			if kv.Key == key {
				return kv.Key, kv.Value, nil
			}
		}
	}
	var emptyValue T
	return "", emptyValue, errors.New("key does not exist in hashmap")
}

func (hm *HashMap[T]) Insert(key string, value T) error {
	if hm.BucketSize == 0 {
		return errors.New("bucket size is 0")
	}

	if hm.FilledSize == hm.BucketSize {
		hm.BucketSize = hm.BucketSize * 2
		tempBucket := hm.Bucket
		hm.Bucket = make([][]KeyVal[T], hm.BucketSize)
		for _, i := range tempBucket {
			for _, j := range i {
				hash := hm.Hash(j.Key)
				hm.Bucket[hash] = append(hm.Bucket[hash], KeyVal[T]{Key: j.Key, Value: j.Value})
			}
		}
	}
	index := hm.Hash(key)
	bucket := hm.Bucket[index]
	for i, kv := range bucket {
		if kv.Key == key {
			bucket[i].Value = value
			return nil
		}
	}
	hm.Bucket[index] = append(hm.Bucket[index], KeyVal[T]{Key: key, Value: value})
	hm.FilledSize++
	return nil
}

func (hm *HashMap[T]) Remove(key string) error {
	index := hm.Hash(key)
	bucket := hm.Bucket[index]
	if len(bucket) > 0 {
		for i, kv := range bucket {
			if kv.Key == key {
				hm.Bucket[index] = slices.Delete(hm.Bucket[index], i, i+1)
				hm.FilledSize--
				return nil
			}
		}
	}
	return errors.New("Key was not found in hashmap")
}

func (hm *HashMap[T]) Size() int {
	return hm.FilledSize
}
