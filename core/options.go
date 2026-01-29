package core

type renderConfig struct {
	Indent bool
}

// Option applies a configuration to the renderer.
type Option func(*renderConfig)

// WithIndent enables indentation of elements with two spaces per depth level.
func WithIndent() Option {
	return func(cfg *renderConfig) {
		cfg.Indent = true
	}
}
