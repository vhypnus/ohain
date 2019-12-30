package repo

type node struct {
	// 定义一个能变化量量

	// shot for type . struct interface func method
	t int

	block interface  

	// edges array

}

// 重写equals 方法

type edge struct {

	// include,dependency
	relation int 

	start node

	end node
}

const (
	INCLUDE = ,
	DEPENENCY = 
)

// 

type method struct{

	node
}
