package test

import (
	"sync"
	"testing"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

const (
	Queue = "balance-queue"
)

func TestRequestNATS(t *testing.T) {
	cfg := config.New("setup_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	set := setup.GetSetupForTest(t, cfg)
	// -----------------------------------------------------------------------------------------------------------------
	data := []byte("some request data")
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			msg, err := set.NATS().Request(Queue, data, time.Second*10)
			if err != nil {
				if set.NATS().LastError() != nil {
					t.Errorf("0. %v for request", set.NATS().LastError())
				}
				t.Errorf("1. %v for request", err)
			} else {
				t.Logf("Published [%s] : '%s'", Queue, data)
				t.Logf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSingleRequest(t *testing.T) {
	cfg := config.New("setup_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	set := setup.GetSetupForTest(t, cfg, false)
	// -----------------------------------------------------------------------------------------------------------------
	data := []byte("some request data single")
	subj := "balance/wallet/create"
	msg, err := set.NATS().Request(subj, data, time.Second*10)
	if err != nil {
		if set.NATS().LastError() != nil {
			t.Errorf("%v", set.NATS().LastError())
		} else {
			t.Errorf("%v", err)
		}
	} else {
		t.Logf("Published [%s] : '%s'", subj, data)
		t.Logf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
	}
}

func TestSubReq(t *testing.T) {
	cfg := config.New("setup_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	set := setup.GetSetupForTest(t, cfg, false)
	// -----------------------------------------------------------------------------------------------------------------

	h := func() nats.MsgHandler {
		return func(msg *nats.Msg) {
			t.Logf("Subject: %s", msg.Subject)
			t.Logf("получили сообщение: %s", msg.Data)
			_ = msg.Respond([]byte(`{"test": "result"}`))
		}
	}
	q, err := set.NATS().QueueSubscribe("test-subj", "test-queue", h())

	if q == nil || err != nil {
		t.Fatal("error")
	}

	time.Sleep(time.Millisecond * 500)

	defer func() {
		if err := q.Drain(); err != nil {
			t.Fatal(err)
		}
	}()

	d := []byte(`{"kek": 101"}`)
	rpl, err := set.NATS().Request("test-subj", d, time.Second*5)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("rpl: %s", rpl.Data)
}

func TestNewNATSRequest(t *testing.T) {
	cfg := config.New("setup_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	set := setup.GetSetupForTest(t, cfg, false)
	// -----------------------------------------------------------------------------------------------------------------
	data := []byte("some request data single")
	subj := "balance/wallet/create"
	rpl, err := api.NewNATSRequest(set, subj, data, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rpl: %s", rpl)
}
