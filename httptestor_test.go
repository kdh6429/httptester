package httptester

import (
	"testing"
	"time"
)

func TestCtxTimeout(t *testing.T) {
	ntest := 3
	nfactor := 2

	tester, rchan := NewFakeHttpTest().
		Timeout(time.Second * 2).
		NClient(ntest).
		NFactor(nfactor).
		Build()

	tester.Do()

LOOP:
	for {
		select {
		case _, ok := <-rchan:
			if !ok {
				break LOOP
			}
		}
	}
	// shoudl be reach here around timeout(2 sec) time
}
