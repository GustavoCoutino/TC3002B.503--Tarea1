package hashmap

import (
	"errors"
	"slices"
)

type KeyVal[K any, V any] struct {
	Key   K
	Value V
}

type HashMap[K comparable, V any] struct {
	BucketSize int
	FilledSize int
	Bucket     [][]KeyVal[K, V]
	HashFunc   func(K) uint32
}

func (hm *HashMap[K, V]) _hash(key K) uint32 {
	return hm.HashFunc(key) % uint32(hm.BucketSize)
}

func New[K comparable, V any](BucketSize int, HashFunc func(K) uint32) *HashMap[K, V] {
	return &HashMap[K, V]{
		BucketSize: BucketSize,
		Bucket:     make([][]KeyVal[K, V], BucketSize),
		FilledSize: 0,
		HashFunc:   HashFunc,
	}
}

func (hm *HashMap[K, V]) Get(key K) (K, V, error) {
	if hm.BucketSize > 0 {
		index := hm._hash(key)
		bucket := hm.Bucket[index]
		for _, kv := range bucket {
			if kv.Key == key {
				return kv.Key, kv.Value, nil
			}
		}
	}
	var emptyKey K
	var emptyValue V
	return emptyKey, emptyValue, errors.New("key does not exist in hashmap")
}

func (hm *HashMap[K, V]) Insert(key K, value V) error {
	if hm.BucketSize == 0 {
		return errors.New("bucket size is 0")
	}

	if hm.FilledSize == hm.BucketSize {
		hm.BucketSize = hm.BucketSize * 2
		tempBucket := hm.Bucket
		hm.Bucket = make([][]KeyVal[K, V], hm.BucketSize)
		for _, i := range tempBucket {
			for _, j := range i {
				hash := hm._hash(j.Key)
				hm.Bucket[hash] = append(hm.Bucket[hash], KeyVal[K, V]{Key: j.Key, Value: j.Value})
			}
		}
	}
	index := hm._hash(key)
	bucket := hm.Bucket[index]
	for i, kv := range bucket {
		if kv.Key == key {
			hm.Bucket[index][i].Value = value
			return nil
		}
	}
	hm.Bucket[index] = append(hm.Bucket[index], KeyVal[K, V]{Key: key, Value: value})
	hm.FilledSize++
	return nil
}

func (hm *HashMap[K, V]) Remove(key K) error {
	index := hm._hash(key)
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

func (hm *HashMap[K, V]) Size() int {
	return hm.FilledSize
}
