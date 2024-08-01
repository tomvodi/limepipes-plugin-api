package main

import (
	"fmt"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
)

func main() {
	notePitch := pitch.Pitch_HighA

	//test := yaml.InterfaceMarshaler()
	//test2 := yaml.InterfaceUnmarshaler()

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

	note := &symbols.Note{
		Pitch:  pitch.Pitch_D,
		Length: length.Length_Sixteenth,
		Dots:   1,
	}

	rawNoteD, err := yaml.Marshal(&note)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshaled raw note: \n%s\n", rawNoteD)

	sym := symbols.Symbol{
		Note: note,
	}

	symD, err := yaml.Marshal(&sym)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshaled Symbol: \n%s\n", string(symD))

	restSym := &symbols.Rest{Length: length.Length_Eighth}

	sym2 := symbols.Symbol{Rest: restSym}

	syms := []symbols.Symbol{
		sym,
		sym2,
	}

	symsD, err := yaml.Marshal(&syms)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Marshaled Symbols: \n%s\n", string(symsD))

	outSym := symbols.Symbol{}
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
