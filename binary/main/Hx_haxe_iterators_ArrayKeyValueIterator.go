package main

var Hx_Obj_haxe_iterators_arraykeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.ArrayKeyValueIterator",
)

type Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator interface {
    Hx_Field_next() any
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_arraykeyvalueiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator
    Hx_Field_current int
    Hx_Field_array any
}

func Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
    obj := &Hx_Obj_haxe_iterators_arraykeyvalueiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateInstance(array any) *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
    obj := Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateEmptyInstance()
    obj.Hx_New(array)
    return obj
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_New(array any) {
    this.Hx_Field_array = array
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_arraykeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 int = this.Hx_Field_current; _ = _hx_tmp_0
    return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_array))
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field_next() any {
    var _hx_tmp_0 any = Hx_Field_go_haxe_hxdynamic_getArrayIndex(this.Hx_Field_array, this.Hx_Field_current); _ = _hx_tmp_0
    var _hx_tmp_1 int = this.Hx_Field_current; _ = _hx_tmp_1
    this.Hx_Field_current = (this.Hx_Field_current + 1)
    return any(map[string]any{ "value": _hx_tmp_0, "key": ((any)(_hx_tmp_1)) })
}
