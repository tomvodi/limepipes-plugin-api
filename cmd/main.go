package main

import (
	"fmt"
	limepipes_music_model "github.com/tomvodi/limepipes-music-model/musicmodel"
	"github.com/tomvodi/limepipes-music-model/musicmodel/length"
	"github.com/tomvodi/limepipes-music-model/musicmodel/pitch"
	"github.com/tomvodi/limepipes-music-model/musicmodel/symbols"

	"gopkg.in/yaml.v3"
)

func main() {
	notePitch := pitch.Pitch_HighA

	d, err := yaml.Marshal(&notePitch)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshaled Pitch: %s\n", string(d))

	pitchData := []byte("HighA")
	err = yaml.Unmarshal(pitchData, &notePitch)
	if err != nil {
		panic(err)
	}

	fmt.Println("Unmarshaled Pitch:", notePitch)

	rawNote := &symbols.Note{
		Pitch:  pitch.Pitch_D,
		Length: length.Length_Sixteenth,
		Dots:   1,
	}

	rawNoteD, err := yaml.Marshal(&rawNote)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshaled raw note: \n%s\n", rawNoteD)

	note := limepipes_music_model.Symbol_Note{
		Note: rawNote,
	}

	sym := limepipes_music_model.Symbol{
		Symbol: &note,
	}

	symD, err := yaml.Marshal(&sym)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshaled Symbol: \n%s\n", string(symD))

	restSym := limepipes_music_model.Symbol_Rest{
		Rest: &symbols.Rest{Length: length.Length_Eighth},
	}

	sym2 := limepipes_music_model.Symbol{Symbol: &restSym}

	syms := []limepipes_music_model.Symbol{
		sym,
		sym2,
	}

	symsD, err := yaml.Marshal(&syms)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshaled Symbols: \n%s\n", string(symsD))

	outSym := limepipes_music_model.Symbol{}
	noteData := []byte(
		`note:
    pitch: D
    length: Sixteenth
    dots: 1`)

	err = yaml.Unmarshal(noteData, &outSym)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshaled Symbol: \n%+v\n", outSym)
}
