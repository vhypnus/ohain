package project

//
type Variable struct {

	// scope
	s int

	// type 
	t string

	// name
	n string

	// value expression
	v string

	// note
	note string
}

//
func GetVariable(line string) *Variable {
	var words []string = strings.Split(line,'')

	var v = &Variable {}
	return v
}