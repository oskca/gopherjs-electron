package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"strings"

	"os"

	"github.com/PuerkitoBio/goquery"
)

type Base struct {
	Name        string      `json:"name,omitempty"`
	RawType     interface{} `json:"type,omitempty"`
	Description string      `json:"description,omitempty"`

	Platforms []string `json:"platforms,omitempty"`
	Process   struct {
		Main     bool `json:"main,omitempty"`
		Renderer bool `json:"renderer,omitempty"`
	} `json:"process,omitempty"`

	Required bool `json:"required,omitempty"`

	Version    string `json:"version,omitempty"`
	RepoURL    string `json:"repoUrl,omitempty"`
	WebsiteURL string `json:"websiteUrl,omitempty"`
	Slug       string `json:"slug,omitempty"`
}

func (b *Base) Type() string {
	// default type set to string
	if b.RawType == nil {
		return ""
	}
	// v := reflect.Indirect(reflect.ValueOf(b.RawType))
	// if v.IsNil() {
	// 	return ""
	// }
	// t := reflect.TypeOf(b.RawType)
	if v, ok := b.RawType.(string); ok {
		return v
	}
	if v, ok := b.RawType.([]interface{}); ok {
		return v[0].(string)
	}
	return ""
}

var replacer = strings.NewReplacer(
	"`", " ",
	"\"", " ",
	".", " ",
	"_", " ",
	"-", " ",
	" ", " ",
	"\t", " ",
	"(", " ",
	")", " ",
	// name conversions
	"url", "URL",
	"Url", "URL",
	// ",", "",
)

func goSym(name string) string {
	name = replacer.Replace(name)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return name
}

func (b *Base) goSym() string {
	name := goSym(b.Name)
	if b.isModule() {
		return name + "Module"
	}
	return name
}

func text(desc string) string {
	desc = strings.Replace(desc, "\n", "\n//", -1)
	desc = html.UnescapeString(desc)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(desc))
	if err != nil {
		println(err.Error())
		return desc
	}
	return doc.Text()
}

func (b *Base) comment(w io.Writer) {
	if b.Description != "" {
		if b.Version != "" {
			fmt.Fprintf(w, "// %s version@%s \n//\n// %s\n", b.goSym(), b.Version, text(b.Description))
		} else {
			fmt.Fprintf(w, "// %s\n", text(b.Description))
		}
	} else {
		if b.isStructure() {
			fmt.Fprintf(w, "// %s a Structure\n", b.goSym())
		}
	}
}

func (b *Base) isModule() bool {
	return b.Type() == "Module"
}

func (b *Base) isClass() bool {
	return b.Type() == "Class"
}

func (b *Base) isStructure() bool {
	return b.Type() == "Structure"
}

func (b *Base) isObject() bool {
	return b.Type() == "Object"
}

func (b *Base) isFunction() bool {
	return b.Type() == "Function"
}

func (b *Base) isBasic() bool {
	return !(b.isModule() ||
		b.isClass() ||
		b.isStructure() ||
		b.isFunction() ||
		b.isObject())
}

func basicType(typ string) string {
	switch typ {
	case "", "String":
		return "string"
	case "Integer", "INTEGER":
		return "int64"
	case "Number", "NUMBER", "Double", "DOUBLE", "Float", "FLOAT":
		return "float64"
	case "Boolean", "BOOLEAN":
		return "bool"
	}
	return "*js.Object"
}

func (b *Base) decl(w *Context, parent *Base) {
	fmt.Fprintf(w, "%s %s",
		b.goSym(),
		basicType(b.Type()),
	)
}

func (b *Base) annotate(w *Context, parent *Base) {
	if parent == nil || parent.isFunction() {
		return
	}
	fmt.Fprintf(w, " `js:\"%s\"` ", b.Name)
}

type Property struct {
	*Base
	Properties     []*Property      `json:"properties,omitempty"`     // object or structure
	Parameters     []*Property      `json:"parameters,omitempty"`     // func
	PossibleValues []*PossibleValue `json:"possibleValues,omitempty"` // const
}

