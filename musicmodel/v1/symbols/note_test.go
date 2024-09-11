package symbols

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/length"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/pitch"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/accidental"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/embellishment"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/movement"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols/tie"
	"testing"
)

func TestNote_IsIncomplete(t *testing.T) {
	type fields struct {
		Pitch         pitch.Pitch
		Length        length.Length
		Dots          uint32
		Accidental    accidental.Accidental
		Fermata       bool
		Tie           tie.Tie
		Embellishment *embellishment.Embellishment
		Movement      *movement.Movement
		Comment       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "pitch and no length",
			fields: fields{
				Pitch: pitch.Pitch_D,
			},
			want: false,
		},
		{
			name: "no pitch but length",
			fields: fields{
				Length: length.Length_Sixteenth,
			},
			want: false,
		},
		{
			name: "embellishment",
			fields: fields{
				Embellishment: &embellishment.Embellishment{},
			},
			want: true,
		},
		{
			name: "movement",
			fields: fields{
				Movement: &movement.Movement{},
			},
			want: true,
		},
		{
			name: "accidental",
			fields: fields{
				Accidental: accidental.Accidental_Sharp,
			},
			want: true,
		},
		{
			name: "tie",
			fields: fields{
				Tie: tie.Tie_Start,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Note{
				Pitch:         tt.fields.Pitch,
				Length:        tt.fields.Length,
				Dots:          tt.fields.Dots,
				Accidental:    tt.fields.Accidental,
				Fermata:       tt.fields.Fermata,
				Tie:           tt.fields.Tie,
				Embellishment: tt.fields.Embellishment,
				Movement:      tt.fields.Movement,
				Comment:       tt.fields.Comment,
			}
			if got := x.IsIncomplete(); got != tt.want {
				t.Errorf("IsIncomplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_HasPitchAndLength(t *testing.T) {
	g := NewGomegaWithT(t)
	n := &Note{}
	g.Expect(n.HasPitchAndLength()).To(BeFalse())
	n.Pitch = pitch.Pitch_D
	g.Expect(n.HasPitchAndLength()).To(BeFalse())
	n.Length = length.Length_Sixteenth
	g.Expect(n.HasPitchAndLength()).To(BeTrue())
}

func TestNote_HasPitchOrLength(t *testing.T) {
	g := NewGomegaWithT(t)
	n := &Note{}
	g.Expect(n.HasPitchOrLength()).To(BeFalse())
	n.Pitch = pitch.Pitch_D
	g.Expect(n.HasPitchOrLength()).To(BeTrue())
	n.Pitch = pitch.Pitch_NoPitch
	n.Length = length.Length_Sixteenth
	g.Expect(n.HasPitchOrLength()).To(BeTrue())
}

func TestNote_IsOnlyAccidental(t *testing.T) {
	g := NewGomegaWithT(t)
	n := &Note{}
	g.Expect(n.IsOnlyAccidental()).To(BeFalse())

	n.Accidental = accidental.Accidental_Sharp
	g.Expect(n.IsOnlyAccidental()).To(BeTrue())

	n.Pitch = pitch.Pitch_D
	g.Expect(n.IsOnlyAccidental()).To(BeFalse())
}
