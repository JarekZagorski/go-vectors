package vec

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
    var tests = []struct {
        a, b Vi2d
        ans float64
    }{
        {Vi2d{4, 0}, Vi2d{0, 3}, 5.0},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s -> %s", tt.a.String(), tt.b.String())
        t.Run(testname, func(t *testing.T){
            ans := Distance(tt.a, tt.b)

            if ans != tt.ans {
                t.Errorf("expected %g, got %g", tt.ans, ans)
            }
        })
    }
}
