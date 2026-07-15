package main

var Hx_Obj_ereg_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "EReg",
)

type Hx_Obj_VTable_ereg interface {
    Hx_Field_split(s string) *[]string
    Hx_Field_replace(s string, by string) string
    Hx_Field_matchedRight() string
    Hx_Field_matchedPos() any
    Hx_Field_matchedNum() int
    Hx_Field_matchedLeft() string
    Hx_Field_matched(n int) string
    Hx_Field_matchSub(s string, pos int, _hx_reserved_len int) bool
    Hx_Field_match(s string) bool
    Hx_Field_map(s string, f func(*Hx_Obj_ereg) string) string
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_ereg struct {
    VTable Hx_Obj_VTable_ereg
}

func Hx_Obj_ereg_CreateEmptyInstance() *Hx_Obj_ereg {
    obj := &Hx_Obj_ereg{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_ereg_CreateInstance(r string, opt string) *Hx_Obj_ereg {
    obj := Hx_Obj_ereg_CreateEmptyInstance()
    obj.Hx_New(r, opt)
    return obj
}

func (this *Hx_Obj_ereg) Hx_New(r string, opt string) {
    panic(Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance("Regular expressions are not implemented for this platform", struct { Value *Hx_Obj_haxe_exception; Valid bool }{}, any(map[string]any{ "fileName": ((any)("EReg.hx")), "lineNumber": ((any)(48)), "className": ((any)("EReg")), "methodName": ((any)("new")) })))
}

func (this *Hx_Obj_ereg) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_ereg_RTTI
}

func (this *Hx_Obj_ereg) Hx_Field_match(s string) bool {
    return false
}

func (this *Hx_Obj_ereg) Hx_Field_matched(n int) string {
    return ``
}

func (this *Hx_Obj_ereg) Hx_Field_matchedLeft() string {
    return ``
}

func (this *Hx_Obj_ereg) Hx_Field_matchedRight() string {
    return ``
}

func (this *Hx_Obj_ereg) Hx_Field_matchedPos() any {
    return nil
}

func (this *Hx_Obj_ereg) Hx_Field_matchSub(s string, pos int, _hx_reserved_len int) bool {
    return false
}

func (this *Hx_Obj_ereg) Hx_Field_matchedNum() int {
    return 0
}

func (this *Hx_Obj_ereg) Hx_Field_split(s string) *[]string {
    return nil
}

func (this *Hx_Obj_ereg) Hx_Field_replace(s string, by string) string {
    return ``
}

func (this *Hx_Obj_ereg) Hx_Field_map(s string, f func(*Hx_Obj_ereg) string) string {
    return ``
}

func Hx_Field_ereg_escape(s string) string {
    return ``
}
