package main

var Hx_Obj_haxe_exception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.Exception",
)

type Hx_Obj_VTable_haxe_exception interface {
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

type Hx_Obj_haxe_exception struct {
    VTable Hx_Obj_VTable_haxe_exception
    Hx_Field__message string
}

func Hx_Obj_haxe_exception_CreateEmptyInstance() *Hx_Obj_haxe_exception {
    obj := &Hx_Obj_haxe_exception{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_exception_CreateInstance(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, native any) *Hx_Obj_haxe_exception {
    obj := Hx_Obj_haxe_exception_CreateEmptyInstance()
    obj.Hx_New(message, previous, native)
    return obj
}

func (this *Hx_Obj_haxe_exception) Hx_New(message string, previous struct { Value *Hx_Obj_haxe_exception; Valid bool }, native any) {
    this.Hx_Field__message = message
}

func (this *Hx_Obj_haxe_exception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_exception_RTTI
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_message() string {
    return this.Hx_Field__message
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_stack() *[]Hx_Enum_haxe_stackitem {
    return &([]Hx_Enum_haxe_stackitem{})
}

func (this *Hx_Obj_haxe_exception) Hx_Field_set_stack(stack *[]Hx_Enum_haxe_stackitem) *[]Hx_Enum_haxe_stackitem {
    return &([]Hx_Enum_haxe_stackitem{})
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_previous() struct { Value *Hx_Obj_haxe_exception; Valid bool } {
    return struct { Value *Hx_Obj_haxe_exception; Valid bool }{}
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_native() any {
    return ((any)(0))
}

func (this *Hx_Obj_haxe_exception) Hx_Field_unwrap() any {
    return ((any)(0))
}

func (this *Hx_Obj_haxe_exception) Hx_Field_toString() string {
    return ("Error: " + this.VTable.Hx_Field_get_message())
}

func (this *Hx_Obj_haxe_exception) Hx_Field_details() string {
    var _hx_tmp_0 string = (this.VTable.Hx_Field_toString() + "\n"); _ = _hx_tmp_0
    return (_hx_tmp_0 + Hx_Field_haxe__callstack_callstack_impl__toString(this.VTable.Hx_Field_get_stack()))
}

func Hx_Field_haxe_exception_caught(value any) *Hx_Obj_haxe_exception {
    var _hx_tmp_0 string; _ = _hx_tmp_0
    if (Hx_Field_go_haxe_hxdynamic_equals(value, nil)) {
        _hx_tmp_0 = "null"
    } else {
        _hx_tmp_0 = Hx_Field_std_string(value)
    }

    return Hx_Obj_haxe_exception_CreateInstance(_hx_tmp_0, struct { Value *Hx_Obj_haxe_exception; Valid bool }{}, nil)
}

func Hx_Field_haxe_exception_thrown(value any) any {
    return ((any)(0))
}
