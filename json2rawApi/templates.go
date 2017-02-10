package main

var moduleGetterTemplate = `
func GetModule() *Module {
	o := electron.Get("%s")
	return &Module{
		Object: o,
	}
}
`

var moduleGetterTemplateWithEmitter = `
func GetModule() *Module {
	o := electron.Get("%s")
	return &Module{
		Emitter: events.New(o),
	}
}
`

var classWrapperTemplateWithEmitter = `
func WrapClassWithEmitter(o *js.Object) *ClassWithEmitter {
	return &ClassWithEmitter{
		Emitter: events.New(o),
	}
}
`

var classWrapperTemplate = `
func WrapClass(o *js.Object) *Class {
	return &Class{
		Object: o,
	}
}
`
