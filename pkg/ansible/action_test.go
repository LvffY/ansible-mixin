package ansible

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
	"fmt"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	testcases := []struct {
		name            string        // Test case name
		file            string        // Path to th test input yaml
		expectedActionName string // Expected action name
		expectedDescription string        // Description that you expect to be found
		expectedArguments   []string      // Arguments that you expect to be found
		expectedFlags       builder.Flags // Flags that you expect to be found
		expectedCommand  string      // Command that you expect to be found
		expectedSuffixArgs  []string      // Suffix arguments that you expect to be found
		expectedWorkingDir  string      // WorkingDir that you expect to be found
	}{
		{
			name: "AdHoc install no outputs", 
			file: "testdata/action_test/adhoc/no_outputs.yaml", 
			expectedActionName: "install",
			expectedDescription: "Run our ansible adhoc command" , 
			expectedArguments: []string{"localhost"}, 
			expectedFlags: builder.Flags{builder.NewFlag("module-name", "debug"), builder.NewFlag("args", "msg='Hello from ansible AdHoc command !'")}, 
			expectedCommand: "ansible",
			expectedSuffixArgs: []string{},
			expectedWorkingDir: "",
		},
		{
			name: "AdHoc install with outputs", 
			file: "testdata/action_test/adhoc/with_outputs.yaml", 
			expectedActionName: "install",
			expectedDescription: "Run our ansible adhoc command and capture output" , 
			expectedArguments: []string{"host_pattern*"}, 
			expectedFlags: builder.Flags{builder.NewFlag("module-name", "debug"), builder.NewFlag("args", "var=variable_name"), builder.NewFlag("inventory", "/etc/ansible/hosts")}, 
			expectedCommand: "ansible",
			expectedSuffixArgs: []string{},
			expectedWorkingDir: "/working/dir",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			var action Action
			err = yaml.Unmarshal(b, &action)
			require.NoError(t, err)

			assert.Equal(t, 
				tc.expectedActionName, 
				action.Name, 
				fmt.Sprintf("Bad action name. Received %s while expected %s", action.Name, tc.expectedActionName),
			)
			require.Len(t, action.Steps, 1)
			
			step := action.Steps[0]

			assert.Equal(t, 
				tc.expectedDescription, 
				step.Description, 
				fmt.Sprintf("Bad action description. Received %s while expected %s", step.Description, tc.expectedDescription),
			)
			
			args := step.GetArguments()
			assert.Equal(t, 
				tc.expectedArguments, 
				args, 
				fmt.Sprintf("Bad arguments. Received %s while expected %s", args, tc.expectedArguments),
			)

			flags := step.GetFlags()
			sort.Sort(flags)
			sort.Sort(tc.expectedFlags)
			assert.Equal(t, 
				tc.expectedFlags, 
				flags,
				fmt.Sprintf("Bad flags. Received %s while expected %s", flags, tc.expectedFlags),
			)

			command := step.GetCommand()
			assert.Equal(t, 
				tc.expectedCommand, 
				command, 
				fmt.Sprintf("Bad command. Received %s while expected %s", command, tc.expectedCommand),
			)


			suffixArgs := step.GetSuffixArguments()
			assert.Equal(t, 
				tc.expectedSuffixArgs, 
				suffixArgs, 
				fmt.Sprintf("Bad suffix args. Received %s while expected %s", suffixArgs, tc.expectedSuffixArgs),
			)

			workingDir := step.GetWorkingDir()
			assert.Equal(t, 
				tc.expectedWorkingDir, 
				workingDir, 
				fmt.Sprintf("Bad working dir. Received %s while expected %s", workingDir, tc.expectedWorkingDir),
			)
		})
	}
}
