package main

import "unicode/utf8"

var Hx_Obj_stringtools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "StringTools",
)

type Hx_Obj_VTable_stringtools interface {
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_stringtools struct {
    VTable Hx_Obj_VTable_stringtools
}

func Hx_Obj_stringtools_CreateEmptyInstance() *Hx_Obj_stringtools {
    obj := &Hx_Obj_stringtools{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_stringtools_CreateInstance() *Hx_Obj_stringtools {
    obj := Hx_Obj_stringtools_CreateEmptyInstance()
    return obj
}

func (this *Hx_Obj_stringtools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_stringtools_RTTI
}

func Hx_Field_stringtools_urlEncode(s string) string {
    return ``
}

func Hx_Field_stringtools_urlDecode(s string) string {
    return ``
}

func Hx_Field_stringtools_htmlEscape(s string, quotes bool) string {
    var buf_b string = ""; _ = buf_b
    {
        var _g_offset int = 0; _ = _g_offset
        var _g_s string = s; _ = _g_s
        for  {
            var _hx_tmp_0 int = _g_offset; _ = _hx_tmp_0
            if (!((_hx_tmp_0 < utf8.RuneCountInString(_g_s)))) {
                break
            }
        
            var _hx_tmp_1 any = Hx_Field_go_haxe_hxdynamic_getField(_g_s, "cca"); _ = _hx_tmp_1
            var _hx_tmp_2 int = _g_offset; _ = _hx_tmp_2
            _g_offset = (_g_offset + 1)
            var code int = ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, &([]any{ ((any)(_hx_tmp_2)) }))))); _ = code
            var _hx_tmp_3 string = buf_b; _ = _hx_tmp_3
            switch (code) {
                case 34:
                    if (quotes) {
                        buf_b = (buf_b + "&quot;")
                    } else {
                        var _hx_tmp_3 string = buf_b; _ = _hx_tmp_3
                        buf_b = (_hx_tmp_3 + Hx_Field_go_haxe_hxstring_fromCharCode(code))
                    }
            
                case 38:
                    buf_b = (buf_b + "&amp;")
            
                case 39:
                    if (quotes) {
                        buf_b = (buf_b + "&#039;")
                    } else {
                        var _hx_tmp_3 string = buf_b; _ = _hx_tmp_3
                        buf_b = (_hx_tmp_3 + Hx_Field_go_haxe_hxstring_fromCharCode(code))
                    }
            
                case 60:
                    buf_b = (buf_b + "&lt;")
            
                case 62:
                    buf_b = (buf_b + "&gt;")
            
                default: 
                    buf_b = (_hx_tmp_3 + Hx_Field_go_haxe_hxstring_fromCharCode(code))
            }
        }
    }

    return buf_b
}

