package api

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
	"github.com/nats-io/nats.go"
)

// WorkerNATS subscribes to NATS queue
func WorkerNATS(set *setup.Setup, queue string) *errs.Error {
	l := set.Log()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var worker *nats.Subscription
	var err error

	i := 0 // delete

	worker, err = set.NATS().QueueSubscribe(queue, queue, func(m *nats.Msg) {

		/*go func() {
			l.Info("got message from nats")
		}()*/

		i++
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		l.Info("[#%d] Received on [%s] Queue[%s] Pid[%d]: '%s'", i, m.Subject, m.Sub.Queue, os.Getpid(), string(m.Data))
		m.Respond([]byte("some reply string"))

	})

	if err != nil || worker == nil {
		l.Error("unable to subscribe to queue")
		signalChan <- syscall.SIGTERM
	}

	set.NATS().Flush()

	if err := set.NATS().LastError(); err != nil {
		l.Fatal("LastError: %s", err)
		signalChan <- syscall.SIGTERM
	}

	l.Info("Listening to " + queue + " queue")

	<-signalChan

	l.Info("Draining...")
	set.NATS().Drain()

	l.Info("Exiting")

	/*defer func(worker *nats.Subscription) {
		err := worker.Unsubscribe()
		if err != nil {
			l.Error("unable to unsubscribe queue:", err)
			signalChan <- syscall.SIGTERM
		}
	}(worker)

	l.Info("Listening to " + queue + " queue")

	<-signalChan

	l.Info("Exiting ...")*/

	return nil
}
