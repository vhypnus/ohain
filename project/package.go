package project


type Package struct {

	path string

	packages []Package

	files []File
}