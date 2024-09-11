package tuplet

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/boundary"
	"reflect"
	"testing"
)

func TestNewTuplet(t *testing.T) {
	type args struct {
		bound boundary.Boundary
		ttype Type
	}
	tests := []struct {
		name string
		args args
		want *Tuplet
	}{
		{
			name: "23",
			args: args{
				bound: boundary.Boundary_NoBoundary,
				ttype: Type23,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_NoBoundary,
				VisibleNotes: 2,
				PlayedNotes:  3,
			},
		},
		{
			name: "32",
			args: args{
				bound: boundary.Boundary_Start,
				ttype: Type32,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_Start,
				VisibleNotes: 3,
				PlayedNotes:  2,
			},
		},
		{
			name: "43",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type43,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 4,
				PlayedNotes:  3,
			},
		},
		{
			name: "46",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type46,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 4,
				PlayedNotes:  6,
			},
		},
		{
			name: "53",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type53,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 5,
				PlayedNotes:  3,
			},
		},
		{
			name: "54",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type54,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 5,
				PlayedNotes:  4,
			},
		},
		{
			name: "64",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type64,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 6,
				PlayedNotes:  4,
			},
		},
		{
			name: "74",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type74,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 7,
				PlayedNotes:  4,
			},
		},
		{
			name: "76",
			args: args{
				bound: boundary.Boundary_End,
				ttype: Type76,
			},
			want: &Tuplet{
				BoundaryType: boundary.Boundary_End,
				VisibleNotes: 7,
				PlayedNotes:  6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTuplet(tt.args.bound, tt.args.ttype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTuplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
