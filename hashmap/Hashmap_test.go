package hashmap

import (
	"gustavocoutino/utils"
	"testing"
)

func TestHashmap_Get(t *testing.T) {
	t.Run("getting existing key in hashmap", func(t *testing.T) {
		hm := New[int](5, utils.FNVHash)
		hm.Insert("Google", 5)
		k, v, err := hm.Get("Google")
		if err != nil {
			t.Error("expected existing key")
		}
		wantedKey := 5
		wantedKeyValue := "Google"
		assertCorrectValue(t, k, wantedKeyValue)
		assertCorrectValue(t, v, wantedKey)
	})

	t.Run("returning error after searching inexistent key", func(t *testing.T) {
		hm := New[int](5, utils.FNVHash)
		hm.Insert("Google", 5)
		_, _, err := hm.Get("Microsoft")
		if err == nil {
			t.Error("expected key does not exist in hashmap error")
		}
	})

	t.Run("returning error after searching hashmap with size 0", func(t *testing.T) {
		hm := New[int](0, utils.FNVHash)
		_, _, err := hm.Get("Microsoft")
		if err == nil {
			t.Error("expected key does not exist in hashmap error")
		}
	})
}

func TestHashmap_Insert(t *testing.T) {
	t.Run("insert hashmap item in empty hashmap", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		hm.Insert("Microsoft", "California")
		k, v, err := hm.Get("Microsoft")
		if err != nil {
			t.Error("expected key does not exist in hashmap error")
		}
		wantedKey := "Microsoft"
		wantedKeyValue := "California"
		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, wantedKeyValue, v)
	})

	t.Run("insert hashmap item in hashmap of size 0", func(t *testing.T) {
		hm := New[string](0, utils.FNVHash)
		err := hm.Insert("Microsoft", "California")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("insert hashmap item in with existing key in hashmap updates the existing one", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Insert("Microsoft", "California")
		err = hm.Insert("Microsoft", "Colorado")
		k, v, err := hm.Get("Microsoft")
		if err != nil {
			t.Error("expected nil, got error")
		}
		wantedKey := "Microsoft"
		wantedValue := "Colorado"

		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, v, wantedValue)
	})

	t.Run("insert hashmap item in hashmap with non-repeating items", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Insert("Microsoft", "Washington")
		err = hm.Insert("Uber", "San Francisco")
		err = hm.Insert("Meta", "Menlo Park")
		err = hm.Insert("Apple", "Cupertino")
		err = hm.Insert("Netflix", "Los Gatos")
		k, v, err := hm.Get("Netflix")
		if err != nil {
			t.Error("expected nil, got error")
		}
		wantedKey := "Netflix"
		wantedValue := "Los Gatos"

		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, v, wantedValue)
	})

	t.Run("insert hashmap item in hashmap with bucket size already full", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Insert("Microsoft", "Washington")
		err = hm.Insert("Uber", "San Francisco")
		err = hm.Insert("Meta", "Menlo Park")
		err = hm.Insert("Apple", "Cupertino")
		err = hm.Insert("Netflix", "Los Gatos")
		err = hm.Insert("Cisco", "San Jose")
		k, v, err := hm.Get("Cisco")
		if err != nil {
			t.Error("expected nil, got error")
		}
		wantedKey := "Cisco"
		wantedValue := "San Jose"
		wantedLength := 6

		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, v, wantedValue)
		assertCorrectValue(t, hm.Size(), wantedLength)
	})
}

func TestHashmap_Remove(t *testing.T) {
	t.Run("removing from empty hashmap returns error", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Remove("Microsoft")
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("removing from hashmap with correct key removes value", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Insert("Microsoft", "Washington")
		err = hm.Remove("Microsoft")
		if err != nil {
			t.Error("Expected nil, got error")
		}
		wantedLength := 0
		assertCorrectValue(t, hm.Size(), wantedLength)
	})

	t.Run("removing from hashmap with inexistent key returns error", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		err := hm.Insert("Microsoft", "Washington")
		err = hm.Remove("Google")
		if err == nil {
			t.Error("Expected error, got nil")
		}
		wantedLength := 1
		assertCorrectValue(t, hm.Size(), wantedLength)
	})
}

func TestHashmap_Size(t *testing.T) {
	t.Run("checking size of empty hashmap", func(t *testing.T) {
		hm := New[string](5, utils.FNVHash)
		got := hm.Size()
		want := 0
		assertCorrectValue(t, got, want)
	})

	t.Run("checking size of hashmap with elements", func(t *testing.T) {
		hm := New[int](5, utils.FNVHash)
		hm.Insert("Hello", 1)
		hm.Insert("Goodbye", 2)
		hm.Insert("Later", 3)
		hm.Insert("Evening", 1)
		hm.Insert("Adieu", 4)
		hm.Insert("Hasta la vista", 2)
		got := hm.Size()
		want := 6
		assertCorrectValue(t, got, want)
	})
}

func assertCorrectValue[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
