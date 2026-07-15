package main

type Hx_Enum_go_resultkind interface {
    Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
    M_Hx_Enum_go_resultkind()
}

type Hx_Enum_go_resultkind_Ok struct {
    Hx_Field_r any
}

func (this Hx_Enum_go_resultkind_Ok) M_Hx_Enum_go_resultkind() {}
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumIndex() int { return 0 }
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_go_resultkind_RTTI }
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_r) }) }
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_r)
        default: return nil
    }
}

type Hx_Enum_go_resultkind_Err struct {
    Hx_Field_e any
}

func (this Hx_Enum_go_resultkind_Err) M_Hx_Enum_go_resultkind() {}
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumIndex() int { return 1 }
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum { return Hx_Enum_go_resultkind_RTTI }
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumParams() any { return &([]any{ any(this.Hx_Field_e) }) }
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumParameter(index int) any {
    switch index {
        case 0: return any(this.Hx_Field_e)
        default: return nil
    }
}

var Hx_Enum_go_resultkind_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
    "go.ResultKind",
    &([]string{ "Ok", "Err" }),
    &([]int{ 1, 1 }),
    func (index int, params any) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
        return nil
    },
)
