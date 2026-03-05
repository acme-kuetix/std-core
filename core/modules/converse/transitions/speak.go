package transitions

import (
	"github.com/kuetix/engine/pkg/domain"
	"github.com/kuetix/engine/pkg/domain/interfaces"
	"github.com/kuetix/engine/pkg/workflow"
)

type speakTransitions struct {
	workflow.BaseServiceTransition
}

func NewSpeakTransitions() interfaces.ServiceTransitions { return &speakTransitions{} }

// Say returns a classic greeting. It does not depend on any input.
func (t *speakTransitions) Say(on string, v ...any) (r domain.FlowStepResult) {
	var value interface{} = v
	if len(v) == 1 {
		value = v[0]
	}
	msg := map[string]interface{}{on: value}
	// logger.Info("["+on+"] ", msg[on])

	r.Success = true
	r.Response = msg
	return
}
