package utils

import "testing"

func TestUtils_FNVHash(t *testing.T){
	t.Run("returns hash for key", func(t *testing.T){
		hash := FNVHash("Hello")
		var wanted uint32 
		wanted = 4116459851	
		assertCorrectValue(t, hash, wanted)
	})
}

func assertCorrectValue[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}