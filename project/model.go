package project



type Clazz struct {
	name string

	note string

	fields []string
}

func (m Model) tostring(){
	
}


//
func NewModel(p string) string {
	rd, err := ioutil.ReadDir(p)
    for _, fi := range rd {
        if fi.IsDir() {
            //todo
        } else {
        	var model = ""
            var f *File = NewFile(p)

            var s,e = f.Block(1)
            
            var cs,ce = f.CurrentLine(s)
            // model name
            var mn = modelName(line)

            //current index
            var p = ce+1
            //next return '\n'
            var nr = f.NextChar(ce,'\n')
            for ;;{

            	ns,ne = f.NextLine(p)

            	
            	
            	
            	//reset
            	p = ne + 1
            	if p > e {
            		break 
            	}

            }


        }
    }
}

// // return model name
// func modelName(s string) string {

// }


// // n note c code
// // format
// func field(n string,c string) string {

// }
