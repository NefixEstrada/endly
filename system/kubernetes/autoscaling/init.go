package autoscaling

import (
	"github.com/viant/endly"
	_ "github.com/viant/endly/system/kubernetes/autoscaling/v1"
)

func init() {
	_ = endly.Registry.Register(func() endly.Service {
		return New()
	})
}
