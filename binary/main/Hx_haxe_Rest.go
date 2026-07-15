package main

type Hx_Typedef_haxe__rest_nativerest = any

var Hx_Obj_haxe__rest_rest_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe._Rest.Rest_Impl_",
)

type Hx_Obj_VTable_haxe__rest_rest_impl_ interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__rest_rest_impl_ struct {
    VTable Hx_Obj_VTable_haxe__rest_rest_impl_
}

func Hx_Obj_haxe__rest_rest_impl__CreateEmptyInstance() *Hx_Obj_haxe__rest_rest_impl_ {
    obj := &Hx_Obj_haxe__rest_rest_impl_{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe__rest_rest_impl__CreateInstance() *Hx_Obj_haxe__rest_rest_impl_ {
    obj := Hx_Obj_haxe__rest_rest_impl__CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe__rest_rest_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe__rest_rest_impl__RTTI
}

func Hx_Field_haxe__rest_rest_impl__get_length(this any) int {
    return ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this, "length"))))
}

func Hx_Field_haxe__rest_rest_impl__of(array any) any {
    return (array).(any)
}

func Hx_Field_haxe__rest_rest_impl___hx_new(array any) any {
    return (array).(any)
}

func Hx_Field_haxe__rest_rest_impl__get(this any, index int) any {
    return Hx_Field_go_haxe_hxdynamic_getArrayIndex(this, index)
}

func Hx_Field_haxe__rest_rest_impl__toArray(this any) any {
    var slice []any = Hx_Field_go_haxe_hxdynamic_toAnySlice(this); _ = slice
    var self []any = slice; _ = self
    return (&self)
}

func Hx_Field_haxe__rest_rest_impl__iterator(this any) *Hx_Obj_haxe_iterators_restiterator {
    return Hx_Obj_haxe_iterators_restiterator_CreateInstance(this)
}

func Hx_Field_haxe__rest_rest_impl__keyValueIterator(this any) *Hx_Obj_haxe_iterators_restkeyvalueiterator {
    return Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateInstance(this)
}

func Hx_Field_haxe__rest_rest_impl__append(this any, item any) any {
    var slice []any = Hx_Field_go_haxe_hxdynamic_toAnySlice(((any)(this))); _ = slice
    var self []any = slice; _ = self
    var result any = (&self); _ = result
    {
        var data []any = (*((result).(*[]any))); _ = data
        var _hx_tmp_0 any = result; _ = _hx_tmp_0
        (*((_hx_tmp_0).(*[]any))) = append(data, item)
        var _hx_tmp_1 int = len(data); _ = _hx_tmp_1
        var this2 int = (_hx_tmp_1 + int(1)); _ = this2
    }

    return (result).(any)
}

func Hx_Field_haxe__rest_rest_impl__prepend(this any, item any) any {
    var slice []any = Hx_Field_go_haxe_hxdynamic_toAnySlice(((any)(this))); _ = slice
    var self []any = slice; _ = self
    var result any = (&self); _ = result
    var _hx_tmp_0 any = result; _ = _hx_tmp_0
    var _hx_tmp_1 []any = (*&([]any{ item })); _ = _hx_tmp_1
    (*((_hx_tmp_0).(*[]any))) = append(_hx_tmp_1, (*((result).(*[]any)))...)
    return (result).(any)
}

func Hx_Field_haxe__rest_rest_impl__toString(this any) string {
    return Hx_Field_std_string(this)
}
