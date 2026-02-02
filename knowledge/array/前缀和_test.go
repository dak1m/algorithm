package array

import (
	"fmt"
	"testing"
)

func TestScoreRangeQuery(t *testing.T) {
	fmt.Println(scoreRangeQuery([]int{0, 1, 2, 2, 3, 4, 5}, 1, 2))
}
