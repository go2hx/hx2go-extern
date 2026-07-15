package main

var Hx_Obj_go__map_map_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "go._Map.Map_Impl_",
)

type Hx_Obj_VTable_go__map_map_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go__map_map_impl_ struct {
    VTable Hx_Obj_VTable_go__map_map_impl_
}

func Hx_Obj_go__map_map_impl__CreateEmptyInstance() *Hx_Obj_go__map_map_impl_ {
    obj := &Hx_Obj_go__map_map_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_go__map_map_impl__CreateInstance() *Hx_Obj_go__map_map_impl_ {
    obj := Hx_Obj_go__map_map_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_go__map_map_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_go__map_map_impl__RTTI
}

func Hx_Field_go__map_map_impl___hx_new() map[any]any {
    return map[any]any{}
}
