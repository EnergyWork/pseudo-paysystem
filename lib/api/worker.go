package api

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

// WorkerNATS subscribes to NATS queue
func WorkerNATS(set *setup.Setup, queue string) *errs.Error {
	l := set.Log()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var worker *nats.Subscription
	var err error

	worker, err = set.NATS().Subscribe(queue, func(m *nats.Msg) {

		go func() {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			l.Info("Received on [%s] Queue[%s] Pid[%d]: '%s'", m.Subject, m.Sub.Queue, os.Getpid(), string(m.Data))
			m.Respond([]byte("some reply string 1"))
		}()

	})

	if err != nil || worker == nil {
		l.Error("unable to subscribe to queue")
		signalChan <- syscall.SIGTERM
	}

	_ = set.NATS().Flush()

	if err = set.NATS().LastError(); err != nil {
		l.Fatal("LastError: %s", err)
		signalChan <- syscall.SIGTERM
	}

	l.Info("Listening to " + queue + " queue")

	<-signalChan

	l.Info("Draining...")
	err = set.NATS().Drain()
	if err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}

	l.Info("Exiting...")
	return nil
}
