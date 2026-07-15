package main

var Hx_Obj_haxe_exceptions_notimplementedexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.exceptions.NotImplementedException",
)

type Hx_Obj_VTable_haxe_exceptions_notimplementedexception interface {
    Hx_Field_unwrap() any
    Hx_Field_toString() string
    Hx_Field_set_stack(stack *[]Hx_Enum_haxe_stackitem) *[]Hx_Enum_haxe_stackitem
    Hx_Field_get_stack() *[]Hx_Enum_haxe_stackitem
    Hx_Field_get_previous() struct { Value *Hx_Obj_haxe_exception; Valid bool }
    Hx_Field_get_native() any
    Hx_Field_get_message() string
    Hx_Field_details() string
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_exceptions_notimplementedexception struct {
    Hx_Obj_haxe_exceptions_posexception
    VTable Hx_Obj_VTable_haxe_exceptions_notimplementedexception
}

func Hx_Obj_haxe_exceptions_notimplementedexception_CreateEmptyInstance() *Hx_Obj_haxe_exceptions_notimplementedexception {
    obj := &Hx_Obj_haxe_exceptions_notimplementedexception{}
    obj.VTable = obj
    obj.Hx_Obj_haxe_exceptions_posexception.VTable = obj
    obj.Hx_Obj_haxe_exception.VTable = obj
    return obj
}

func Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, pos any) *Hx_Obj_haxe_exceptions_notimplementedexception {
    obj := Hx_Obj_haxe_exceptions_notimplementedexception_CreateEmptyInstance()
    obj.Hx_New(message, previous, pos)
    return obj
}

func (this *Hx_Obj_haxe_exceptions_notimplementedexception) Hx_New(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, pos any) {
    this.Hx_Obj_haxe_exceptions_posexception.Hx_New(message, previous, pos)
}

func (this *Hx_Obj_haxe_exceptions_notimplementedexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_exceptions_notimplementedexception_RTTI
}
