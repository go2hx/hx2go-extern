package main

var Hx_Obj_haxe_iterators_restiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.iterators.RestIterator",
)

type Hx_Obj_VTable_haxe_iterators_restiterator interface {
    Hx_Field_next() any
    Hx_Field_hasNext() bool
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_restiterator struct {
    VTable Hx_Obj_VTable_haxe_iterators_restiterator
    Hx_Field_args any
    Hx_Field_current int
}

func Hx_Obj_haxe_iterators_restiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_restiterator {
    obj := &Hx_Obj_haxe_iterators_restiterator{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_iterators_restiterator_CreateInstance(args any) *Hx_Obj_haxe_iterators_restiterator {
    obj := Hx_Obj_haxe_iterators_restiterator_CreateEmptyInstance()
    obj.Hx_New(args)
    return obj
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_New(args any) {
    this.Hx_Field_args = (args).(any)
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_iterators_restiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field_hasNext() bool {
    var _hx_tmp_0 int = this.Hx_Field_current; _ = _hx_tmp_0
    return (_hx_tmp_0 < ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(((any)(this.Hx_Field_args)), "length")))))
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field_next() any {
    var _hx_tmp_0 any = ((any)(this.Hx_Field_args)); _ = _hx_tmp_0
    var _hx_tmp_1 int = this.Hx_Field_current; _ = _hx_tmp_1
    this.Hx_Field_current = (this.Hx_Field_current + 1)
    return Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_0, _hx_tmp_1)
}
