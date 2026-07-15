package main

var Hx_Obj_go_haxe_hxarray_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go.haxe.HxArray",
)

type Hx_Obj_VTable_go_haxe_hxarray interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxarray struct {
    VTable Hx_Obj_VTable_go_haxe_hxarray
}

func Hx_Obj_go_haxe_hxarray_CreateEmptyInstance() *Hx_Obj_go_haxe_hxarray {
    obj := &Hx_Obj_go_haxe_hxarray{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go_haxe_hxarray_CreateInstance() *Hx_Obj_go_haxe_hxarray {
    obj := Hx_Obj_go_haxe_hxarray_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_go_haxe_hxarray) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go_haxe_hxarray_RTTI
}
