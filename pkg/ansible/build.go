package ansible

import (
	"fmt"
	"strings"
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
	OtherPipDependencies []string `yaml:"otherPipDependencies,omitempty"`
	RequirementsFiles []string `yaml:"requirementsFiles,omitempty"`
	ConstraintsFiles []string `yaml:"constraintsFiles,omitempty"`
}

func parseConfig(m *Mixin, input *BuildInput) {
	if suppliedClientVersion := input.Config.ClientVersion; suppliedClientVersion != "" {
		m.ClientVersion = suppliedClientVersion
	}

	if otherPipDependencies := input.Config.OtherPipDependencies; len(otherPipDependencies) > 0 {
		var otherPipDependenciesWithQuotes []string
		for _, x := range otherPipDependencies {
			// Ensure each string to have the format 'str'
			// Without these quotes, we could fall into issues by having lines like pip install dep<2 and being misinterpreted in shell 
			otherPipDependenciesWithQuotes = append(otherPipDependenciesWithQuotes, "'" + x + "'") 
		}
		m.OtherPipDependencies = otherPipDependenciesWithQuotes
	}

	if requirementsFiles := input.Config.RequirementsFiles; len(requirementsFiles) > 0 {
		var requirementsFilesWithPrefix []string
		for _, x := range requirementsFiles {
			// Ensure each string to have the format <--requirement 'str'>
			requirementsFilesWithPrefix = append(requirementsFilesWithPrefix, "--requirement '" + x + "'") 
		}
		m.RequirementsFiles = requirementsFilesWithPrefix
	}

	if constraintsFiles := input.Config.ConstraintsFiles; len(constraintsFiles) > 0 {
		var constraintsFilesWithPrefix []string
		for _, x := range constraintsFiles {
			// Ensure each string to have the format <--constraint 'str'>
			constraintsFilesWithPrefix = append(constraintsFilesWithPrefix, "--constraint '" + x + "'") 
		}
		m.ConstraintsFiles = constraintsFilesWithPrefix
	}
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

	// Parse all configs received from users
	parseConfig(m, &input)

	// Example of pulling and defining a client version for your mixin
	fmt.Fprintf(m.Out, `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible%s' %s %s %s
	`, 
	m.ClientVersion, 
	strings.Join(m.OtherPipDependencies[:], " "), 
	strings.Join(m.RequirementsFiles[:], " "), 
	strings.Join(m.ConstraintsFiles[:], " "))

	return nil
}
