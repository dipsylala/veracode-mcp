package types

import "context"

// taskCancelRegistrarKey is the context key used to pass a cancel-func
// registrar into a task-augmented tool handler.
type taskCancelRegistrarKey struct{}

// WithTaskCancelRegistrar returns a context that lets a tool handler register
// a function to stop its in-flight work if the task is cancelled via
// tasks/cancel. Only set when a tool call is task-augmented.
func WithTaskCancelRegistrar(ctx context.Context, register func(cancel func() error)) context.Context {
	return context.WithValue(ctx, taskCancelRegistrarKey{}, register)
}

// RegisterTaskCancel lets a tool handler supply a function that stops its
// work early. It is a no-op if the call is not task-augmented.
func RegisterTaskCancel(ctx context.Context, cancel func() error) {
	if register, ok := ctx.Value(taskCancelRegistrarKey{}).(func(func() error)); ok {
		register(cancel)
	}
}
