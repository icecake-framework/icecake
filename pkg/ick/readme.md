# html component memo
? Tokeniser
? HtmlElement


Attribute types:
* Id -> id="uniqueid"
* Classes -> class="c1 c2 c3"
* Style -> style="s1=value1;s2=value2;"
* Boolean -> set or unset aka Hidden
* Key/Value -> key={"}value{"}

2 ways to set an attribute :
a) parsing an escaped string, with different rules according to the attribute type
b) setting
    - Id ==> overwrite existing Key/Value, or create it
    - Classes -> append it to the value if the value does not contains it, otherwise does nothing
    - Style -> append it to the value if the value does not contains it, otherwise update the value
    - Boolean -> overwrite existing Key with an empty value
    - Key/Value -> update


Attributes
is a map (or a slice ?) of attribute of any type


Node
* Type
* child nodes

Text Node
* escaped string

Element Node
* tagname
* autoclosing
* attributes
* body 

Html Snippet Node
* ickname
* data properties (exposed)
* private properties
* default rendering tagname
* default rendering attributes
* customized attributes
* Template

