package ansible

import (
	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var _ builder.ExecutableAction = Action{}
var _ builder.BuildableAction = Action{}

type Action struct {
	Name  string
	Steps []AnsibleStep // using UnmarshalYAML so that we don't need a custom type per action
}

// MarshalYAML converts the action back to a YAML representation
// install:
//   ansible:
//     ...
func (a Action) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{a.Name: a.Steps}, nil
}

// MakeSteps builds a slice of Step for data to be unmarshaled into.
func (a Action) MakeSteps() interface{} {
	return &[]AnsibleStep{}
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - ansible: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, a)
	if err != nil {
		return err
	}

	for actionName, action := range results {
		a.Name = actionName
		for _, result := range action {
			step := result.(*[]AnsibleStep)
			a.Steps = append(a.Steps, *step...)
		}
		break // There is only 1 action
	}
	return nil
}

func (a Action) GetSteps() []builder.ExecutableStep {
	// Go doesn't have generics, nothing to see here...
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}

type Step struct {
	Instruction `yaml:"ansible"`
}

type AnsibleStep struct {
	Description string
	AnsibleInstruction
	Output
}

// Actions is a set of actions, and the steps, passed from Porter.
type Actions []Action

// UnmarshalYAML takes chunks of a porter.yaml file associated with this mixin
// and populates it on the current action set.
// install:
//   ansible:
//     ...
//   ansible:
//     ...
// upgrade:
//   ansible:
//     ...
func (a *AnsibleStep) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Turn the yaml into a raw map so we can iterate over the values and
	// look for which command was used
	stepMap := map[string]map[string]interface{}{}
	err := unmarshal(&stepMap)
	if err != nil {
		return errors.Wrap(err, "Could not unmarshal yaml into a raw ansible command")
	}

	// get at the values defined under "ansible"
	ansibleStep := stepMap["ansible"]

	for actionName, action := range ansibleStep {
		var cmd AnsibleInstruction
		switch actionName {
		case "description":
			a.Description = action.(string)
			continue
		case "adhoc":
			cmd = &AdhocCommand{}
		default:
			return errors.Errorf("Unsupported ansible mixin command %s", actionName)
		}

		b, err := yaml.Marshal(action)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(b, cmd)
		if err != nil {
			return err
		}

		a.AnsibleInstruction = cmd
	}
	return nil
}

var _ builder.HasOrderedArguments = Instruction{}
var _ builder.ExecutableStep = Instruction{}
var _ builder.StepWithOutputs = Instruction{}

type Instruction struct {
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

func (s Instruction) GetCommand() string {
	return "ansible"
}

func (s Instruction) GetWorkingDir() string {
	return s.WorkingDir
}

func (s Instruction) GetArguments() []string {
	return s.Arguments
}

func (s Instruction) GetSuffixArguments() []string {
	return s.SuffixArguments
}

func (s Instruction) GetFlags() builder.Flags {
	return s.Flags
}

func (s Instruction) SuppressesOutput() bool {
	return s.SuppressOutput
}

func (s Instruction) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}

type AnsibleInstruction interface {
	builder.ExecutableStep
	builder.HasOrderedArguments
	builder.SuppressesOutput
}

var _ AnsibleInstruction = Instruction{}

var _ builder.OutputJsonPath = Output{}
var _ builder.OutputFile = Output{}
var _ builder.OutputRegex = Output{}

type Output struct {
	Name string `yaml:"name"`

	// See https://porter.sh/mixins/exec/#outputs
	// TODO: If your mixin doesn't support these output types, you can remove these and the interface assertions above, and from #/definitions/outputs in schema.json
	JsonPath string `yaml:"jsonPath,omitempty"`
	FilePath string `yaml:"path,omitempty"`
	Regex    string `yaml:"regex,omitempty"`
}

func (o Output) GetName() string {
	return o.Name
}

func (o Output) GetJsonPath() string {
	return o.JsonPath
}

func (o Output) GetFilePath() string {
	return o.FilePath
}

func (o Output) GetRegex() string {
	return o.Regex
}
