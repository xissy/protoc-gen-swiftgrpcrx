package swiftgrpcrx

import (
	"strings"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/xissy/protoc-gen-swiftgrpcrx/generator"
)

func init() {
	generator.RegisterPlugin(new(swiftgrpcrx))
}

// swiftgrpcrx is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for go-swiftgrpcrx support.
type swiftgrpcrx struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "swiftgrpcrx".
func (g *swiftgrpcrx) Name() string {
	return "swiftgrpcrx"
}

// Init initializes the plugin.
func (g *swiftgrpcrx) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *swiftgrpcrx) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *swiftgrpcrx) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *swiftgrpcrx) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *swiftgrpcrx) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	g.GenerateImports(file)
	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *swiftgrpcrx) GenerateImports(file *generator.FileDescriptor) {
	g.P("import Foundation")
	g.P("import RxSwift")
	g.P("import SwiftGRPC")
	g.P()
}

// methodTypeName returns the method inputType or outputType's type name only.
// e.g. `.users.GetRequest` to `GetRequest`
func (g *swiftgrpcrx) methodTypeName(methodType string) string {
	i := strings.LastIndex(methodType, ".")
	return methodType[i+1:]
}

// generateService generates all the code for the named service.
func (g *swiftgrpcrx) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	origServName := service.GetName()
	servName := generator.CamelCase(origServName)
	servAlias := servName + "Service"

	g.P("internal extension ", servName, "_", servAlias, "Client {")
	g.P()

	for _, method := range service.Method {
		if !method.GetServerStreaming() {
			// Unary RPC method
			methodName := strings.ToLower((*method.Name)[:1]) + (*method.Name)[1:]
			inputTypeName := servName + "_" + g.typeName(*method.InputType)
			outputTypeName := servName + "_" + g.typeName(*method.OutputType)
			g.P("  /// RxSwift. Unary.")
			g.P("  internal func ", methodName, "(_ request: ", inputTypeName, ", metadata customMetadata: Metadata?) -> Observable<", outputTypeName, "> {")
			g.P("    return Observable.create { observer in")
			g.P("      _ = try? self.", methodName, "(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in")
			g.P("        guard let resp: ", outputTypeName, " = resp else {")
			g.P("          observer.onError(RPCError.callError(result))")
			g.P("          return")
			g.P("        }")
			g.P("        observer.onNext(resp)")
			g.P("      })")
			g.P("      return Disposables.create()")
			g.P("    }")
			g.P("  }")
		} else {
			// Streaming RPC method
		}
		g.P()
	}

	g.P("}")
	g.P()
}
