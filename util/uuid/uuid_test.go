package uuid

import (
	"fmt"
	"testing"
)

func TestUniqueFilename(t *testing.T) {
	str := UniqueFilename(10)
	fmt.Println(str)
}
