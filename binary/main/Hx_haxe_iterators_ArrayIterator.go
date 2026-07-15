package main

var Hx_Obj_haxe_iterators_arrayiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.ArrayIterator",
)

type Hx_Obj_VTable_haxe_iterators_arrayiterator interface {
    Hx_Field_next() any
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_arrayiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_arrayiterator
    Hx_Field_array any
    Hx_Field_current int
}

func Hx_Obj_haxe_iterators_arrayiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_arrayiterator {
    obj := &Hx_Obj_haxe_iterators_arrayiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(array any) *Hx_Obj_haxe_iterators_arrayiterator {
    obj := Hx_Obj_haxe_iterators_arrayiterator_CreateEmptyInstance()
    obj.Hx_New(array)
    return obj
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_New(array any) {
    this.Hx_Field_array = array
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_arrayiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 any = ((any)(this.Hx_Field_current)); _ = _hx_tmp_0
    return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_lt(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_array, "length")))))
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field_next() any {
    var _hx_tmp_0 any = this.Hx_Field_array; _ = _hx_tmp_0
    var _hx_tmp_1 int = this.Hx_Field_current; _ = _hx_tmp_1
    this.Hx_Field_current = (this.Hx_Field_current + 1)
    return Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_0, _hx_tmp_1)
}
