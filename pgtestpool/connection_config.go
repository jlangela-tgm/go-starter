package pgtestpool

import (
	"fmt"
	"sort"
	"strings"
)

type ConnectionConfig struct {
	Host             string
	Port             int
	Username         string
	Password         string
	Database         string
	AdditionalParams map[string]string // Optional additional connection parameters mapped into the connection string
}

// Generates a connection string to be passed to sql.Open or equivalents, assuming Postgres syntax
func (c ConnectionConfig) ConnectionString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Host, c.Port, c.Username, c.Password, c.Database))

	if _, ok := c.AdditionalParams["sslmode"]; !ok {
		b.WriteString(" sslmode=disable")
	}

	if len(c.AdditionalParams) > 0 {
		params := make([]string, 0, len(c.AdditionalParams))
		for param := range c.AdditionalParams {
			params = append(params, param)
		}

		sort.Strings(params)

		for _, param := range params {
			fmt.Fprintf(&b, " %s=%s", param, c.AdditionalParams[param])
		}
	}

	return b.String()
}
