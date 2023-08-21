package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func newTestServer(url string) *Server {
	srv := &Server{
		echo: echo.New(),
	}

	srv.configureMiddleware()
	srv.configureRoutes()

	return srv
}

func Test_Health(t *testing.T) {
	srv := newTestServer("")

	req, _ := http.NewRequest(echo.GET, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Test_Health expected code %v, got %v", http.StatusOK, rec.Code)
	}
}

func Test_Root(t *testing.T) {
	srv := newTestServer("")

	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Test_Root expected code %v, got %v", http.StatusOK, rec.Code)
	}
}

func Test_StartShutdown(t *testing.T) {
	srv := newTestServer("")

	go func() {
		_ = srv.Start(":3928")
	}()
	assert.Nil(t, srv.Shutdown(context.Background()))
}
