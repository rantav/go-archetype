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
	default:
		panic(fmt.Sprintf("Unknown transformation type: %s", spec.Type))
	}
}