func Hx_Field_stringtools_htmlUnescape(s string) string {
    var data []string = (*Hx_Field_go_haxe_hxstring_split(s, "&gt;")); _ = data
    var length int = len(data); _ = length
    var sep string = ">"; _ = sep
    var _this string; _ = _this
    var _hx_tmp_0 int = length; _ = _hx_tmp_0
    if ((_hx_tmp_0 == int(0))) {
        _this = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length)) {
            var _hx_tmp_1 string = result; _ = _hx_tmp_1
            result = (_hx_tmp_1 + Hx_Field_std_string(data[((int)(i))]))
            var _hx_tmp_2 int = i; _ = _hx_tmp_2
            var _hx_tmp_3 int = length; _ = _hx_tmp_3
            if ((_hx_tmp_2 < (_hx_tmp_3 - int(1)))) {
                result = (result + sep)
            }
        
            i = (i + ((int)(1)))
        }
    
        _this = result
    }

    var data1 []string = (*Hx_Field_go_haxe_hxstring_split(_this, "&lt;")); _ = data1
    var length1 int = len(data1); _ = length1
    var sep1 string = "<"; _ = sep1
    var _this1 string; _ = _this1
    var _hx_tmp_1 int = length1; _ = _hx_tmp_1
    if ((_hx_tmp_1 == int(0))) {
        _this1 = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length1)) {
            var _hx_tmp_2 string = result; _ = _hx_tmp_2
            result = (_hx_tmp_2 + Hx_Field_std_string(data1[((int)(i))]))
            var _hx_tmp_3 int = i; _ = _hx_tmp_3
            var _hx_tmp_4 int = length1; _ = _hx_tmp_4
            if ((_hx_tmp_3 < (_hx_tmp_4 - int(1)))) {
                result = (result + sep1)
            }
        
            i = (i + ((int)(1)))
        }
    
        _this1 = result
    }

    var data2 []string = (*Hx_Field_go_haxe_hxstring_split(_this1, "&quot;")); _ = data2
    var length2 int = len(data2); _ = length2
    var sep2 string = "\""; _ = sep2
    var _this2 string; _ = _this2
    var _hx_tmp_2 int = length2; _ = _hx_tmp_2
    if ((_hx_tmp_2 == int(0))) {
        _this2 = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length2)) {
            var _hx_tmp_3 string = result; _ = _hx_tmp_3
            result = (_hx_tmp_3 + Hx_Field_std_string(data2[((int)(i))]))
            var _hx_tmp_4 int = i; _ = _hx_tmp_4
            var _hx_tmp_5 int = length2; _ = _hx_tmp_5
            if ((_hx_tmp_4 < (_hx_tmp_5 - int(1)))) {
                result = (result + sep2)
            }
        
            i = (i + ((int)(1)))
        }
    
        _this2 = result
    }

    var data3 []string = (*Hx_Field_go_haxe_hxstring_split(_this2, "&#039;")); _ = data3
    var length3 int = len(data3); _ = length3
    var sep3 string = "'"; _ = sep3
    var _this3 string; _ = _this3
    var _hx_tmp_3 int = length3; _ = _hx_tmp_3
    if ((_hx_tmp_3 == int(0))) {
        _this3 = ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length3)) {
            var _hx_tmp_4 string = result; _ = _hx_tmp_4
            result = (_hx_tmp_4 + Hx_Field_std_string(data3[((int)(i))]))
            var _hx_tmp_5 int = i; _ = _hx_tmp_5
            var _hx_tmp_6 int = length3; _ = _hx_tmp_6
            if ((_hx_tmp_5 < (_hx_tmp_6 - int(1)))) {
                result = (result + sep3)
            }
        
            i = (i + ((int)(1)))
        }
    
        _this3 = result
    }

    var data4 []string = (*Hx_Field_go_haxe_hxstring_split(_this3, "&amp;")); _ = data4
    var length4 int = len(data4); _ = length4
    var sep4 string = "&"; _ = sep4
    var _hx_tmp_4 int = length4; _ = _hx_tmp_4
    if ((_hx_tmp_4 == int(0))) {
        return ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length4)) {
            var _hx_tmp_5 string = result; _ = _hx_tmp_5
            result = (_hx_tmp_5 + Hx_Field_std_string(data4[((int)(i))]))
            var _hx_tmp_6 int = i; _ = _hx_tmp_6
            var _hx_tmp_7 int = length4; _ = _hx_tmp_7
            if ((_hx_tmp_6 < (_hx_tmp_7 - int(1)))) {
                result = (result + sep4)
            }
        
            i = (i + ((int)(1)))
        }
    
        return result
    }
}

func Hx_Field_stringtools_contains(s string, value string) bool {
    var _hx_tmp_0 string = s; _ = _hx_tmp_0
    var _hx_tmp_1 string = value; _ = _hx_tmp_1
    return (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, _hx_tmp_1, struct { Value int; Valid bool }{}) != -1)
}

func Hx_Field_stringtools_startsWith(s string, start string) bool {
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 int = utf8.RuneCountInString(s); _ = _hx_tmp_1
    if ((_hx_tmp_1 >= utf8.RuneCountInString(start))) {
        var _hx_tmp_2 string = s; _ = _hx_tmp_2
        var _hx_tmp_3 string = start; _ = _hx_tmp_3
        var _hx_tmp_4 string = _hx_tmp_2; _ = _hx_tmp_4
        var _hx_tmp_5 string = _hx_tmp_3; _ = _hx_tmp_5
        _hx_tmp_0 = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_4, _hx_tmp_5, struct { Value int; Valid bool }{ Value: 0, Valid: true }) == 0)
    } else {
        _hx_tmp_0 = false
    }

    return _hx_tmp_0
}

