package ansible

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMixin_Build(t *testing.T) {
	m := NewTestMixin(t)

	err  := m.Build()
	require.NoError(t, err, "Build failed")

	gotOutput := m.TestContext.GetOutput()

	wantOutput := `RUN pip install --upgrade --no-cache-dir pip && \
	pip install --upgrade --no-cache-dir setuptools wheel && \
	pip install --upgrade --no-cache-dir 'ansible'   
	`

	assert.Equal(t, wantOutput, gotOutput)
}