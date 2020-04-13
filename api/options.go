package api

type apiState int
type graphicsState int

const (
	asUninitialized apiState = 0
	asOptionsBlock  apiState = 1
	asWorldBlock    apiState = 2
	asReady         apiState = 3
)

// Options struct hows the options for the application
type Options struct {
	apiState
}

// Init initializes the Options struct
func (o *Options) Init() error {
	// Init Api State
	o.apiState = asReady
	// Init Render Options
	// Init Graphic State
	// Init SampledSpectrum
	// Init Parallel
	// Init Profiler
	return nil
}
