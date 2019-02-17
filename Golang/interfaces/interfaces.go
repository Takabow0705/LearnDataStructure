package interfaces

type Queue interface{
	Add(x interface{})
	Remove()
}

type List interface{
	Size()
	Get(i int)
	Set(i int,x interface{})
	Add(i int,x interface{})
	Remove(i int)
}

type Set interface{
	Size()
	Add()
	Remove()
	Find()
}

