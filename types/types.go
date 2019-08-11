package types

type FilePattern string
type File struct {
	Contents string
	Path     string
	// Mark this file as needs to be discarded (as opposed to just an empty file)
	Discarded bool
}
