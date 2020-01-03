package project

type Block struct {
	
	// package --> file path
	p string

	// block type virtual,reality
	t int

	// block name. first param is class
	n string

	// input. type array
	i []string

	//output 
	o []string

	//start index
	s int

	//end index 
	e int
}


func (b Block) GetCodeBlockDependency(){

}




