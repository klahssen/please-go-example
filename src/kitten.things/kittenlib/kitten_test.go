package kittenlib

import (
	"testing"
  "fmt"
  // "time"

	"github.com/stretchr/testify/assert"

	pb "proto/kitten"
)

func TestProvideKitten(t *testing.T) {
	kitten := ProvideKitten()
  // time.Sleep(2000 * time.Millisecond)
	assert.Equal(t, pb.Breed_RUSSIAN_BLUE, kitten.Breed)
	assert.Equal(t, "Mighty Paws", kitten.Name)
  fmt.Println(kitten.Age)
  if kitten.Age != 6 {
    t.Errorf("expected %s, was %s", 6, kitten.Age)
  }
}
