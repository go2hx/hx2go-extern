package main

var Hx_Obj_haxe_iterators_restkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.RestKeyValueIterator",
)

type Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator interface {
    Hx_Field_next() any
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_restkeyvalueiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator
    Hx_Field_args any
    Hx_Field_current int
}

func Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_restkeyvalueiterator {
    obj := &Hx_Obj_haxe_iterators_restkeyvalueiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateInstance(args any) *Hx_Obj_haxe_iterators_restkeyvalueiterator {
    obj := Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateEmptyInstance()
    obj.Hx_New(args)
    return obj
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_New(args any) {
    this.Hx_Field_args = (args).(any)
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_restkeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 int = this.Hx_Field_current; _ = _hx_tmp_0
    return (_hx_tmp_0 < ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(((any)(this.Hx_Field_args)), "length")))))
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field_next() any {
    var _hx_tmp_0 any = ((any)(this.Hx_Field_current)); _ = _hx_tmp_0
    var _hx_tmp_1 any = ((any)(this.Hx_Field_args)); _ = _hx_tmp_1
    var _hx_tmp_2 int = this.Hx_Field_current; _ = _hx_tmp_2
    this.Hx_Field_current = (this.Hx_Field_current + 1)
    return any(map[string]any{ "key": _hx_tmp_0, "value": Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_1, _hx_tmp_2) })
}
