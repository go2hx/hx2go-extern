package main

var Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.ds._ReadOnlyArray.ReadOnlyArray_Impl_",
)

type Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ struct {
    VTable Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_
}

func Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateEmptyInstance() *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ {
    obj := &Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateInstance() *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ {
    obj := Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__RTTI
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get_length(this any) int {
    return ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this, "length"))))
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get(this any, i int) any {
    return Hx_Field_go_haxe_hxdynamic_getArrayIndex(this, i)
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__concat(this any, a any) any {
    var newArr any = &([]any{}); _ = newArr
    var _hx_tmp_0 any = newArr; _ = _hx_tmp_0
    var _hx_tmp_1 []any = (*((newArr).(*[]any))); _ = _hx_tmp_1
    (*((_hx_tmp_0).(*[]any))) = append(_hx_tmp_1, (*((this).(*[]any)))...)
    var newArr1 any = newArr; _ = newArr1
    var _hx_tmp_2 any = newArr1; _ = _hx_tmp_2
    var _hx_tmp_3 []any = (*((newArr1).(*[]any))); _ = _hx_tmp_3
    (*((_hx_tmp_2).(*[]any))) = append(_hx_tmp_3, (*((a).(*[]any)))...)
    return newArr1
}
