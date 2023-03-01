package pufs

import (
	"fmt"

	"github.com/StephenGriese/kitty/kitty"
)

func Pufs() string {
	return fmt.Sprintf("Pufs returns kitty number: %v", kitty.KittyValue)
}
