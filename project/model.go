package project



type Clazz struct {
	name string

	note string

	fields []string
}


//
// func NewModel(p string) string {
// 	rd, err := ioutil.ReadDir(p)
//     for _, fi := range rd {
//         if fi.IsDir() {
//             //todo
//         } else {
//         	var model = ""
//             var f *File = NewFile(p)

//             var s,e = f.GetBlock(1)
            
//             var line = f.GetCurrentLine(s)
//             // model name
//             var mn = modelName(line)

//             // line,note line,code line
//             var l,nl,cl = "","",""

//             //current index
//             var ci = 0
//             for {

//             	l = f.GetNextLine(s)

//             	if l[:2] == "//" || l[:2] == "/*" {
//             		nl += l
//             	} else {

//             		if cl != nil {
// 	            		// parse nl,cl
// 	            		var field = nil

// 	            		//reset
// 	            		var nl,cl = "",""
// 	            	} 
//             	}

            	

//             	if ci + len(l) > e {
//             		break 
//             	}
//             }


//         }
//     }
// }

// // return model name
// func modelName(s string) string {

// }


// // n note c code
// // format
// func field(n string,c string) string {

// }