func Hx_Field_stringtools_endsWith(s string, end string) bool {
    var elen int = utf8.RuneCountInString(end); _ = elen
    var slen int = utf8.RuneCountInString(s); _ = slen
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    if ((slen >= elen)) {
        var _hx_tmp_1 string = s; _ = _hx_tmp_1
        var _hx_tmp_2 string = end; _ = _hx_tmp_2
        var _hx_tmp_3 string = _hx_tmp_1; _ = _hx_tmp_3
        var _hx_tmp_4 string = _hx_tmp_2; _ = _hx_tmp_4
        _hx_tmp_0 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, _hx_tmp_4, struct { Value int; Valid bool }{ Value: (slen - elen), Valid: true }) == (slen - elen))
    } else {
        _hx_tmp_0 = false
    }

    return _hx_tmp_0
}

func Hx_Field_stringtools_isSpace(s string, pos int) bool {
    var c struct { Value int; Valid bool } = Hx_Field_go_haxe_hxstring_charCodeAt(s, pos); _ = c
    var _hx_tmp_0 bool; _ = _hx_tmp_0
    var _hx_tmp_1 bool; _ = _hx_tmp_1
    if ((c.Value > 8)) {
        _hx_tmp_1 = (c.Value < 14)
    } else {
        _hx_tmp_1 = false
    }

    if (_hx_tmp_1) {
        _hx_tmp_0 = true
    } else {
        _hx_tmp_0 = (c.Value == 32)
    }

    return _hx_tmp_0
}

func Hx_Field_stringtools_ltrim(s string) string {
    var l int = utf8.RuneCountInString(s); _ = l
    var r int = 0; _ = r
    for  {
        var _hx_tmp_0 bool; _ = _hx_tmp_0
        if ((r < l)) {
            _hx_tmp_0 = Hx_Field_stringtools_isSpace(s, r)
        } else {
            _hx_tmp_0 = false
        }
    
        if (!(_hx_tmp_0)) {
            break
        }
    
        r++
    }

    if ((r > 0)) {
        var _hx_tmp_0 string = s; _ = _hx_tmp_0
        var _hx_tmp_1 int = r; _ = _hx_tmp_1
        return Hx_Field_go_haxe_hxstring_substr(_hx_tmp_0, _hx_tmp_1, struct { Value int; Valid bool }{ Value: (l - r), Valid: true })
    } else {
        return s
    }
}

func Hx_Field_stringtools_rtrim(s string) string {
    var l int = utf8.RuneCountInString(s); _ = l
    var r int = 0; _ = r
    for  {
        var _hx_tmp_0 bool; _ = _hx_tmp_0
        if ((r < l)) {
            _hx_tmp_0 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
        } else {
            _hx_tmp_0 = false
        }
    
        if (!(_hx_tmp_0)) {
            break
        }
    
        r++
    }

    if ((r > 0)) {
        var _hx_tmp_0 string = s; _ = _hx_tmp_0
        return Hx_Field_go_haxe_hxstring_substr(_hx_tmp_0, 0, struct { Value int; Valid bool }{ Value: (l - r), Valid: true })
    } else {
        return s
    }
}

func Hx_Field_stringtools_trim(s string) string {
    var l int = utf8.RuneCountInString(s); _ = l
    var r int = 0; _ = r
    for  {
        var _hx_tmp_0 bool; _ = _hx_tmp_0
        if ((r < l)) {
            _hx_tmp_0 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
        } else {
            _hx_tmp_0 = false
        }
    
        if (!(_hx_tmp_0)) {
            break
        }
    
        r++
    }

    var _hx_tmp_0 string; _ = _hx_tmp_0
    if ((r > 0)) {
        var _hx_tmp_1 string = s; _ = _hx_tmp_1
        var _hx_tmp_2 string = _hx_tmp_1; _ = _hx_tmp_2
        _hx_tmp_0 = Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, struct { Value int; Valid bool }{ Value: (l - r), Valid: true })
    } else {
        _hx_tmp_0 = s
    }

    var s1 string = _hx_tmp_0; _ = s1
    var l1 int = utf8.RuneCountInString(s1); _ = l1
    var r1 int = 0; _ = r1
    for  {
        var _hx_tmp_1 bool; _ = _hx_tmp_1
        if ((r1 < l1)) {
            _hx_tmp_1 = Hx_Field_stringtools_isSpace(s1, r1)
        } else {
            _hx_tmp_1 = false
        }
    
        if (!(_hx_tmp_1)) {
            break
        }
    
        r1++
    }

    if ((r1 > 0)) {
        var _hx_tmp_1 string = s1; _ = _hx_tmp_1
        var _hx_tmp_2 int = r1; _ = _hx_tmp_2
        return Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, struct { Value int; Valid bool }{ Value: (l1 - r1), Valid: true })
    } else {
        return s1
    }
}

