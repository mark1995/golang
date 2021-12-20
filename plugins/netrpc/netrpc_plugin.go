package netrpc

import (
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"log"
)

type NetrpcPlugin struct {
	*generator.Generator
}

func (p *NetrpcPlugin) Name() string {
	return "netrpc"
}

func (p *NetrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *NetrpcPlugin) GenerateImports(file *generator.FileDescriptor)  {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *NetrpcPlugin) Generate(file *generator.FileDescriptor)  {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *NetrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	//p.P("// TODO: import code")

	p.P(`import "net/rpc"`)
}

func (p *NetrpcPlugin) genServiceCode(svc *descriptorpb.ServiceDescriptorProto) {
	//p.P("// TODO: service code, Name = " + svc.GetName())
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}
	p.P(buf.String())
}

func (p *NetrpcPlugin) buildServiceSpec(svc *descriptorpb.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:    generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}
	return spec
}


type ServiceSpec struct {
	ServiceName string
	MethodList []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName string
	InputTypeName string
	OutputTypeName string
}


const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}} (
	srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
	if err := srv.RegisterName("{{.ServiceName}}", x); err != nil { 
		return err
	}
	return nil
}

type {{.ServiceName}}Client struct {
	*rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}} (network, address string)  (
	*{{.ServiceName}}Client, error,
) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err	
	}
	return &{{.ServiceName}}Client{Client:c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}}) error {
	return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`
