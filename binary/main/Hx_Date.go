package main

import "time"

var Hx_Obj_date_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
    "Date",
)

type Hx_Obj_VTable_date interface {
    Hx_Field_toString() string
    Hx_Field_getUTCSeconds() int
    Hx_Field_getUTCMonth() int
    Hx_Field_getUTCMinutes() int
    Hx_Field_getUTCHours() int
    Hx_Field_getUTCFullYear() int
    Hx_Field_getUTCDay() int
    Hx_Field_getUTCDate() int
    Hx_Field_getTimezoneOffset() int
    Hx_Field_getTime() float64
    Hx_Field_getSeconds() int
    Hx_Field_getMonth() int
    Hx_Field_getMinutes() int
    Hx_Field_getHours() int
    Hx_Field_getFullYear() int
    Hx_Field_getDay() int
    Hx_Field_getDate() int
    Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_date struct {
    VTable Hx_Obj_VTable_date
    Hx_Field_t time.Time
}

func Hx_Obj_date_CreateEmptyInstance() *Hx_Obj_date {
    obj := &Hx_Obj_date{}
    obj.VTable = obj
    return obj
}

func Hx_Obj_date_CreateInstance(year int, month int, day int, hour int, min int, sec int) *Hx_Obj_date {
    obj := Hx_Obj_date_CreateEmptyInstance()
    obj.Hx_New(year, month, day, hour, min, sec)
    return obj
}

func (this *Hx_Obj_date) Hx_New(year int, month int, day int, hour int, min int, sec int) {
    var _hx_tmp_0 int = year; _ = _hx_tmp_0
    var _hx_tmp_1 time.Month = ((time.Month)(int64(month))); _ = _hx_tmp_1
    var _hx_tmp_2 int = day; _ = _hx_tmp_2
    var _hx_tmp_3 int = hour; _ = _hx_tmp_3
    var _hx_tmp_4 int = min; _ = _hx_tmp_4
    var _hx_tmp_5 int = sec; _ = _hx_tmp_5
    this.Hx_Field_t = time.Date(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, 0, time.Local)
}

func (this *Hx_Obj_date) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
    return Hx_Obj_date_RTTI
}

func (this *Hx_Obj_date) Hx_Field_getTime() float64 {
    return ((float64)(this.Hx_Field_t.Unix()))
}

func (this *Hx_Obj_date) Hx_Field_getFullYear() int {
    return this.Hx_Field_t.Year()
}

func (this *Hx_Obj_date) Hx_Field_getMonth() int {
    return ((int)(this.Hx_Field_t.Month()))
}

func (this *Hx_Obj_date) Hx_Field_getDate() int {
    return ((int)(this.Hx_Field_t.Unix()))
}

func (this *Hx_Obj_date) Hx_Field_getHours() int {
    return this.Hx_Field_t.Hour()
}

func (this *Hx_Obj_date) Hx_Field_getMinutes() int {
    return this.Hx_Field_t.Minute()
}

func (this *Hx_Obj_date) Hx_Field_getSeconds() int {
    return this.Hx_Field_t.Second()
}

func (this *Hx_Obj_date) Hx_Field_getDay() int {
    return this.Hx_Field_t.Day()
}

func (this *Hx_Obj_date) Hx_Field_getUTCFullYear() int {
    return this.Hx_Field_t.UTC().Year()
}