func Hx_Field_stringtools_lpad(s string, c string, l int) string {
    if ((utf8.RuneCountInString(c) <= 0)) {
        return s
    }

    var buf_b string = ""; _ = buf_b
    var _hx_tmp_0 int = l; _ = _hx_tmp_0
    l = (_hx_tmp_0 - utf8.RuneCountInString(s))
    for  {
        if (!((utf8.RuneCountInString(buf_b) < l))) {
            break
        }
    
        var _hx_tmp_1 string = buf_b; _ = _hx_tmp_1
        buf_b = (_hx_tmp_1 + Hx_Field_std_string(c))
    }

    var _hx_tmp_1 string = buf_b; _ = _hx_tmp_1
    buf_b = (_hx_tmp_1 + Hx_Field_std_string(s))
    return buf_b
}

func Hx_Field_stringtools_rpad(s string, c string, l int) string {
    if ((utf8.RuneCountInString(c) <= 0)) {
        return s
    }

    var buf_b string = ""; _ = buf_b
    var _hx_tmp_0 string = buf_b; _ = _hx_tmp_0
    buf_b = (_hx_tmp_0 + Hx_Field_std_string(s))
    for  {
        if (!((utf8.RuneCountInString(buf_b) < l))) {
            break
        }
    
        var _hx_tmp_1 string = buf_b; _ = _hx_tmp_1
        buf_b = (_hx_tmp_1 + Hx_Field_std_string(c))
    }

    return buf_b
}

func Hx_Field_stringtools_replace(s string, sub string, by string) string {
    var data []string = (*Hx_Field_go_haxe_hxstring_split(s, sub)); _ = data
    var length int = len(data); _ = length
    var _hx_tmp_0 string; _ = _hx_tmp_0
    if (((by != ``) == false)) {
        _hx_tmp_0 = ","
    } else {
        _hx_tmp_0 = by
    }

    var sep string = _hx_tmp_0; _ = sep
    var _hx_tmp_1 int = length; _ = _hx_tmp_1
    if ((_hx_tmp_1 == int(0))) {
        return ""
    } else {
        var result string = ""; _ = result
        var i int = int(0); _ = i
        for ((i < length)) {
            var _hx_tmp_2 string = result; _ = _hx_tmp_2
            result = (_hx_tmp_2 + Hx_Field_std_string(data[((int)(i))]))
            var _hx_tmp_3 int = i; _ = _hx_tmp_3
            var _hx_tmp_4 int = length; _ = _hx_tmp_4
            if ((_hx_tmp_3 < (_hx_tmp_4 - int(1)))) {
                result = (result + sep)
            }
        
            i = (i + ((int)(1)))
        }
    
        return result
    }
}

func Hx_Field_stringtools_hex(n int, digits struct { Value int; Valid bool }) string {
    var s string = ""; _ = s
    var hexChars string = "0123456789ABCDEF"; _ = hexChars
    for  {
        s = (Hx_Field_go_haxe_hxstring_charAt(hexChars, (n & 15)) + s)
        n = ((int)((((uint32)(n)) >> 4)))
        if (!((n > 0))) {
            break
        }
    }

    if ((digits.Valid != false)) {
        for  {
            var _hx_tmp_0 int = utf8.RuneCountInString(s); _ = _hx_tmp_0
            if (!((_hx_tmp_0 < digits.Value))) {
                break
            }
        
            s = ("0" + s)
        }
    }

    return s
}

