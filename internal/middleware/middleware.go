package middleware

import "net/http"

// Middleware defines a function type that wraps an http.Handler with additional functionality.
type Middleware func(http.Handler) http.Handler

// CreateStack takes a variadic list of Middleware functions and returns a single Middleware that
// applies them in reverse order. This allows for composing multiple middleware layers, where the
// last provided middleware is executed first, and the first provided middleware is executed last.
// The function iterates over the slice of Middleware from the end, wrapping the 'next' http.Handler
// with each Middleware in turn, creating a stacked chain of middleware.
func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
