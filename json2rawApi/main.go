package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"reflect"
	"strings"

	"bytes"
	"fmt"
	"go/format"
	"path/filepath"
)

var (
	enableComment bool
	outDir        string
	doFormat      bool
)

var (
	globalTypeNames = make(map[string]struct{})
)

type Context struct {
	w             *bytes.Buffer
	compoundTypes map[string]*Property
	consts        map[string][]*PossibleValue
	// targeting/toplevel block/base
	base *Base
	// global type names
}

func newContext(b *Base) (w *Context, err error) {
	w = new(Context)
	w.base = b
	w.w = bytes.NewBuffer(nil)
	w.compoundTypes = make(map[string]*Property)
	w.consts = make(map[string][]*PossibleValue)
	return w, nil
}

func getOutputFileName(baseFileName string) string {
	opath := outDir + "/" + baseFileName + ".go"
	_, err := os.Stat(opath)
	if err != nil && os.IsNotExist(err) {
		return opath
	}
	for i := 2; ; i++ {
		opath = outDir + "/" + baseFileName + fmt.Sprint(i) + ".go"
		_, err = os.Stat(opath)
		if err != nil && os.IsNotExist(err) {
			break
		}
	}
	return opath
}

func (w *Context) adjustImport(b *Block) {
	if len(b.Properties)+len(b.Methods)+len(b.InstanceProperties)+len(b.InstanceMethods) > 0 {
	}
	buf := bytes.NewBuffer(nil)
	src := w.w.Bytes()
	fmt.Fprintf(buf, "package electron\n")
	if b.isEventEmitter() {
		fmt.Fprintf(buf, "import \"github.com/oskca/gopherjs-nodejs/events\"\n")
	}
	if bytes.Contains(src, []byte("js.Object")) {
		fmt.Fprintf(buf, "\nimport \"github.com/gopherjs/gopherjs/js\"\n")
	}
	buf.Write(src)
	w.w = buf
}

func (w *Context) formatCode() (err error) {
	src := w.w.Bytes()
	// format
	if doFormat {
		src, err = format.Source(src)
		if err != nil {
			return
		}
	}
	w.w = bytes.NewBuffer(src)
	return nil
}

func (w *Context) OutputToFile() error {
	// flush compoundTypes
	opath := getOutputFileName("raw_" + strings.ToLower(w.base.Name))
	ow, err := os.Create(opath)
	if err != nil {
		return err
	}
	// flush and close file
	io.Copy(ow, w.w)
	if err != nil {
		return err
	}
	return ow.Close()
}

func (c *Context) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

func (c *Context) newType(p *Property, parent *Base) string {
	tname := parent.goSym() + p.goSym()
	if c.base != nil { // prefix module/class name
		tname = c.base.goSym() + tname
	}
	// avoid duplicated names
	if _, alreadyExist := globalTypeNames[tname]; alreadyExist {
		for i := 2; ; i++ {
			tmp := fmt.Sprintf("%s%d", tname, i)
			if _, ok := globalTypeNames[tmp]; !ok {
				tname = tmp
				break
			}
		}
	}
	// put into global names
	globalTypeNames[tname] = struct{}{}
	c.compoundTypes[tname] = p
	//
	if p.isObject() || p.isStructure() {
		return "*" + tname
	}
	return tname
}

func (c *Context) newConst(p *Property, parent *Base) string {
	tname := parent.goSym() + p.goSym()
	if c.base != nil { // prefix module/class name
		tname = c.base.goSym() + tname
	}
	// avoid duplicated names
	if _, alreadyExist := globalTypeNames[tname]; alreadyExist {
		for i := 1; ; i++ {
			tmp := fmt.Sprintf("%s%d", tname, i)
			if _, ok := globalTypeNames[tmp]; !ok {
				tname = tmp
				break
			}
		}
	}
	// put into global names
	globalTypeNames[tname] = struct{}{}
	c.consts[tname] = p.PossibleValues
	return tname
}

func (c *Context) sub() *Context {
	nc := new(Context)
	nc.w = c.w
	nc.base = c.base
	nc.compoundTypes = make(map[string]*Property)
	nc.consts = make(map[string][]*PossibleValue)
	return nc
}

func (c *Context) doNewCompound(tname string, p *Property) {
	nc := c.sub()
	defer func() {
		// recursive
		if len(nc.compoundTypes) > 0 || len(nc.consts) > 0 {
			nc.declNewTypes()
		}
	}()
	if p.isFunction() {
		fmt.Fprintf(nc, "\ntype %s func(", tname)
		declSlice(p.Parameters, nc, p.Base, ",")
		fmt.Fprintf(nc, ")")
		return
	}
	fmt.Fprintf(nc, "\ntype %s struct {\n\t", tname)
	fmt.Fprintf(nc, "*js.Object\n\t")
	declSlice(p.Properties, nc, p.Base)
	fmt.Fprintf(nc, "}\n")
}

func (c *Context) declNewTypes() {
	// obj/func/structure etc
	for tname, p := range c.compoundTypes {
		c.doNewCompound(tname, p)
	}
	// consts
	for tname, vals := range c.consts {
		fmt.Fprintf(c, "\ntype %s string\n", tname)
		fmt.Fprintf(c, "\n// consts \nconst (")
		for _, val := range vals {
			fmt.Fprintf(c, "\n\t%s%s %s = \"%s\"",
				tname,
				goSym(val.Value),
				tname,
				val.Value,
			)
		}
		fmt.Fprintf(c, "\n)")
	}
}

type decler interface {
	decl(w *Context, parent *Base)
	annotate(w *Context, parent *Base)
	comment(w io.Writer)
}

func declSlice(ar interface{}, w *Context, parent *Base, seperators ...string) {
	// seperator
	sep := "\n\t"
	if len(seperators) > 0 {
		sep = seperators[0]
	}
	// arrange
	v := reflect.ValueOf(ar)
	if v.IsNil() {
		return
	}
	var t []decler
	ret := reflect.MakeSlice(reflect.TypeOf(t), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret = reflect.Append(ret, v.Index(i))
	}
	// declers to string
	ds := ret.Interface().([]decler)
	for i := 0; i < len(ds); i++ {
		if enableComment {
			ds[i].comment(w)
		}
		ds[i].decl(w, parent)
		ds[i].annotate(w, parent)
		fmt.Fprintf(w, sep)
	}
}

func process(fpath string) error {
	r, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer r.Close()
	// parse
	var a ApiFile
	dec := json.NewDecoder(r)
	err = dec.Decode(&a)
	if err != nil {
		return err
	}
	log.Println("Processing api file:", fpath)
	if err = a.decl(); err != nil {
		return err
	}
	log.Println("Done with", len(a), "modules.")
	return nil
}

func main() {
	// mkdir
	err := os.MkdirAll(outDir, 0777)
	if err != nil {
		log.Fatalln("error cleaning", outDir, ":", err.Error())
	}
	// clean all raw_* files
	filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(path, "raw_") {
			os.Remove(path)
		}
		return nil
	})
	// begin process
	if err = process(flag.Arg(0)); err != nil {
		log.Println(err.Error())
	}
}

func init() {
	flag.BoolVar(&enableComment, "c", false, "generate comment")
	flag.BoolVar(&doFormat, "f", true, "format the output code")
	flag.StringVar(&outDir, "o", "rawapi", "output directory for raw api")
	flag.Parse()
}
