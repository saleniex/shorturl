package params

const Repository = "REPOSITORY"
const ListenAddr = "LISTEN_ADDR"

// Params interface defines methods to retrieve parameters independent from parameter storage implementation
type Params interface {
	// GetWithDefault parameter value by name with default value being provided
	GetWithDefault(name, defaultVal string) string

	// Get parameter value
	Get(name string) string

	// GetInt return parameter value as integer. In case parameter not found error !nil
	GetInt(name string) (int, error)

	// GetIntWithDefault get parameter value with provided default value
	GetIntWithDefault(name string, defaultVal int) int
}
