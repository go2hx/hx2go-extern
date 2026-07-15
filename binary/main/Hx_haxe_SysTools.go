package main

import "unicode/utf8"

var Hx_Obj_haxe_systools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "haxe.SysTools",
)

type Hx_Obj_VTable_haxe_systools interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_systools struct {
    VTable Hx_Obj_VTable_haxe_systools
}

func Hx_Obj_haxe_systools_CreateEmptyInstance() *Hx_Obj_haxe_systools {
    obj := &Hx_Obj_haxe_systools{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_haxe_systools_CreateInstance() *Hx_Obj_haxe_systools {
    obj := Hx_Obj_haxe_systools_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_haxe_systools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_haxe_systools_RTTI
}

var Hx_Field_haxe_systools_winMetaCharacters any = Hx_Init_Hx_Field_haxe_systools_winMetaCharacters()
func Hx_Init_Hx_Field_haxe_systools_winMetaCharacters() any {
    return ((any)(&([]int{ 32, 40, 41, 37, 33, 94, 34, 60, 62, 38, 124, 10, 13, 44, 59 })))
}

func Hx_Field_haxe_systools_quoteUnixArg(argument string) string {
    if ((argument == "")) {
        return "''"
    }

    if (!Hx_Obj_ereg_CreateInstance("[^a-zA-Z0-9_@%+=:,./-]", "").VTable.Hx_Field_match(argument)) {
        return argument
    }

    return (("'" + Hx_Field_stringtools_replace(argument, "'", "'\"'\"'")) + "'")
}

func Hx_Field_haxe_systools_quoteWinArg(argument string, escapeMetaCharacters bool) string {
    if (!Hx_Obj_ereg_CreateInstance("^(/)?[^ 	/\\\\\"]+$", "").VTable.Hx_Field_match(argument)) {
        var result_b string = ""; _ = result_b
        var _hx_tmp_0 bool; _ = _hx_tmp_0
        var _hx_tmp_1 bool; _ = _hx_tmp_1
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        var _hx_tmp_3 string = argument; _ = _hx_tmp_3
        if ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, " ", struct { Value int; Valid bool }{}) != -1)) {
            _hx_tmp_2 = true
        } else {
            var _hx_tmp_4 string = argument; _ = _hx_tmp_4
            var _hx_tmp_5 string = _hx_tmp_4; _ = _hx_tmp_5
            _hx_tmp_2 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "	", struct { Value int; Valid bool }{}) != -1)
        }
    
        if (_hx_tmp_2) {
            _hx_tmp_1 = true
        } else {
            _hx_tmp_1 = (argument == "")
        }
    
        if (_hx_tmp_1) {
            _hx_tmp_0 = true
        } else {
            var _hx_tmp_4 string = argument; _ = _hx_tmp_4
            var _hx_tmp_5 string = _hx_tmp_4; _ = _hx_tmp_5
            _hx_tmp_0 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "/", struct { Value int; Valid bool }{}) > 0)
        }
    
        var needquote bool = _hx_tmp_0; _ = needquote
        if (needquote) {
            result_b = (result_b + "\"")
        }
    
        var bs_buf *Hx_Obj_stringbuf = Hx_Obj_stringbuf_CreateInstance(); _ = bs_buf
        {
            var _g int = 0; _ = _g
            var _g1 int = utf8.RuneCountInString(argument); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_4 int = _g; _ = _hx_tmp_4
                _g = (_g + 1)
                var i int = _hx_tmp_4; _ = i
                {
                    var tmp__g1_1 struct { Value int; Valid bool } = Hx_Field_go_haxe_hxstring_charCodeAt(argument, i); _ = tmp__g1_1
                    if ((tmp__g1_1.Valid == false)) {
                        var c struct { Value int; Valid bool } = tmp__g1_1; _ = c
                        {
                            if ((utf8.RuneCountInString(bs_buf.Hx_Field_b) > 0)) {
                                var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
                                result_b = (_hx_tmp_5 + Hx_Field_std_string(bs_buf.Hx_Field_b))
                                bs_buf = Hx_Obj_stringbuf_CreateInstance()
                            }
                        
                            var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
                            result_b = (_hx_tmp_5 + Hx_Field_go_haxe_hxstring_fromCharCode(c.Value))
                        }
                    } else {
                        switch (tmp__g1_1).Value {
                            case 34:
                                {
                                    var bs string = bs_buf.Hx_Field_b; _ = bs
                                    var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
                                    result_b = (_hx_tmp_5 + Hx_Field_std_string(bs))
                                    var _hx_tmp_6 string = result_b; _ = _hx_tmp_6
                                    result_b = (_hx_tmp_6 + Hx_Field_std_string(bs))
                                    bs_buf = Hx_Obj_stringbuf_CreateInstance()
                                    result_b = (result_b + "\\\"")
                                }
                        
                            case 92:
                                bs_buf.Hx_Field_b = (bs_buf.Hx_Field_b + "\\")
                        
                            default: 
                                {
                                    var c struct { Value int; Valid bool } = tmp__g1_1; _ = c
                                    {
                                        if ((utf8.RuneCountInString(bs_buf.Hx_Field_b) > 0)) {
                                            var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
                                            result_b = (_hx_tmp_5 + Hx_Field_std_string(bs_buf.Hx_Field_b))
                                            bs_buf = Hx_Obj_stringbuf_CreateInstance()
                                        }
                                    
                                        var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
                                        result_b = (_hx_tmp_5 + Hx_Field_go_haxe_hxstring_fromCharCode(c.Value))
                                    }
                                }
                        }
                    }
                }
            }
        }
    
        var _hx_tmp_4 string = result_b; _ = _hx_tmp_4
        result_b = (_hx_tmp_4 + Hx_Field_std_string(bs_buf.Hx_Field_b))
        if (needquote) {
            var _hx_tmp_5 string = result_b; _ = _hx_tmp_5
            result_b = (_hx_tmp_5 + Hx_Field_std_string(bs_buf.Hx_Field_b))
            result_b = (result_b + "\"")
        }
    
        argument = result_b
    }

    if (escapeMetaCharacters) {
        var result_b string = ""; _ = result_b
        {
            var _g int = 0; _ = _g
            var _g1 int = utf8.RuneCountInString(argument); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_0 int = _g; _ = _hx_tmp_0
                _g = (_g + 1)
                var i int = _hx_tmp_0; _ = i
                var c struct { Value int; Valid bool } = Hx_Field_go_haxe_hxstring_charCodeAt(argument, i); _ = c
                var fromIndex struct { Value int; Valid bool } = struct { Value int; Valid bool }{}; _ = fromIndex
                var data []int = (*(Hx_Field_haxe_systools_winMetaCharacters).(*[]int)); _ = data
                var length int = len(data); _ = length
                var start int; _ = start
                if ((fromIndex.Valid == false)) {
                    start = int(0)
                } else {
                    if ((fromIndex.Value < 0)) {
                        var _hx_tmp_1 int = length; _ = _hx_tmp_1
                        var idx int = (_hx_tmp_1 + int(fromIndex.Value)); _ = idx
                        var _hx_tmp_2 int; _ = _hx_tmp_2
                        var _hx_tmp_3 int = idx; _ = _hx_tmp_3
                        if ((_hx_tmp_3 < int(0))) {
                            _hx_tmp_2 = int(0)
                        } else {
                            _hx_tmp_2 = idx
                        }
                    
                        start = _hx_tmp_2
                    } else {
                        start = int(fromIndex.Value)
                    }
                }
            
                var this1 int; _ = this1
                if ((start >= length)) {
                    this1 = int(-1)
                } else {
                    var tmp_i_1 int = start; _ = tmp_i_1
                    var res int = int(-1); _ = res
                    for ((tmp_i_1 < length)) {
                        var _hx_tmp_1 int = data[((int)(tmp_i_1))]; _ = _hx_tmp_1
                        if ((_hx_tmp_1 == c.Value)) {
                            res = tmp_i_1
                            break
                        }
                    
                        tmp_i_1 = (tmp_i_1 + ((int)(1)))
                    }
                
                    this1 = res
                }
            
                if ((((int)(((int)(this1)))) >= 0)) {
                    var _hx_tmp_1 string = result_b; _ = _hx_tmp_1
                    result_b = (_hx_tmp_1 + Hx_Field_go_haxe_hxstring_fromCharCode(94))
                }
            
                var _hx_tmp_1 string = result_b; _ = _hx_tmp_1
                result_b = (_hx_tmp_1 + Hx_Field_go_haxe_hxstring_fromCharCode(c.Value))
            }
        }
    
        return result_b
    } else {
        return argument
    }
}
