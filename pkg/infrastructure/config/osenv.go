package config

import (
	"fmt"
	"os"
	"strings"
)

type Environment struct {
	Port             string
	CORSAllowOrigins []string
}

func Get() (*Environment, error) {
	var env Environment
	var missed []string
	var corsAllowOrigins string

	for _, tmp := range []struct {
		field *string
		name  string
	}{
		{&env.Port, "PORT"},
		{&corsAllowOrigins, "CORS_ALLOW_ORIGINS"},
	} {
		v := os.Getenv(tmp.name)
		if v == "" {
			missed = append(missed, tmp.name)
		}
		*tmp.field = v
	}

	if 0 < len(missed) {
		return nil, fmt.Errorf("%s cannot be empty.", strings.Join(missed, ", "))
	}

	corsaos := strings.Split(corsAllowOrigins, ",")
	env.CORSAllowOrigins = append(env.CORSAllowOrigins, corsaos...)

	return &env, nil
}
