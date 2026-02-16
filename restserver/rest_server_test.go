package restserver

import (
	"testing"
)

func TestRegisterController(t *testing.T) {
	// Reset global state
	ServerRoutes = nil

	controller := &mockController{
		routes: []Route{
			{Path: "/api/users", HttpMethod: Get},
			{Path: "/api/users", HttpMethod: Post},
		},
	}

	RegisterController(controller)

	if len(ServerRoutes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(ServerRoutes))
	}
	if ServerRoutes[0].Path != "/api/users" {
		t.Fatalf("expected /api/users, got %s", ServerRoutes[0].Path)
	}
	if ServerRoutes[0].HttpMethod != Get {
		t.Fatalf("expected GET, got %s", ServerRoutes[0].HttpMethod)
	}
	if ServerRoutes[1].HttpMethod != Post {
		t.Fatalf("expected POST, got %s", ServerRoutes[1].HttpMethod)
	}
}

func TestRegisterController_Multiple(t *testing.T) {
	ServerRoutes = nil

	c1 := &mockController{routes: []Route{{Path: "/a", HttpMethod: Get}}}
	c2 := &mockController{routes: []Route{{Path: "/b", HttpMethod: Post}}}

	RegisterController(c1)
	RegisterController(c2)

	if len(ServerRoutes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(ServerRoutes))
	}
}

func TestRegisterMiddleware(t *testing.T) {
	ServerMiddlewares = nil

	m := &mockMiddleware{}
	RegisterMiddleware(m)

	if len(ServerMiddlewares) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(ServerMiddlewares))
	}
}

func TestRegisterMiddleware_Multiple(t *testing.T) {
	ServerMiddlewares = nil

	RegisterMiddleware(&mockMiddleware{})
	RegisterMiddleware(&mockMiddleware{})
	RegisterMiddleware(&mockMiddleware{})

	if len(ServerMiddlewares) != 3 {
		t.Fatalf("expected 3 middlewares, got %d", len(ServerMiddlewares))
	}
}

type mockController struct {
	routes []Route
}

func (m *mockController) Routes() []Route {
	return m.routes
}

type mockMiddleware struct{}

func (m *mockMiddleware) Apply(ctx WebContext) error {
	return nil
}
