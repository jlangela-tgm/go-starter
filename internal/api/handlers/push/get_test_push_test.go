package push_test

import (
	"net/http"
	"testing"

	"github.com/jlangela-tgm/go-starter/internal/api"
	"github.com/jlangela-tgm/go-starter/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTestPush(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/push/test", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}

func TestGetTestPushUnauthorized(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/api/v1/push/test", nil, nil)

		assert.Equal(t, http.StatusUnauthorized, res.Result().StatusCode)
	})
}
