package client

import (
	"os"
	"strings"
)

var DEBUG = strings.TrimSpace(os.Getenv("DEBUG")) != ""
