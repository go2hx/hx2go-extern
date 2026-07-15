package main

var Hx_Obj_haxe__int32_int32_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe._Int32.Int32_Impl_",
)

type Hx_Obj_VTable_haxe__int32_int32_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__int32_int32_impl_ struct {
    VTable Hx_Obj_VTable_haxe__int32_int32_impl_
}

func Hx_Obj_haxe__int32_int32_impl__CreateEmptyInstance() *Hx_Obj_haxe__int32_int32_impl_ {
    obj := &Hx_Obj_haxe__int32_int32_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe__int32_int32_impl__CreateInstance() *Hx_Obj_haxe__int32_int32_impl_ {
    obj := Hx_Obj_haxe__int32_int32_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe__int32_int32_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe__int32_int32_impl__RTTI
}

func Hx_Field_haxe__int32_int32_impl__negate(this int) int {
    return (^this + 1)
}

func Hx_Field_haxe__int32_int32_impl__preIncrement(this int) int {
    this = (this + 1)
    var _hx_tmp_0 int = this; _ = _hx_tmp_0
    this = _hx_tmp_0
    return this
}

func Hx_Field_haxe__int32_int32_impl__postIncrement(this int) int {
    var _hx_tmp_0 int = this; _ = _hx_tmp_0
    this = (this + 1)
    var ret int = _hx_tmp_0; _ = ret
    return ret
}

func Hx_Field_haxe__int32_int32_impl__preDecrement(this int) int {
    this = (this - 1)
    var _hx_tmp_0 int = this; _ = _hx_tmp_0
    this = _hx_tmp_0
    return this
}

func Hx_Field_haxe__int32_int32_impl__postDecrement(this int) int {
    var _hx_tmp_0 int = this; _ = _hx_tmp_0
    this = (this - 1)
    var ret int = _hx_tmp_0; _ = ret
    return ret
}

func Hx_Field_haxe__int32_int32_impl__add(a int, b int) int {
    return (((int)(a)) + ((int)(b)))
}

func Hx_Field_haxe__int32_int32_impl__addInt(a int, b int) int {
    return (((int)(a)) + b)
}

func Hx_Field_haxe__int32_int32_impl__sub(a int, b int) int {
    return (((int)(a)) - ((int)(b)))
}

func Hx_Field_haxe__int32_int32_impl__subInt(a int, b int) int {
    return (((int)(a)) - b)
}

func Hx_Field_haxe__int32_int32_impl__intSub(a int, b int) int {
    return (a - ((int)(b)))
}

func Hx_Field_haxe__int32_int32_impl__toFloat(this int) float64 {
    return ((float64)(this))
}

func Hx_Field_haxe__int32_int32_impl__ucompare(a int, b int) int {
    if ((a < 0)) {
        if ((b < 0)) {
            return ((int)((((int)(^b)) - ((int)(^a)))))
        } else {
            return 1
        }
    }

    if ((b < 0)) {
        return -1
    } else {
        return ((int)((((int)(a)) - ((int)(b)))))
    }
}

func Hx_Field_haxe__int32_int32_impl__clamp(x int) int {
    return x
}
