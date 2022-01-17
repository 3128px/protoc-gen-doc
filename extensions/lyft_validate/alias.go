package extensions

import evp "github.com/3128px/protoc-gen-doc/v2/extensions/envoyproxy_validate"

// ValidateRule represents a single validator rule from the (validate.rules) method option extension.
type ValidateRule = evp.ValidateRule

// ValidateExtension contains the rules set by the (validate.rules) method option extension.
type ValidateExtension = evp.ValidateExtension
