package gofalcon

import (
    "encoding/json"
    "net/http"
)

// Context holds the HTTP request and response
type Context struct {
    Writer  http.ResponseWriter
    Request *http.Request
}

// NewContext creates a new Context
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
    return &Context{
        Writer:  w,
        Request: req,
    }
}

// JSON writes a JSON response
func (c *Context) JSON(status int, obj interface{}) {
    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.WriteHeader(status)
    json.NewEncoder(c.Writer).Encode(obj)
}

// M is a shortcut for map[string]interface{}
type M map[string]interface{}

// HandlerFunc defines the request handler used by GoFalcon
type HandlerFunc func(*Context)

// Router stores the routes and their handlers
type Router struct {
    handlers map[string]HandlerFunc
}

// NewRouter creates a new Router
func NewRouter() *Router {
    return &Router{handlers: make(map[string]HandlerFunc)}
}

// Handle registers a new route with a handler
func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
    key := method + "-" + pattern
    r.handlers[key] = handler
}

// ServeHTTP handles the HTTP request
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    key := req.Method + "-" + req.URL.Path
    if handler, ok := r.handlers[key]; ok {
        c := NewContext(w, req)
        handler(c)
    } else {
        http.NotFound(w, req)
    }
}

// NewServer creates a new Router (alias for NewRouter)
func NewServer() *Router {
    return NewRouter()
}

// Run starts the HTTP server
func (r *Router) Run(addr ...string) error {
    address := ":8080"
    if len(addr) > 0 {
        address = addr[0]
    }
    return http.ListenAndServe(address, r)
}
