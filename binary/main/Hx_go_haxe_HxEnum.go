package main

var Hx_Obj_go_haxe_hxenum_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go.haxe.HxEnum",
)

type Hx_Obj_VTable_go_haxe_hxenum interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxenum struct {
    VTable Hx_Obj_VTable_go_haxe_hxenum
    Hx_Field_name string
    Hx_Field_constructorNames *[]string
    Hx_Field_constructorArgCounts *[]int
    Hx_Field_createByIndex func(int, any) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
}

func Hx_Obj_go_haxe_hxenum_CreateEmptyInstance() *Hx_Obj_go_haxe_hxenum {
    obj := &Hx_Obj_go_haxe_hxenum{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go_haxe_hxenum_CreateInstance(name string, constructorNames *[]string, constructorArgCounts *[]int, createByIndex func(int, any) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue) *Hx_Obj_go_haxe_hxenum {
    obj := Hx_Obj_go_haxe_hxenum_CreateEmptyInstance()
    obj.Hx_New(name, constructorNames, constructorArgCounts, createByIndex)
    return obj
}

func (this *Hx_Obj_go_haxe_hxenum) Hx_New(name string, constructorNames *[]string, constructorArgCounts *[]int, createByIndex func(int, any) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue) {
    this.Hx_Field_name = name
    this.Hx_Field_constructorNames = constructorNames
    this.Hx_Field_constructorArgCounts = constructorArgCounts
    this.Hx_Field_createByIndex = createByIndex
}

func (this *Hx_Obj_go_haxe_hxenum) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go_haxe_hxenum_RTTI
}
