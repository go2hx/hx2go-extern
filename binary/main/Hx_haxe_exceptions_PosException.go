package main

var Hx_Obj_haxe_exceptions_posexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.exceptions.PosException",
)

type Hx_Obj_VTable_haxe_exceptions_posexception interface {
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

type Hx_Obj_haxe_exceptions_posexception struct {
    Hx_Obj_haxe_exception
    VTable Hx_Obj_VTable_haxe_exceptions_posexception
    Hx_Field_posInfos any
}

func Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance() *Hx_Obj_haxe_exceptions_posexception {
    obj := &Hx_Obj_haxe_exceptions_posexception{}
    obj.VTable = obj
    obj.Hx_Obj_haxe_exception.VTable = obj
    return obj
}

func Hx_Obj_haxe_exceptions_posexception_CreateInstance(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, pos any) *Hx_Obj_haxe_exceptions_posexception {
    obj := Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance()
    obj.Hx_New(message, previous, pos)
    return obj
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_New(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, pos any) {
    this.Hx_Obj_haxe_exception.Hx_New(message, previous, nil)
    if (Hx_Field_go_haxe_hxdynamic_equals(pos, nil)) {
        this.Hx_Field_posInfos = any(map[string]any{ "fileName": ((any)("(unknown)")), "lineNumber": ((any)(0)), "className": ((any)("(unknown)")), "methodName": ((any)("(unknown)")) })
    } else {
        this.Hx_Field_posInfos = pos
    }
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_exceptions_posexception_RTTI
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_Field_toString() string {
    var _hx_tmp_3 any = ((any)((("" + this.Hx_Obj_haxe_exception.Hx_Field_toString()) + " in "))); _ = _hx_tmp_3
    var _hx_tmp_2 any = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "className")), ((any)("."))); _ = _hx_tmp_2
    var _hx_tmp_1 any = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "methodName")), ((any)(" at "))); _ = _hx_tmp_1
    var _hx_tmp_0 any = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "fileName")), ((any)(":"))); _ = _hx_tmp_0
    return Hx_Field_std_string(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "lineNumber")))
}
