package test

import (
	"sync"
	"testing"
	"time"

	balance "github.com/energywork/pseudo-paysystem/balance/api"
)

func TestRequestNATS(t *testing.T) {
	set := GetSetup(t)
	data := []byte("some request data")
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			msg, err := set.NATS().Request(balance.Queue, data, time.Second*5)
			if err != nil {
				if set.NATS().LastError() != nil {
					t.Errorf("%v for request", set.NATS().LastError())
				}
				t.Errorf("%v for request", err)
			} else {
				t.Logf("Published [%s] : '%s'", balance.Queue, data)
				t.Logf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
