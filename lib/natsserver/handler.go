package natsserver

import (
	"encoding/json"
	"reflect"

	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

func NatsHandler(req api.Request, rplFunc func(api.Request) (api.Reply, *errs.Error)) nats.MsgHandler {
	return func(msg *nats.Msg) {

		defer func() {
			if p := recover(); p != nil {
				msg.Respond([]byte(`{"Error": {"Code": "ERROR_INTERNAL_SYSTEM", "Msg": "panic recovered"}}`))
			}
		}()

		tmp := reflect.TypeOf(req)
		if tmp.Kind() != reflect.Ptr || tmp.Elem().Kind() != reflect.Struct || !tmp.Implements(reflect.TypeOf((*api.Request)(nil)).Elem()) {
			msg.Respond([]byte(`{"Error": {"Code": "ERROR_INTERNAL_SYSTEM", "Msg": "unable to register Request argument. Must be pointer to struct which implements Request interface."} }`))
			return
		}
		tmp2 := reflect.New(tmp.Elem()).Interface().(api.Request)

		if err := json.Unmarshal(msg.Data, &tmp2); err != nil {
			msg.Respond([]byte(`{"Error": {"Code": "ERROR_INTERNAL_SYSTEM", "Msg": "unable to unmarshal request data"} }`))
			return
		}
		rpl, errApi := rplFunc(tmp2)
		if errApi != nil {
			errApiJs, err := json.Marshal(errApi)
			if err != nil {
				msg.Respond([]byte(`{"Error": {"Code": "ERROR_INTERNAL_SYSTEM", "Msg": "unable to unmarshal request error data"} }`))
				return
			}
			msg.Respond(errApiJs)
			return
		}

		rplJs, err := json.Marshal(rpl)
		if err != nil {
			msg.Respond([]byte(`{"Error": {"Code": "ERROR_INTERNAL_SYSTEM", "Msg": "unable to unmarshal response data" }}`))
			return
		}

		_ = msg.Respond(rplJs)
	}
}
