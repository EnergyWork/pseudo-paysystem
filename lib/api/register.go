package api

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/energywork/pseudo-paysystem/lib/errs"
)

var registeredTypes map[string]Registered
var registeredNames map[string]string

func init() {
	registeredTypes = make(map[string]Registered, 0)
	registeredNames = make(map[string]string, 0)
}

type Registered struct {
	Name   string
	Type   reflect.Type
	Queue  string
	Exec   string
	Secure bool
	Auth   string
}

func (r *Registered) Interface() interface{} {
	return reflect.New(r.Type.Elem()).Interface()
}

// Register ...
func Register(
	queue string,
	name string,
	request interface{},
	execute interface{},
	auth ...interface{},
) {
	if _, exists := registeredTypes[name]; exists {
		return
	}
	execMethod := getMethodName(execute)
	if len(execMethod) == 0 {
		return
	}

	var secure bool
	var authMethod string
	if len(auth) > 0 {
		secure = true
		authMethod = getMethodName(auth[0])
	}

	t := reflect.TypeOf(request)

	registeredTypes[name] = Registered{
		Name:   name,
		Type:   t,
		Queue:  queue,
		Exec:   execMethod,
		Secure: secure,
		Auth:   authMethod,
	}
	registeredNames[t.String()] = name
}

func getMethodName(method interface{}) string {
	if reflect.TypeOf(method).Kind() != reflect.Func {
		return ""
	}
	pc := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
	s := strings.LastIndex(pc, ".") + 1
	e := strings.LastIndex(pc, "-fm")
	if e < s {
		return pc[s:]
	}
	return pc[s:e]
}

func GetRegistered(name string) (*Registered, *errs.Error) {
	reg, ok := registeredTypes[name]
	if !ok {
		return nil, errs.New().SetCode(errs.RequestUnknown).SetMsg("Unknown request")
	}

	return &reg, nil
}

func GetRequestName(r interface{}) string {
	return registeredNames[reflect.TypeOf(r).String()]
}
