package ansible

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"io/ioutil"
)

func TestMixinBuildWithDefaults(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/build_test/build-input-without-config.yaml")
	require.NoError(t, err)

	m := NewTestMixin(t)
	m.Debug = false
	m.In = bytes.NewReader(b)

	err = m.Build()
	require.NoError(t, err, "build failed")

	wantOutput := `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible'   
	`

	gotOutput := m.TestContext.GetOutput()
	assert.Equal(t, wantOutput, gotOutput)
}

func TestMixinBuildWithoutDefault(t *testing.T) {

	b, err := ioutil.ReadFile("testdata/build_test/build-input-with-config.yaml")
	require.NoError(t, err)

	m := NewTestMixin(t)
	m.Debug = false
	m.In = bytes.NewReader(b)

	err = m.Build()
	require.NoError(t, err, "build failed")

	wantOutput := `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible==2.10' 'jmespath' --requirement 'requirements.txt' --constraint 'constraints.txt'
	`

	gotOutput := m.TestContext.GetOutput()
	assert.Equal(t, wantOutput, gotOutput)
}
