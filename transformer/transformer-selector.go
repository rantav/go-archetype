package transformer

import (
	"fmt"

	"github.com/rantav/go-archetype/log"
)

func newTransformer(spec transformationSpec, logger log.Logger) Transformer {
	switch spec.Type {
	case TransformationTypeInclude:
		return newIncludeTransformer(spec, logger)
	case TransformationTypeReplace:
		return newTextReplacer(spec)
	case TransformationTypeRename:
		return newFileRenamer(spec)
	default:
		panic(fmt.Sprintf("Unknown transformation type: %s", spec.Type))
	}
}
