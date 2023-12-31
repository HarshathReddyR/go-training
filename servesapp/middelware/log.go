package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type reqKey int

const TraceIdKey reqKey = 999

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Generate a new unique identifier (UUID)
		traceId := uuid.NewString()

		// Fetch the current context from the gin context
		ctx := c.Request.Context()

		// Add the trace id in context so it can be used by upcoming processes in this request's lifecycle
		ctx = context.WithValue(ctx, TraceIdKey, traceId)

		// The 'WithContext' method on 'c.Request' creates a new copy of the request ('req'),
		// but with an updated context ('ctx') that contains our trace ID.
		// The original request does not get changed by this; we're simply creating a new version of it ('req').
		req := c.Request.WithContext(ctx)

		// Now, we want to carry forward this updated request (that has the new context) through our application.
		// So, we replace 'c.Request' (the original request) with 'req' (the new version with the updated context).
		// After this line, when we use 'c.Request' in this function or pass it to others, it'll be this new version
		// that carries our trace ID in its context.
		c.Request = req

		c.Next()
	}
}
