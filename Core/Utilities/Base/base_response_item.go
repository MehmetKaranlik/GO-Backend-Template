package Base

type ResponseItem interface {
	// To give type constraint to the response item
	Conform()
}
