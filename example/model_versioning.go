package example

import versioner "github.com/lucaspopp0/schema-versioner"

var (
	userVersions = []versioner.SemanticVersion{
		versioner.NewSemanticVersion(userVersionOpts, 0, 0, 0),
		versioner.NewSemanticVersion(userVersionOpts, 0, 1, 0),
	}
)

// Model version 0.0.0 of the User struct
type user_0_0_0 struct {
	Name  string `json:"name"`
	Email string `json:"Email"`

	versioner.Versioned[versioner.SemanticVersion]
}

type user_0_0_1 = User
