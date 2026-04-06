package utils

import "testing"

func TestUtils_FNVHash(t *testing.T){
	t.Run("returns hash for key", func(t *testing.T){
		hash := FNVHash("Hello")
		var wanted uint32 
		wanted = 4116459851	
		assertCorrectValue(t, hash, wanted)
	})

	t.Run("returns hash for integer key", func(t *testing.T){
		hash := FNVHash(123)
		var wanted uint32 
		wanted = 1916298011	
		assertCorrectValue(t, hash, wanted)
	})

	t.Run("returns hash for boolean key", func(t *testing.T){
		hash := FNVHash(false)
		var wanted uint32 
		wanted = 84696351	
		assertCorrectValue(t, hash, wanted)
	})
}

func assertCorrectValue[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}