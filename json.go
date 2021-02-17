package tinytime

var layout = `"2006-01-02T15:04:05Z07:00"`

// MarshalJSON implements the json.Marshaler interface
func (tt TinyTime) MarshalJSON() ([]byte, error) {
	return []byte(tt.Format(layout)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tt *TinyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var err error
	*tt, err = Parse(layout, string(data))
	return err
}
