package ickcore

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/lolorenzo777/verbose"
)

type RMetaProvider interface {

	// Meta returns a reference to render meta data
	RMeta() *RMetaData
}

type ComposerMap map[string]RMetaProvider

// RMetaData is rendering metadata for a single HTMLContentComposer
type RMetaData struct {
	Deep      int           // deepness of the HTMLContentComposer
	VirtualId string        // virtual id allocated to an HTMLContentComposer, always one
	Id        string        // the id allocated to the HTMLContentComposer if any
	Parent    RMetaProvider // optional parent, may be an orphan
	IsRender  bool          // Indicates the HTMLContentComposer has been rendered at least once
	IsMounted bool          // Indicates the HTMLContentComposer has been mounted
	RError    error         // rendering error if any

	childs ComposerMap // embedded child content composer
}

func (rmeta *RMetaData) RMeta() *RMetaData {
	return rmeta
}

// GenerateVirtualId generates a unique id for every rendering composer.
// The composer may not have a TagBuilder, so the Id is not necessarly the id attribute of the composer. The generated id pattern is:
//
//	`{{{parentid|orphan}.[cmpname.index]}|[cmpid]}`
//
// rules:
//   - if the rendering has a rendering parent, the virtual id starts with the parent's virtualid otherwise it's "orphan" followed by the component name
//   - if the rendering does not have a rendering parent, the virtual id is the given cmpid if not empty, otherwise it's "orphan"
//   - the dot "-" seperator is added followed by the cmpname if any
//
// - the cmpname is added
func (rmeta *RMetaData) GenerateVirtualId(cmp RMetaProvider) string {
	prefix := "orphan"
	if rmeta.Parent != nil {
		if pvid := rmeta.Parent.RMeta().VirtualId; pvid != "" {
			prefix = pvid
		}
	}
	prefix += "."
	toporphan := strings.HasPrefix(prefix, "orphan.")

	cmpname := reflect.TypeOf(cmp).Elem().Name()
	cmpname = strings.ToLower(cmpname)

	body := cmpname
	cmpid := cmp.RMeta().Id
	if cmpid != "" {
		body = cmpid
		if toporphan {
			toporphan = false
			prefix = ""
		}
	}

	index := 0
	if rmeta.Parent != nil {
		index = len(rmeta.Parent.RMeta().Embedded())
	} else {
		index, _ = GetUniqueId(cmpname)
	}

	suffix := ""
	if cmpid == "" || toporphan {
		suffix = strconv.Itoa(index)
	}

	rmeta.VirtualId = prefix + body + suffix
	return rmeta.VirtualId
}

// Embed adds child to the map of embedded components.
// If a child with the same key has already been embedded it's replaced and a warning is raised in verbose mode.
// The key is the id of the html element if any otherwise it's its virtual id.
func (rmeta *RMetaData) Embed(child RMetaProvider) {
	if rmeta.childs == nil {
		rmeta.childs = make(ComposerMap, 1)
	}
	key := child.RMeta().Id
	if key == "" {
		key = child.RMeta().VirtualId
	}
	if _, f := rmeta.childs[key]; f {
		verbose.Println(verbose.WARNING, "Embed: duplicate child id:%q for parent id:%q", key, rmeta.VirtualId)
	}
	rmeta.childs[key] = child
	child.RMeta().Parent = rmeta
	// verbose.Debug("embedded (%v) %q", reflect.TypeOf(subcmp).String(), id)
}

// Embedded returns the map of embedded components, keyed by their id.
func (rmeta RMetaData) Embedded() ComposerMap {
	if rmeta.childs == nil {
		rmeta.childs = make(ComposerMap, 0)
	}
	return rmeta.childs
}