func Hx_Field_stringtools_fastCodeAt(s string, index int) int {
    return Hx_Field_go_haxe_hxstring_charCodeAt(s, index).Value
}

func Hx_Field_stringtools_unsafeCodeAt(s string, index int) int {
    var _hx_tmp_0 any = Hx_Field_go_haxe_hxdynamic_getField(s, "cca"); _ = _hx_tmp_0
    return ((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, &([]any{ ((any)(index)) })))))
}

func Hx_Field_stringtools_iterator(s string) *Hx_Obj_haxe_iterators_stringiterator {
    return Hx_Obj_haxe_iterators_stringiterator_CreateInstance(s)
}

func Hx_Field_stringtools_keyValueIterator(s string) *Hx_Obj_haxe_iterators_stringkeyvalueiterator {
    return Hx_Obj_haxe_iterators_stringkeyvalueiterator_CreateInstance(s)
}

func Hx_Field_stringtools_isEof(c int) bool {
    return false
}

func Hx_Field_stringtools_quoteUnixArg(argument string) string {
    if ((argument == "")) {
        return "''"
    } else {
        if (!Hx_Obj_ereg_CreateInstance("[^a-zA-Z0-9_@%+=:,./-]", "").VTable.Hx_Field_match(argument)) {
            return argument
        } else {
            return (("'" + Hx_Field_stringtools_replace(argument, "'", "'\"'\"'")) + "'")
        }
    }
}

var Hx_Field_stringtools_winMetaCharacters *[]int = Hx_Init_Hx_Field_stringtools_winMetaCharacters()
func Hx_Init_Hx_Field_stringtools_winMetaCharacters() *[]int {
    return (Hx_Field_haxe_systools_winMetaCharacters).(*[]int)
}

func Hx_Field_stringtools_quoteWinArg(argument string, escapeMetaCharacters bool) string {
    var argument1 string = argument; _ = argument1
    if (!Hx_Obj_ereg_CreateInstance("^(/)?[^ 	/\\\\\"]+$", "").VTable.Hx_Field_match(argument1)) {
        var result_b string = ""; _ = result_b
        var _hx_tmp_0 bool; _ = _hx_tmp_0
        var _hx_tmp_1 bool; _ = _hx_tmp_1
        var _hx_tmp_2 bool; _ = _hx_tmp_2
        var _hx_tmp_3 string = argument1; _ = _hx_tmp_3
        if ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, " ", struct { Value int; Valid bool }{}) != -1)) {
            _hx_tmp_2 = true
        } else {
            var _hx_tmp_4 string = argument1; _ = _hx_tmp_4
            var _hx_tmp_5 string = _hx_tmp_4; _ = _hx_tmp_5
            _hx_tmp_2 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "	", struct { Value int; Valid bool }{}) != -1)
        }
    
        if (_hx_tmp_2) {
            _hx_tmp_1 = true
        } else {
            _hx_tmp_1 = (argument1 == "")
        }
    
        if (_hx_tmp_1) {
            _hx_tmp_0 = true
        } else {
            var _hx_tmp_4 string = argument1; _ = _hx_tmp_4
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
            var _g1 int = utf8.RuneCountInString(argument1); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_4 int = _g; _ = _hx_tmp_4
                _g = (_g + 1)
                var i int = _hx_tmp_4; _ = i
                {
                    var tmp__g1_1 struct { Value int; Valid bool } = Hx_Field_go_haxe_hxstring_charCodeAt(argument1, i); _ = tmp__g1_1
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
    
        argument1 = result_b
    }

    if (escapeMetaCharacters) {
        var result_b string = ""; _ = result_b
        {
            var _g int = 0; _ = _g
            var _g1 int = utf8.RuneCountInString(argument1); _ = _g1
            for ((_g < _g1)) {
                var _hx_tmp_0 int = _g; _ = _hx_tmp_0
                _g = (_g + 1)
                var i int = _hx_tmp_0; _ = i
                var c struct { Value int; Valid bool } = Hx_Field_go_haxe_hxstring_charCodeAt(argument1, i); _ = c
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
        return argument1
    }
}
