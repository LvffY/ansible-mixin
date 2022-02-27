package ansible

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ AnsibleInstruction = PlaybookCommand{}

type PlaybookCommand struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	WorkingDir  string   `yaml:"dir,omitempty"`
	Arguments   []string `yaml:"arguments,omitempty"`

	// Useful when the CLI you are calling wants some arguments to come after flags
	// Arguments are passed first, then Flags, then SuffixArguments.
	SuffixArguments []string `yaml:"suffix-arguments,omitempty"`

	Flags          builder.Flags `yaml:"flags,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`

	// Allow the user to ignore some errors
	// Adds the ignoreError functionality from the exec mixin
	// https://release-v1.porter.sh/mixins/exec/#ignore-error
	builder.IgnoreErrorHandler `yaml:"ignoreError,omitempty"`
}

func (c PlaybookCommand) GetSuffixArguments() []string {
	return c.SuffixArguments
}

func (c PlaybookCommand) GetCommand() string {
	return "ansible-playbook"
}

func (c PlaybookCommand) GetArguments() []string {
	// Final Command: ansible ARGUMENTS --FLAGS

	args := []string{}
	args = append(args, c.Arguments...)

	return args
}

func (c PlaybookCommand) GetFlags() builder.Flags {
	var flags builder.Flags
	flags = append(flags, c.Flags...)
	return flags
}

func (c PlaybookCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c PlaybookCommand) GetWorkingDir() string {
	return c.WorkingDir
}

func (c PlaybookCommand) GetOutputs() []Output {
	return c.Outputs
}