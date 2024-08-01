package symbols

func (x *Symbol) IsNote() bool {
	return x.Note != nil
}

func (x *Symbol) IsValidNote() bool {
	if x.Note != nil {
		return x.Note.IsValid()
	}

	return false
}
