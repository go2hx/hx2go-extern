package main

var Hx_Obj__any_any_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "_Any.Any_Impl_",
)

type Hx_Obj_VTable__any_any_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj__any_any_impl_ struct {
    VTable Hx_Obj_VTable__any_any_impl_
}

func Hx_Obj__any_any_impl__CreateEmptyInstance() *Hx_Obj__any_any_impl_ {
    obj := &Hx_Obj__any_any_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj__any_any_impl__CreateInstance() *Hx_Obj__any_any_impl_ {
    obj := Hx_Obj__any_any_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj__any_any_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj__any_any_impl__RTTI
}
