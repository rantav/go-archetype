package transformer

import "fmt"

func newTransformer(spec transformationSpec) Transformer {
	switch spec.Type {
	case TransformationTypeInclude:
		return newIncludeTransformer(spec)
	case TransformationTypeReplace:
		return newTextReplacer(spec)
	case TransformationTypeRename:
		return newFileRenamer(spec)
	default:
		panic(fmt.Sprintf("Unknown transformation type: %s ", spec.Type))
	}
}
