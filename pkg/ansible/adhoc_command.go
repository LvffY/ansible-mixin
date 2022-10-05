package ansible

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ AnsibleInstruction = AdhocCommand{}

type AdhocCommand struct {
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
}

func (c AdhocCommand) GetSuffixArguments() []string {
	return c.SuffixArguments
}

func (c AdhocCommand) GetCommand() string {
	return "ansible"
}

func (c AdhocCommand) GetArguments() []string {
	// Final Command: ansible ARGUMENTS --FLAGS

	args := []string{}
	args = append(args, c.Arguments...)

	return args
}

func (c AdhocCommand) GetFlags() builder.Flags {
	var flags builder.Flags
	flags = append(flags, c.Flags...)
	return flags
}

func (c AdhocCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c AdhocCommand) GetWorkingDir() string {
	return c.WorkingDir
}