func (p *Property) decl(w *Context, parent *Base) {
	if p.Name == "" {
		p.Name = "obj"
	}
	// possibleValues
	if p.PossibleValues != nil {
		tname := w.newConst(p, parent)
		fmt.Fprintf(w, "%s %s",
			p.goSym(),
			tname,
		)
		return
	}
	// basic types
	if p.isBasic() {
		p.Base.decl(w, parent)
		return
	}
	// compound types
	tname := w.newType(p, parent)
	fmt.Fprintf(w, "%s %s",
		p.goSym(),
		tname,
	)
}

type PossibleValue struct { // const
	*Base
	Value string `json:"value,omitempty"`
}

type Event struct {
	*Base
	Return []*Property `json:"returns,omitempty"`
}

// w      *io.Writer // main writer
// xw     *io.Writer // seperate writer of paramter/Property Objects
// parent *Base      // for Object
func (e Event) decl(w *Context, p *Base) {
	fmt.Fprintf(w,
		`Evt%s%s = "%s"`,
		strings.Replace(w.base.goSym(), "Module", "", 1),
		e.goSym(),
		e.Name,
	)
}

type Method struct {
	*Base
	Signature  string      `json:"signature,omitempty"`
	Parameters []*Property `json:"parameters,omitempty"`
	Return     *Property   `json:"returns,omitempty"`
}

func (m *Method) decl(w *Context, parent *Base) {
	fmt.Fprintf(w, "%s func(", m.goSym())
	// parameters
	for _, p := range m.Parameters {
		p.decl(w, m.Base)
		fmt.Fprintf(w, ",")
	}
	fmt.Fprintf(w, ")")
	// returns
	if m.Return != nil {
		fmt.Fprintf(w, "(")
		m.Return.decl(w, m.Base)
		fmt.Fprintf(w, ")")
	}
}

// defBody write declaration and function body
func (m *Method) defStaticMethodBody(w *Context, rawMethodName string) {
	fmt.Fprintf(w, "\nfunc %s(", goSym(rawMethodName))
	// parameters
	for _, p := range m.Parameters {
		p.decl(w, w.base)
		fmt.Fprintf(w, ",")
	}
	fmt.Fprintf(w, ")")
	// return
	if m.Return != nil {
		fmt.Fprintf(w, " *js.Object ")
	}
	// body
	fmt.Fprintf(w, "{\n")
	fmt.Fprintf(w, "o := electron.Get(\"%s\")\n", w.base.Name)
	// ret
	if m.Return != nil {
		fmt.Fprintf(w, "ret := ")
	}
	// parameters
	fmt.Fprintf(w, "o.Call(\"%s\"", rawMethodName)
	for _, p := range m.Parameters {
		fmt.Fprintf(w, ", %s", p.goSym())
	}
	fmt.Fprintf(w, ")\n")
	// return
	if m.Return != nil {
		fmt.Fprintf(w, "return ret\n")
	}
	fmt.Fprintf(w, "}")
}

func (m *Method) defConstructorBody(w *Context, rawMethodName string) {
	fmt.Fprintf(w, "\nfunc %s(", goSym(rawMethodName))
	// parameters
	for _, p := range m.Parameters {
		p.decl(w, w.base)
		fmt.Fprintf(w, ",")
	}
	fmt.Fprintf(w, ")")
	// return type
	fmt.Fprintf(w, " *%s", w.base.goSym())
	// body
	fmt.Fprintf(w, "{\n")
	fmt.Fprintf(w, "o := electron.Get(\"%s\")\n", w.base.Name)
	fmt.Fprintf(w, "ret := ")
	// parameters
	fmt.Fprintf(w, "o.Invoke(\"New\"")
	for _, p := range m.Parameters {
		fmt.Fprintf(w, ", %s", p.goSym())
	}
	fmt.Fprintf(w, ")\n")
	// return
	fmt.Fprintf(w, "return Wrap%s(ret)\n", w.base.goSym())
	fmt.Fprintf(w, "}")
}

type Block struct {
	*Base
	// module
	Events     []*Event    `json:"events,omitempty"`
	Properties []*Property `json:"Properties,omitempty"`
	Methods    []*Method   `json:"Methods,omitempty"`
	// class
	InstanceName       string      `json:"instanceName,omitempty"`
	InstanceEvents     []*Event    `json:"instanceEvents,omitempty"`
	InstanceProperties []*Property `json:"instanceProperties,omitempty"`
	InstanceMethods    []*Method   `json:"instanceMethods,omitempty"`
	// standalone
	ConstructorMethod *Method   `json:"constructorMethod,omitempty"`
	StaticMethods     []*Method `json:"staticMethods,omitempty"`
}

