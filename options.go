package telefil

type (
	Option  func(*options) error
	options struct {
		api string
	}
)

func newOptions(o ...Option) (*options, error) {
	return &options{
		api: `https://api.node.glif.io`,
	}, nil
}

// WithFilecoinAPI sets the Filecoin API endpoint.
// Defaults to https://api.node.glif.io
func WithFilecoinAPI(url string) Option {
	return func(o *options) error {
		o.api = url
		return nil
	}
}
