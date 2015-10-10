package auth

import (
	"crypto/subtle"

	"github.com/revel/revel"
)

func ServiceEnabled(service string, v *revel.Validation) bool {
	v.Match(service, enabledServices)

	if v.HasErrors() {
		revel.INFO.Printf("Got not supported service name (%v)", service)
		return false
	}

	return true
}

func StateValid(service, sessionId, state string, v *revel.Validation) bool {
	v.Required(state)
	v.Length(state, 32)

	if v.HasErrors() {
		revel.INFO.Printf("Got invalid state (%v) value", state)
		return false
	}

	cachedState, ok := RetriveState(service, sessionId)
	if !ok {
		revel.INFO.Printf("Cached state not found for service (%v) and session Id (%v)",
			service, sessionId)
		return false
	}

	if eq := subtle.ConstantTimeCompare([]byte(state), []byte(cachedState)); eq != 1 {
		revel.INFO.Printf("State (%s) from request and cached state (%s) not equal",
			state, cachedState)
		return false
	}

	return true
}

func CodeValid(service, code string, v *revel.Validation) bool {
	v.Required(code)
	v.Length(code, 20)

	if v.HasErrors() {
		revel.INFO.Printf("Got invalid code (%s) value", code)
		return false
	}

	return true
}
