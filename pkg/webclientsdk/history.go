package webclientsdk

import "syscall/js"

/*********************************************************************************
* History
 */

type ScrollRestoration int

const (
	AutoScrollRestoration ScrollRestoration = iota
	ManualScrollRestoration
)

var scrollRestorationToWasmTable = []string{
	"auto", "manual",
}

var scrollRestorationFromWasmTable = map[string]ScrollRestoration{
	"auto": AutoScrollRestoration, "manual": ManualScrollRestoration,
}

// JSValue is converting this enum into a javascript object
func (this *ScrollRestoration) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollRestoration) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollRestorationToWasmTable) {
		return scrollRestorationToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollRestorationFromJS is converting a javascript value into
// a ScrollRestoration enum value.
func ScrollRestorationFromJS(value js.Value) ScrollRestoration {
	key := value.String()
	conv, ok := scrollRestorationFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

/*********************************************************************************
* History
 */

// https://developer.mozilla.org/en-US/docs/Web/API/History
type History struct {
	jsValue js.Value
}

// HistoryFromJS is casting a js.Value into History.
func HistoryFromJS(value js.Value) *History {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &History{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *History) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

// ScrollRestoration returning attribute 'scrollRestoration' with
// type ScrollRestoration (idl: ScrollRestoration).
func (_this *History) ScrollRestoration() ScrollRestoration {
	var ret ScrollRestoration
	value := _this.jsValue.Get("scrollRestoration")
	ret = ScrollRestorationFromJS(value)
	return ret
}

// SetScrollRestoration setting attribute 'scrollRestoration' with
// type ScrollRestoration (idl: ScrollRestoration).
func (_this *History) SetScrollRestoration(value ScrollRestoration) {
	input := value.JSValue()
	_this.jsValue.Set("scrollRestoration", input)
}

// State returning attribute 'state' with
// type Any (idl: any).
func (_this *History) State() js.Value {
	var ret js.Value
	value := _this.jsValue.Get("state")
	ret = value
	return ret
}

func (_this *History) Go(delta *int) {
	var (
		_args [1]interface{}
		_end  int
	)
	if delta != nil {

		var _p0 interface{}
		if delta != nil {
			_p0 = *(delta)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	_this.jsValue.Call("go", _args[0:_end]...)
	return
}

func (_this *History) Back() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("back", _args[0:_end]...)
	return
}

func (_this *History) Forward() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("forward", _args[0:_end]...)
	return
}

func (_this *History) PushState(data interface{}, title string, url *string) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_p1 := title
	_args[1] = _p1
	_end++
	if url != nil {

		var _p2 interface{}
		if url != nil {
			_p2 = *(url)
		} else {
			_p2 = nil
		}
		_args[2] = _p2
		_end++
	}
	_this.jsValue.Call("pushState", _args[0:_end]...)
	return
}

func (_this *History) ReplaceState(data interface{}, title string, url *string) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_p1 := title
	_args[1] = _p1
	_end++
	if url != nil {

		var _p2 interface{}
		if url != nil {
			_p2 = *(url)
		} else {
			_p2 = nil
		}
		_args[2] = _p2
		_end++
	}
	_this.jsValue.Call("replaceState", _args[0:_end]...)
	return
}
