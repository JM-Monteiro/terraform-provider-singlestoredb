package provider_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/singlestore-labs/terraform-provider-singlestoredb/examples"
	"github.com/singlestore-labs/terraform-provider-singlestoredb/internal/provider/config"
	"github.com/singlestore-labs/terraform-provider-singlestoredb/internal/provider/testutil"
	"github.com/stretchr/testify/require"
)

func TestProviderAuthenticates(t *testing.T) {
	apiKey := "buzz"
	actualAPIKey := ""

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualAPIKey = r.Header.Get("Authorization")
	}))
	defer server.Close()

	testutil.UnitTest(t, testutil.Config{
		APIServiceURL: server.URL,
		APIKey:        apiKey,
	}, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: examples.Regions,
			},
		},
	})

	require.Equal(t, fmt.Sprintf("Bearer %s", apiKey), actualAPIKey)
}

func TestProviderAuthenticationError(t *testing.T) {
	apiKey := "foo"
	actualAPIKey := ""

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualAPIKey = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	r := regexp.MustCompile(http.StatusText(http.StatusUnauthorized))

	testutil.UnitTest(t, testutil.Config{
		APIServiceURL: server.URL,
		APIKey:        apiKey,
	}, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:      examples.Regions,
				ExpectError: r,
			},
		},
	})

	require.Equal(t, fmt.Sprintf("Bearer %s", apiKey), actualAPIKey)
}

func TestProviderAuthenticatesFromEnv(t *testing.T) {
	apiKey := "bar"
	actualAPIKey := ""

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualAPIKey = r.Header.Get("Authorization")
	}))
	defer server.Close()

	testutil.UnitTest(t, testutil.Config{
		APIServiceURL: server.URL,
		APIKeyFromEnv: apiKey,
	}, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: examples.Regions,
			},
		},
	})

	require.Equal(t, fmt.Sprintf("Bearer %s", apiKey), actualAPIKey)
}

func TestProviderAuthenticationErrorIntegration(t *testing.T) {
	apiKey := "foo"

	r := regexp.MustCompile(http.StatusText(http.StatusUnauthorized))

	testutil.IntegrationTest(t, apiKey, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:      examples.Regions,
				ExpectError: r,
			},
		},
	})
}

func TestProviderAuthenticatesIntegration(t *testing.T) {
	apiKey := os.Getenv(config.EnvTestAPIKey)

	testutil.IntegrationTest(t, apiKey, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: examples.Regions,
			},
		},
	})
}
