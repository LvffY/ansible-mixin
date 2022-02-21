package ansible

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the ansible mixin in porter.yaml
// mixins:
// - ansible:
//	  clientVersion: "v0.0.0"

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
	OtherPipDependencies string `yaml:"otherPipDependencies,omitempty"`
}

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	// Create new Builder.
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "" {
		m.ClientVersion = suppliedClientVersion
	}

	// Example of pulling and defining a client version for your mixin
	fmt.Fprintf(m.Out, `RUN pip install --upgrade --no-cache-dir pip && \
    pip install --upgrade --no-cache-dir setuptools wheel && \
    pip install --upgrade --no-cache-dir 'ansible%s'
    `, m.ClientVersion)

	return nil
}
