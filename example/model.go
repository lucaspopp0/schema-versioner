package example

import versioner "github.com/lucaspopp0/schema-versioner"

var (
	userVersionOpts = versioner.SemanticVersionOpts{
		NumParts: 3,
	}
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	versioner.Versioned[versioner.SemanticVersion]
}
