package aws

import (
	"os"
	"testing"
)

// Custom environment variables used in the Terraform AWS Provider testing.
// Additions should also be documented in the Environment Variable Dictionary
// of the Maintainers Guide: docs/MAINTAINING.md
const (
	// Default static credential identifier for tests (AWS Go SDK does not provide this as constant)
	// See also AWS_SECRET_ACCESS_KEY and AWS_PROFILE
	EnvVarAwsAccessKeyId = "AWS_ACCESS_KEY_ID"

	// For tests using an alternate AWS account, the equivalent of AWS_ACCESS_KEY_ID for that account
	EnvVarAwsAlternateAccessKeyId = "AWS_ALTERNATE_ACCESS_KEY_ID"

	// For tests using an alternate AWS account, the equivalent of AWS_PROFILE for that account
	EnvVarAwsAlternateProfile = "AWS_PROFILE"

	// For tests using an alternate AWS region, the equivalent of AWS_DEFAULT_REGION for that account
	EnvVarAwsAlternateRegion = "AWS_ALTERNATE_REGION"

	// For tests using an alternate AWS account, the equivalent of AWS_SECRET_ACCESS_KEY for that account
	EnvVarAwsAlternateSecretAccessKey = "AWS_ALTERNATE_SECRET_ACCESS_KEY"

	// Container credentials endpoint
	// See also AWS_ACCESS_KEY_ID and AWS_PROFILE
	EnvVarAwsContainerCredentialsFullUri = "AWS_CONTAINER_CREDENTIALS_FULL_URI"

	// Default AWS region for tests (AWS Go SDK does not provide this as constant)
	EnvVarAwsDefaultRegion = "AWS_DEFAULT_REGION"

	// Default AWS shared configuration profile for tests (AWS Go SDK does not provide this as constant)
	EnvVarAwsProfile = "AWS_PROFILE"

	// Default static credential value for tests (AWS Go SDK does not provide this as constant)
	// See also AWS_ACCESS_KEY_ID and AWS_PROFILE
	EnvVarAwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"

	// For tests using a third AWS region, the equivalent of AWS_DEFAULT_REGION for that region
	EnvVarAwsThirdRegion = "AWS_THIRD_REGION"

	// For tests requiring GitHub permissions
	EnvVarGithubToken = "GITHUB_TOKEN"
)

// TestAccEnvironmentVariableSetPreCheck verifies that an environment variable is non-empty or skips the test.
func TestAccEnvironmentVariableSetPreCheck(t *testing.T, variable string, usageMessage string) string {
	value := os.Getenv(variable)

	if value == "" {
		t.Skipf("skipping tests; environment variable %s must be set. Usage: %s", variable, usageMessage)
	}

	return value
}
