package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

func FNVHash[K comparable](key K) uint32 {
	h := fnv.New32a()

    if s, ok := any(key).(string); ok {
		h.Write([]byte(s))
		return h.Sum32()
	}

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, key); err != nil {
		h.Write([]byte(fmt.Sprintf("%v", key)))
		return h.Sum32()
	}
	h.Write(buf.Bytes())

    return h.Sum32()
}
