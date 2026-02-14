// protoc-gen-go-identifier generates Go helper methods for identifier-type
// protobuf messages marked with the (plugins.identifier.v1.identifier) option.
//
// Generated methods include:
//   - As<Type>(v) - Constructor function
//   - Unwrap() - Nil-safe value extraction
//   - Equal(other) - Value-based equality comparison
//   - Clone() - Deep copy using proto.Clone
package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := generateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}

