package base

import "fmt"

func init() {
	fmt.Println("class init")
	testObj()
}

type Obj_st struct {
	objId   int
	objName string
}

func testObj() {
	/*
		创建对象，对表对象引用
		a := Obj_st{}

		两种写法的意义一样，代表对象指针
		aa := new(Obj_st)
		bb := &Obj_st{}
	*/

	useObj := Obj_st{}
	tmp := Obj_st{objId: 1, objName: "abc"}
	fmt.Println("objId unset", useObj.objId)
	useObj.copyObjName(tmp)
	fmt.Println("objId set", useObj.objId)
	retid, retname := useObj.GetObj_stInfo()
	fmt.Println("obj info", retid, retname)
}

/**
getObjName
返回传入参数的 objName

getObjNameOther
返回调用对象的 objName
*/
func getObjName(obj Obj_st) string {
	return obj.objName
}

func (obj Obj_st) getObjNameOther() string {
	return obj.objName
}

func (obj Obj_st) getObjId() int {
	return obj.objId
}

func (obj Obj_st) GetObj_stInfo() (int, string) {
	return obj.objId, obj.objName
}

func (obj *Obj_st) copyObjName(pa Obj_st) {
	obj.objId = pa.objId
	obj.objName = pa.objName
}
