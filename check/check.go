package model


type Check struct {

	filepaths []string

}



func MaxLineCheck() {

}

//注释检查
func NoteCheck() {
	// 类注释

	// public 方法注释，如果是实现，需要从接口获取

	// 属性注释
}

func Check() (bool ,CheckError ){

	return false,nil
}