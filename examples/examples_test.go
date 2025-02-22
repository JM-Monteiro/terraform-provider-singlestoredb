package examples_test

import (
	"testing"

	"github.com/singlestore-labs/terraform-provider-singlestoredb/examples"
	"github.com/stretchr/testify/require"
)

func TestEmbedsExamples(t *testing.T) {
	require.NotEmpty(t, examples.Provider)
	require.NotEmpty(t, examples.Regions)
	require.NotEmpty(t, examples.WorkspaceGroupsListDataSource)
	require.NotEmpty(t, examples.WorkspaceGroupsGetDataSource)
	require.NotEmpty(t, examples.WorkspaceGroupsResource)
	require.NotEmpty(t, examples.WorkspacesListDataSource)
	require.NotEmpty(t, examples.WorkspacesGetDataSource)
	require.NotEmpty(t, examples.WorkspacesResource)
}
