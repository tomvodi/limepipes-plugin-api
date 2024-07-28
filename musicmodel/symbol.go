package limepipes_music_model

import (
	"fmt"
	"github.com/tomvodi/limepipes-music-model/musicmodel/symbols"
	"gopkg.in/yaml.v3"
)

// YamlSymbol is a helper struct for marshalling and unmarshalling the Symbol struct
// into YAML. It is used to determine which field of the Symbol struct is being marshalled
// so that a slice of symbols won't end in:
// - symbol:
//   - note:
//   - pitch: D
//
// but in
// - note:
//   - pitch: D
type YamlSymbol struct {
	Note *symbols.Note `yaml:"note,omitempty"`
	Rest *symbols.Rest `yaml:"rest,omitempty"`
}

func (i Symbol) MarshalYAML() (interface{}, error) {
	if i.GetNote() != nil {
		return YamlSymbol{Note: i.GetNote()}, nil
	}
	if i.GetRest() != nil {
		return YamlSymbol{Rest: i.GetRest()}, nil
	}

	return i.String(), nil
}

func (i *Symbol) UnmarshalYAML(value *yaml.Node) error {
	if len(value.Content) == 0 {
		return fmt.Errorf("symbol content is empty, can't unmarshal")
	}

	return nil
}
