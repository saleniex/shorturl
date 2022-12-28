package shorturl

// IdGenerator provides ID generation service
type IdGenerator interface {
	// Generate new identifier
	Generate() string
}
