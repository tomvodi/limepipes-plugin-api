package measure

import (
	"fmt"
)

func (x *TimeSignature) DisplayString() string {
	return fmt.Sprintf("%d/%d", x.Beats, x.BeatType)
}
