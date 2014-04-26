package natto

import (
	"./terst"
	"testing"
	"time"
)

var is = terst.Is

func Test(t *testing.T) {
	terst.Terst(t, func() {
		start := time.Now()
		Run(`
            setTimeout(function(){}, 1000);
        `)
		duration := time.Since(start)
		is(duration, ">=", 1*time.Second)
	})
}
