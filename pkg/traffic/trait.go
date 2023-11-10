package traffic

import "io"

type Former interface {
	IntoForm() map[string]string
}

type MultiPartForm struct {
	Param    string
	Filename string
	Reader   io.Reader
}

func NewImageUploadForm(filename string, reader io.Reader) MultiPartForm {
	return MultiPartForm{
		Param:    "image",
		Filename: filename,
		Reader:   reader,
	}
}
