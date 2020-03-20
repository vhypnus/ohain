package doc


type Doc struct{

}

func NewDoc(path string){

}


func (d *Doc) NewSegement(title string,content) *Doc {

}

func (d *Doc) NewTable(title string,heads []string,rows []string) *Doc {

}

func (d *Doc) output() (bool,error) {}