func (this *Hx_Obj_date) Hx_Field_getUTCMonth() int {
    return ((int)(this.Hx_Field_t.UTC().Month()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCDate() int {
    return ((int)(this.Hx_Field_t.UTC().Unix()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCHours() int {
    return this.Hx_Field_t.UTC().Hour()
}

func (this *Hx_Obj_date) Hx_Field_getUTCMinutes() int {
    return this.Hx_Field_t.UTC().Minute()
}

func (this *Hx_Obj_date) Hx_Field_getUTCSeconds() int {
    return this.Hx_Field_t.UTC().Second()
}

func (this *Hx_Obj_date) Hx_Field_getUTCDay() int {
    return this.Hx_Field_t.UTC().Day()
}

func (this *Hx_Obj_date) Hx_Field_getTimezoneOffset() int {
    var hx_tuple_1 struct { Name string; Offset int }; _ = hx_tuple_1
    hx_tuple_1.Name, hx_tuple_1.Offset = this.Hx_Field_t.Zone()
    return ((struct { Name string; Offset int })(hx_tuple_1)).Offset
}

func (this *Hx_Obj_date) Hx_Field_toString() string {
    return this.Hx_Field_t.String()
}

func Hx_Field_date_now() *Hx_Obj_date {
    var d *Hx_Obj_date = Hx_Field_date_createEmpty(); _ = d
    d.Hx_Field_t = time.Now()
    return d
}

func Hx_Field_date_createEmpty() *Hx_Obj_date {
    return Hx_Obj_date_CreateInstance(0, 0, 0, 0, 0, 0)
}

func Hx_Field_date_fromInt(t int) *Hx_Obj_date {
    var d *Hx_Obj_date = Hx_Field_date_createEmpty(); _ = d
    var _hx_tmp_0 int64 = int64(t); _ = _hx_tmp_0
    d.Hx_Field_t = time.Unix(_hx_tmp_0, int64(0))
    return d
}

func Hx_Field_date_fromTime(t float64) *Hx_Obj_date {
    var d *Hx_Obj_date = Hx_Field_date_createEmpty(); _ = d
    d.Hx_Field_t = time.UnixMilli(((int64)(Hx_Field_haxe_int64helper_fromFloat(t))))
    return d
}

func Hx_Field_date_fromString(s string) *Hx_Obj_date {
    var d *Hx_Obj_date = Hx_Field_date_createEmpty(); _ = d
    var hx_result_4 struct { Error error; Result time.Time }; _ = hx_result_4
    hx_result_4.Result, hx_result_4.Error = time.Parse(time.DateTime, s)
    var res struct { Error error; Result time.Time } = hx_result_4; _ = res
    var tmp bool; _ = tmp
    if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
        tmp = false
    } else {
        tmp = true
    }

    if (tmp) {
        var tmp_tmp_1 time.Time; _ = tmp_tmp_1
        if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
            var e error = ((struct { Error error; Result time.Time })(res)).Error; _ = e
            panic(e)
        } else {
            var r time.Time = ((struct { Error error; Result time.Time })(res)).Result; _ = r
            tmp_tmp_1 = r
        }
    
        d.Hx_Field_t = tmp_tmp_1
        return d
    }

    var hx_result_5 struct { Error error; Result time.Time }; _ = hx_result_5
    hx_result_5.Result, hx_result_5.Error = time.Parse(time.DateOnly, s)
    res = hx_result_5
    var tmp1 bool; _ = tmp1
    if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
        tmp1 = false
    } else {
        tmp1 = true
    }

    if (tmp1) {
        var tmp_tmp_1 time.Time; _ = tmp_tmp_1
        if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
            var e error = ((struct { Error error; Result time.Time })(res)).Error; _ = e
            panic(e)
        } else {
            var r time.Time = ((struct { Error error; Result time.Time })(res)).Result; _ = r
            tmp_tmp_1 = r
        }
    
        d.Hx_Field_t = tmp_tmp_1
        return d
    }

    var hx_result_6 struct { Error error; Result time.Time }; _ = hx_result_6
    hx_result_6.Result, hx_result_6.Error = time.Parse(time.TimeOnly, s)
    res = hx_result_6
    var tmp2 bool; _ = tmp2
    if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
        tmp2 = false
    } else {
        tmp2 = true
    }

    if (tmp2) {
        var tmp_tmp_1 time.Time; _ = tmp_tmp_1
        if ((((struct { Error error; Result time.Time })(res)).Error != nil)) {
            var e error = ((struct { Error error; Result time.Time })(res)).Error; _ = e
            panic(e)
        } else {
            var r time.Time = ((struct { Error error; Result time.Time })(res)).Result; _ = r
            tmp_tmp_1 = r
        }
    
        d.Hx_Field_t = tmp_tmp_1
        return d
    }

    panic(("Invalid date format : " + s))
}
