package env

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
	Unknown     Environment = "unknown"
)

func FromString(env string) Environment {
	switch env {
	case "development":
		return Development
	case "production":
		return Production
	default:
		return Unknown
	}
}
