package lint

import (
	"github.com/emicklei/proto"
	"github.com/uber/prototool/internal/strs"
	"github.com/uber/prototool/internal/text"
	"path/filepath"
	"strings"
)

var messageNamesSameAsFileNameLinter = NewLinter(
	"MESSAGE_NAME_SAME_AS_FILE_NAME",
	"Verifies message name is same as file name",
	checkMessageNamesSameAsFileName,
)

func checkMessageNamesSameAsFileName(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(&messageNamesSameAsFileNameVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type messageNamesSameAsFileNameVisitor struct {
	baseAddVisitor
	filename string
}


func (v *messageNamesSameAsFileNameVisitor) OnStart(descriptor *FileDescriptor) error {
	v.filename = descriptor.Filename
	return nil
}



func (v *messageNamesSameAsFileNameVisitor) Finally() error {

	return nil
}


func (v *messageNamesSameAsFileNameVisitor) VisitMessage(message *proto.Message) {
	// fmt.Println("filename: ", v.filename)
	// for nested messages
	for _, child := range message.Elements {
		child.Accept(v)
	}
	if message.IsExtend {
		return
	}
	if v.filename == "" {
		return
	}
	// 验证文件名与Message名称是否相同
	fileName := filepath.Base(v.filename)
	fileNameNoSuffix := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	if message.Name != strs.ToUpperCamelCase(fileNameNoSuffix) {
		v.AddFailuref(message.Position, "Message name %q must be same with file name convert result of Camel Case", message.Name)
	}
}
