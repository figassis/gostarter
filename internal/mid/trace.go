package mid

import (
	"context"
	"fmt"
	"geeks-accelerator/oss/saas-starter-kit/internal/platform/web/webcontext"
	"net/http"

	"geeks-accelerator/oss/saas-starter-kit/internal/platform/web"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Trace adds the base tracing info for requests
func Trace() web.Middleware {

	// This is the actual middleware function to be executed.
	f := func(before web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			// Span options with request info
			opts := []ddtrace.StartSpanOption{
				tracer.SpanType(ext.SpanTypeWeb),
				tracer.ResourceName(r.URL.Path),
				tracer.Tag(ext.HTTPMethod, r.Method),
				tracer.Tag(ext.HTTPURL, r.RequestURI),
			}

			// Continue server side request tracing from previous request.
			if spanctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(r.Header)); err == nil {
				opts = append(opts, tracer.ChildOf(spanctx))
			}

			// Start the span for tracking
			span, ctx := tracer.StartSpanFromContext(ctx, "http.request", opts...)
			defer span.Finish()

			// Load the context values.
			v, err := webcontext.ContextValues(ctx)
			if err != nil {
				return err
			}

			v.TraceID = span.Context().TraceID()
			v.SpanID = span.Context().SpanID()

			// Execute the request handler
			err = before(ctx, w, r, params)

			// Set the span status code for the trace
			span.SetTag(ext.HTTPCode, v.StatusCode)

			// If there was an error, append it to the span
			if err != nil {
				span.SetTag(ext.Error, fmt.Sprintf("%+v", err))
			}

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return f
}
