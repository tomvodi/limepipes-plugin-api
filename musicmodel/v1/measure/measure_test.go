package measure

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
	"testing"
)

func TestMeasure_IsNil(t *testing.T) {
	type fields struct {
		LeftBarline    *barline.Barline
		RightBarline   *barline.Barline
		Time           *TimeSignature
		Symbols        []*symbols.Symbol
		Comments       []string
		InlineText     []string
		ImportMessages []*ImportMessage
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   true,
		},
		{
			name: "non-empty",
			fields: fields{
				LeftBarline: &barline.Barline{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Measure{
				LeftBarline:    tt.fields.LeftBarline,
				RightBarline:   tt.fields.RightBarline,
				Time:           tt.fields.Time,
				Symbols:        tt.fields.Symbols,
				Comments:       tt.fields.Comments,
				InlineText:     tt.fields.InlineText,
				ImportMessages: tt.fields.ImportMessages,
			}
			if got := x.IsNil(); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeasure_AddMessage(t *testing.T) {
	g := NewGomegaWithT(t)
	m := &Measure{}
	g.Expect(m.ImportMessages).To(BeNil())
	im := &ImportMessage{
		Symbol:   "sym",
		Severity: 0,
		Text:     "",
		Fix:      0,
	}
	m.AddMessage(im)
	g.Expect(m.ImportMessages).To(HaveLen(1))
	g.Expect(m.ImportMessages[0]).To(Equal(im))
}
