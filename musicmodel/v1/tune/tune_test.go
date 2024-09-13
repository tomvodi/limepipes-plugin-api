package tune

import (
	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"
	"testing"
)

func TestTune_ImportMessages(t *testing.T) {
	g := NewGomegaWithT(t)
	tu := &Tune{}
	g.Expect(tu.ImportMessages()).To(BeNil())
	msg1 := &measure.ParserMessage{
		Severity: measure.Severity_Error,
		Text:     "error",
	}
	msg2 := &measure.ParserMessage{
		Symbol: "asdf",
		Text:   "warning",
	}
	msg3 := &measure.ParserMessage{
		Severity: measure.Severity_Warning,
		Text:     "warning",
	}
	tu.Measures = []*measure.Measure{
		{
			ParserMessages: []*measure.ParserMessage{
				msg1,
				msg2,
			},
		},
		{
			ParserMessages: []*measure.ParserMessage{
				msg3,
			},
		},
	}

	g.Expect(tu.ImportMessages()).To(HaveLen(3))
	g.Expect(tu.ImportMessages()).To(ConsistOf(msg1, msg2, msg3))
}

func TestTune_FirstTimeSignature(t *testing.T) {
	g := NewGomegaWithT(t)
	tu := &Tune{}
	g.Expect(tu.FirstTimeSignature()).To(BeNil())
	ts1 := &measure.TimeSignature{
		Beats:    4,
		BeatType: 4,
	}
	ts2 := &measure.TimeSignature{
		Beats:    3,
		BeatType: 4,
	}

	tu.Measures = []*measure.Measure{
		{
			Time: ts1,
		},
		{
			Time: ts2,
		},
	}

	g.Expect(tu.FirstTimeSignature()).To(Equal(ts1))
}
