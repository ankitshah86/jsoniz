package config

import (
	"github.com/ankitshah86/jsoniz/internal/utils"
)

var enableJsonDuplicateKeyCheck = utils.GetEnvBool("ENABLE_JSON_DUPLICATE_KEY_CHECK", false)

func IsJsonDuplicateKeyCheckEnabled() bool {
	return enableJsonDuplicateKeyCheck
}