type ApiFile []*Block

func (b *Block) declOther(w *Context) {
	if enableComment {
		b.comment(w)
	}
	// props
	fmt.Fprintf(w, "type %s struct {\n\t", b.goSym())
	fmt.Fprintf(w, "*js.Object\n\t")
	declSlice(b.Properties, w, b.Base)
	fmt.Fprintf(w, "}\n")
}

func (b *Block) declModule(w *Context) {
	// evnents
	if len(b.Events) > 0 {
		fmt.Fprintf(w, "\nconst (\n\t")
		declSlice(b.Events, w, nil)
		fmt.Fprintf(w, ")\n")
	}
	// props and methods
	if enableComment {
		b.comment(w)
	}
	fmt.Fprintf(w, "type %s struct {\n\t", b.goSym())
	if b.isEventEmitter() {
		fmt.Fprintf(w, "*events.Emitter\n\t")
	} else {
		fmt.Fprintf(w, "*js.Object\n\t")
	}
	declSlice(b.Properties, w, b.Base)
	declSlice(b.Methods, w, b.Base)
	fmt.Fprintf(w, "}\n")
	// getters
	getter := moduleGetterTemplate
	if b.isEventEmitter() {
		getter = moduleGetterTemplateWithEmitter
	}
	getter = strings.Replace(getter, "Module", b.goSym(), -1)
	getter = fmt.Sprintf(getter, b.Name)
	fmt.Fprintf(w, "%s", getter)
}

func (b *Block) isEventEmitter() bool {
	return len(b.Events)+len(b.InstanceEvents) > 0
}

func (b *Block) declClass(w *Context) {
	// evnents
	if len(b.InstanceEvents) > 0 {
		fmt.Fprintf(w, "\n// events \nconst (\n\t")
		declSlice(b.InstanceEvents, w, nil)
		fmt.Fprintf(w, ")\n")
	}
	// props and methods
	if enableComment {
		b.comment(w)
	}
	fmt.Fprintf(w, "type %s struct {\n\t", b.goSym())
	if b.isEventEmitter() {
		fmt.Fprintf(w, "*events.Emitter\n\t")
	} else {
		fmt.Fprintf(w, "*js.Object\n\t")
	}
	declSlice(b.InstanceProperties, w, b.Base)
	declSlice(b.InstanceMethods, w, b.Base)
	fmt.Fprintf(w, "}\n")
	// wrapper
	wrapper := strings.Replace(classWrapperTemplate, "Class", b.goSym(), -1)
	if b.isEventEmitter() {
		wrapper = strings.Replace(classWrapperTemplateWithEmitter, "ClassWithEmitter", b.goSym(), -1)
	}
	fmt.Fprintf(w, "%s", wrapper)
	// static methods
	for _, m := range b.StaticMethods {
		m.defStaticMethodBody(w, m.Name)
	}
	// constructorMethod
	if b.ConstructorMethod != nil {
		b.ConstructorMethod.defConstructorBody(w, "New"+b.goSym())
	}
}

func (a ApiFile) test() {
	for _, b := range a {
		if b.Name == "BrowserWindow" {
			log.Printf("BrowserWindow: %+v", *b)
		}
	}
}

func (a ApiFile) createElectronTemplate() error {
	opath := outDir + "/electron.go"
	w, err := os.Create(opath)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, electronTemplate)
	return err
}

func (a ApiFile) decl() error {
	// electron template
	err := a.createElectronTemplate()
	if err != nil {
		return err
	}
	// blocks
	for _, b := range a {
		log.Println("Processing", b.Base.Name)
		ctx, err := newContext(b.Base)
		if err != nil {
			log.Println(b.Name, err)
		}
		// decl
		if b.isModule() {
			b.declModule(ctx)
		} else if b.isClass() {
			b.declClass(ctx)
		} else {
			b.declOther(ctx)
		}
		// decl extra types
		ctx.declNewTypes()
		// adjust import
		ctx.adjustImport(b)
		// format
		ctx.formatCode()
		// write output
		if err = ctx.OutputToFile(); err != nil {
			log.Println(b.Name, err)
		}
	}
	// a.test()
	return nil
}
