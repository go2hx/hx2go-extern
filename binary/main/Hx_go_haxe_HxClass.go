package main

var Hx_Obj_go_haxe_hxclass_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go.haxe.HxClass",
)

type Hx_Obj_VTable_go_haxe_hxclass interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxclass struct {
    VTable Hx_Obj_VTable_go_haxe_hxclass
    Hx_Field_name string
}

func Hx_Obj_go_haxe_hxclass_CreateEmptyInstance() *Hx_Obj_go_haxe_hxclass {
    obj := &Hx_Obj_go_haxe_hxclass{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go_haxe_hxclass_CreateInstance(name string) *Hx_Obj_go_haxe_hxclass {
    obj := Hx_Obj_go_haxe_hxclass_CreateEmptyInstance()
    obj.Hx_New(name)
    return obj
}

func (this *Hx_Obj_go_haxe_hxclass) Hx_New(name string) {
    this.Hx_Field_name = name
}

func (this *Hx_Obj_go_haxe_hxclass) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go_haxe_hxclass_RTTI
}
