package adt

import (
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()
	k := "a"
	v := "b"

	k2 := "c"
	v2 := "d"
	dict.Hset((*dictKey)(&k), (*dictValue)(&v))
	dict.Hset((*dictKey)(&k2), (*dictValue)(&v2))

	res := dict.Hget((*dictKey)(&k))
	if res != (*dictValue)(&v) {
		t.Error("hget is wrong ", res)
	}
}

func TestHash(t *testing.T) {
	dict := NewDict()
	k := "a"

	hash := dict.GetHash((*dictKey)(&k))
	if hash == uint64(0) {
		t.Error("hash is error ", hash)
	}

}
