package plugin

import (
	"context"

	"github.com/bufbuild/protoplugin"
	"github.com/mfridman/protoc-gen-go-json/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	defaultFilenameSuffix = ".pb.json.go"
	binaryFilenameSuffix  = ".pb.binary.go" // 二进制序列化文件的后缀
)

func Handle(
	ctx context.Context,
	env protoplugin.PluginEnv,
	w protoplugin.ResponseWriter,
	r protoplugin.Request,
) error {
	p, err := protogen.Options{}.New(r.CodeGeneratorRequest())
	if err != nil {
		return err
	}
	opt, err := parseOptions(r.Parameter())
	if err != nil {
		return err
	}
	if err := generate(p, opt); err != nil {
		p.Error(err)
	}

	response := p.Response()
	w.AddCodeGeneratorResponseFiles(response.GetFile()...)
	w.AddError(response.GetError())
	w.SetFeatureProto3Optional()
	// 修改这里以支持 Proto2 和 Proto3
	w.SetFeatureSupportsEditions(
		descriptorpb.Edition_EDITION_PROTO2, // 添加对 Proto2 的支持
		descriptorpb.Edition_EDITION_PROTO3,
	)
	return nil
}

func generate(p *protogen.Plugin, opt *gen.Options) error {
	for _, f := range p.Files {
		if !f.Generate || len(f.Messages) == 0 {
			continue
		}

		g := p.NewGeneratedFile(f.GeneratedFilenamePrefix+defaultFilenameSuffix, f.GoImportPath)
		if err := gen.ApplyTemplate(g, f, opt); err != nil {
			g.Skip()
			return err
		}

		if opt.GenerateBinary || opt.GenerateMarshal || opt.GenerateUnmarshal {
			binaryG := p.NewGeneratedFile(f.GeneratedFilenamePrefix+binaryFilenameSuffix, f.GoImportPath)
			if err := gen.ApplyBinaryTemplate(binaryG, f, opt); err != nil {
				binaryG.Skip()
				return err
			}
		}
	}
	return nil
}
