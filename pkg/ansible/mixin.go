//go:generate packr2
package ansible

import (
	"get.porter.sh/porter/pkg/context"
)

const defaultClientVersion string = ""

type Mixin struct {
	*context.Context
	ClientVersion string
	//add whatever other context/state is needed here
}

// New azure mixin client, initialized with useful defaults.
func New() (*Mixin, error) {
	return &Mixin{
		Context:       context.New(),
		ClientVersion: defaultClientVersion,
	}, nil

}
