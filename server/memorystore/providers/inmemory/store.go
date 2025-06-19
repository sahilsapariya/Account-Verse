package inmemory

import (
	"fmt"
)


// GetStringStoreEnvVariable to get the env variable from string store object
func (c *provider) GetStringStoreEnvVariable(key string) (string, error) {
	res := c.envStore.Get(key)
	if res == nil {
		return "", nil
	}
	return fmt.Sprintf("%v", res), nil
}
