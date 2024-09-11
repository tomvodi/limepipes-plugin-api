package tuplet

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/boundary"
)

type Type uint8

const (
	NoType Type = iota
	Type23
	Type32
	Type43
	Type46
	Type53
	Type54
	Type64
	Type74
	Type76
)

func NewTuplet(bound boundary.Boundary, ttype Type) *Tuplet {
	tp := &Tuplet{
		BoundaryType: bound,
	}
	tp.VisibleNotes, tp.PlayedNotes = notesConfigFromType(ttype)

	return tp
}

// notesConfigFromType returns the visible notes and played notes for a given type
func notesConfigFromType(ttype Type) (uint32, uint32) {
	switch ttype {
	case Type23:
		return 2, 3
	case Type32:
		return 3, 2
	case Type43:
		return 4, 3
	case Type46:
		return 4, 6
	case Type53:
		return 5, 3
	case Type54:
		return 5, 4
	case Type64:
		return 6, 4
	case Type74:
		return 7, 4
	case Type76:
		return 7, 6
	default:
		panic("tuplet type missed in notesConfigFromType")
	}
}
