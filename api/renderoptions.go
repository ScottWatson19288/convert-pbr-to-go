package api

type renderoptions struct {
	transformStartTime float64 // 0
	transformEndTime   float64 // 1
	filterName         string  // box
	filmName           string  // image
	samplerName        string  // halton
	acceleratorName    string  // bvh
	integratorName     string  // path
	cameraName         string  // perspective

	// ParamSets whatever these are probably structs
	/*
		filterParams paramSet
		filmParams paramSet
		samplerParams paramSet
		acceleratorParams paramSet
		integratorParams paramSet
		cameraParams paramSet
	*/
	// cameraToWorld transformSet
	// lights *[]Light
	// primitives *[]Primitives
	// map[string] *[Medium]{}
	// map[string] *[Primitives]}{}
	haveScatteringMedia bool // false

	// RenderOptions Public Methods
	// MakeInegrator()
	// MakeScene()
	// MakeCamera()
}
