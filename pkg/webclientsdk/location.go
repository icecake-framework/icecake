package webclientsdk

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Location
type Location struct {
	jsValue js.Value
}

// LocationFromJS is casting a js.Value into Location.
func LocationFromJS(value js.Value) *Location {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Location{}
	ret.jsValue = value
	return ret
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *Location) Href() string {
	var ret string
	value := _this.jsValue.Get("href")
	ret = (value).String()
	return ret
}

// ToString is an alias for Href.
func (_this *Location) ToString() string {
	return _this.Href()
}

// SetHref setting attribute 'href' with
// type string (idl: USVString).
func (_this *Location) SetHref(value string) {
	input := value
	_this.jsValue.Set("href", input)
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *Location) Origin() string {
	var ret string
	value := _this.jsValue.Get("origin")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: USVString).
func (_this *Location) Protocol() string {
	var ret string
	value := _this.jsValue.Get("protocol")
	ret = (value).String()
	return ret
}

// SetProtocol setting attribute 'protocol' with
// type string (idl: USVString).
func (_this *Location) SetProtocol(value string) {
	input := value
	_this.jsValue.Set("protocol", input)
}

// Host returning attribute 'host' with
// type string (idl: USVString).
func (_this *Location) Host() string {
	var ret string
	value := _this.jsValue.Get("host")
	ret = (value).String()
	return ret
}

// SetHost setting attribute 'host' with
// type string (idl: USVString).
func (_this *Location) SetHost(value string) {
	input := value
	_this.jsValue.Set("host", input)
}

// Hostname returning attribute 'hostname' with
// type string (idl: USVString).
func (_this *Location) Hostname() string {
	var ret string
	value := _this.jsValue.Get("hostname")
	ret = (value).String()
	return ret
}

// SetHostname setting attribute 'hostname' with
// type string (idl: USVString).
func (_this *Location) SetHostname(value string) {
	input := value
	_this.jsValue.Set("hostname", input)
}

// Port returning attribute 'port' with
// type string (idl: USVString).
func (_this *Location) Port() string {
	var ret string
	value := _this.jsValue.Get("port")
	ret = (value).String()
	return ret
}

// SetPort setting attribute 'port' with
// type string (idl: USVString).
func (_this *Location) SetPort(value string) {
	input := value
	_this.jsValue.Set("port", input)
}

// Pathname returning attribute 'pathname' with
// type string (idl: USVString).
func (_this *Location) Pathname() string {
	var ret string
	value := _this.jsValue.Get("pathname")
	ret = (value).String()
	return ret
}

// SetPathname setting attribute 'pathname' with
// type string (idl: USVString).
func (_this *Location) SetPathname(value string) {
	input := value
	_this.jsValue.Set("pathname", input)
}

// Search returning attribute 'search' with
// type string (idl: USVString).
func (_this *Location) Search() string {
	var ret string
	value := _this.jsValue.Get("search")
	ret = (value).String()
	return ret
}

// SetSearch setting attribute 'search' with
// type string (idl: USVString).
func (_this *Location) SetSearch(value string) {
	input := value
	_this.jsValue.Set("search", input)
}

// Hash returning attribute 'hash' with
// type string (idl: USVString).
func (_this *Location) Hash() string {
	var ret string
	value := _this.jsValue.Get("hash")
	ret = (value).String()
	return ret
}

// SetHash setting attribute 'hash' with
// type string (idl: USVString).
func (_this *Location) SetHash(value string) {
	input := value
	_this.jsValue.Set("hash", input)
}

func (_this *Location) Assign(url string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_this.jsValue.Call("assign", _args[0:_end]...)
}

func (_this *Location) Replace(url string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_this.jsValue.Call("replace", _args[0:_end]...)
}

func (_this *Location) Reload() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("reload", _args[0:_end]...)
}
