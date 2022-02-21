package ansible

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMixinBuildWithDefaults(t *testing.T) {
	m := NewTestMixin(t)

	err := m.Build()
	require.NoError(t, err, "Build failed")

	gotOutput := m.TestContext.GetOutput()

	wantOutput := `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible'   
	`

	assert.Equal(t, wantOutput, gotOutput)
}

// TODO: Add unit tests with some non defaults values
/* func TestMixinBuildWithoutDefault(t *testing.T) {
	m := NewTestMixin(t)

	m.ClientVersion = "<2.10"
	m.RequirementsFiles = []string{"requirements.txt", "http://host.com/requirements.txt"}
	m.ConstraintsFiles = []string{"constraints.txt", "http://host.com/constraints.txt"}
	m.OtherPipDependencies = []string{"aPythonDep", "anotherPythonDep<specificVersion"}

	err := m.Build()
	require.NoError(t, err, "Build failed")

	gotOutput := m.TestContext.GetOutput()

	wantOutput := `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible<2.10' 'aPythonDep' 'anotherPythonDep<specificVersion' --requirement 'requirements.txt' --requirement 'http://host.com/requirements.txt' --constraint 'constraints.txt' --constraint 'http://host.com/constraints.txt'
	`

	assert.Equal(t, wantOutput, gotOutput)
} */
