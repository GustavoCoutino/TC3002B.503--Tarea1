package hashmap

import (
	"testing"
)

func TestHashmap_Get(t *testing.T) {
	t.Run("getting existing key in hashmap", func(t *testing.T) {
		hm := New[string, int](5)
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
		hm := New[string, int](5)
		hm.Insert("Google", 5)
		_, _, err := hm.Get("Microsoft")
		if err == nil {
			t.Error("expected key does not exist in hashmap error")
		}
	})

	t.Run("returning error after searching hashmap with size 0", func(t *testing.T) {
		hm := New[string, int](0)
		_, _, err := hm.Get("Microsoft")
		if err == nil {
			t.Error("expected key does not exist in hashmap error")
		}
	})
}

func TestHashmap_Insert(t *testing.T) {
	t.Run("insert hashmap item in empty hashmap", func(t *testing.T) {
		hm := New[string, string](5)
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
		hm := New[string, string](0)
		err := hm.Insert("Microsoft", "California")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("insert hashmap item in with existing key in hashmap updates the existing one", func(t *testing.T) {
		hm := New[string, string](5)
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
		hm := New[string, string](5)
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
		hm := New[string, string](5)
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

	t.Run("insert hashmap item in hashmap with integers as keys", func(t *testing.T) {
		hm := New[int, string](5)
		err := hm.Insert(1, "Washington")
		err = hm.Insert(2, "San Francisco")
		err = hm.Insert(3, "Menlo Park")
		err = hm.Insert(4, "Cupertino")
		err = hm.Insert(5, "Los Gatos")
		err = hm.Insert(6, "San Jose")
		k, v, err := hm.Get(3)
		if err != nil {
			t.Error("expected nil, got error")
		}
		wantedKey := 3
		wantedValue := "Menlo Park"
		wantedLength := 6

		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, v, wantedValue)
		assertCorrectValue(t, hm.Size(), wantedLength)
	})

	t.Run("insert hashmap item in hashmap with booleans as keys", func(t *testing.T) {
		hm := New[bool, string](5)
		err := hm.Insert(true, "Washington")
		err = hm.Insert(false, "San Francisco")
		err = hm.Insert(true, "Menlo Park")
		k, v, err := hm.Get(true)
		if err != nil {
			t.Error("expected nil, got error")
		}
		wantedKey := true
		wantedValue := "Menlo Park"
		wantedLength := 2

		assertCorrectValue(t, k, wantedKey)
		assertCorrectValue(t, v, wantedValue)
		assertCorrectValue(t, hm.Size(), wantedLength)
	})
}

func TestHashmap_Remove(t *testing.T) {
	t.Run("removing from empty hashmap returns error", func(t *testing.T) {
		hm := New[string, string](5)
		err := hm.Remove("Microsoft")
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("removing from hashmap with correct key removes value", func(t *testing.T) {
		hm := New[string, string](5)
		err := hm.Insert("Microsoft", "Washington")
		err = hm.Remove("Microsoft")
		if err != nil {
			t.Error("Expected nil, got error")
		}
		wantedLength := 0
		assertCorrectValue(t, hm.Size(), wantedLength)
	})

	t.Run("removing from hashmap with inexistent key returns error", func(t *testing.T) {
		hm := New[string, string](5)
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
		hm := New[string, string](5)
		got := hm.Size()
		want := 0
		assertCorrectValue(t, got, want)
	})

	t.Run("checking size of hashmap with elements", func(t *testing.T) {
		hm := New[string, int](5)
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
