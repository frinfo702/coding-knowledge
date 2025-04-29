package errorsample

import (
    "errors"
    "fmt"
    "os"
)

var ErrConfig = errors.New("config error")

func LoadConfig(path string) error {
    if _, err := os.Stat(path); err != nil {
        return fmt.Errorf("%w: %v", ErrConfig, err)
    }
    // あ
    return nil
}
