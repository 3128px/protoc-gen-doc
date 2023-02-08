package extensions

import (
	"github.com/3128px/protoc-gen-doc/v2/extensions"
)

type IstioObjectSpecOptions struct {
	AtType string `json:"@type"`
}

type RbacRequires struct {
	Permissions                      []string `json:"permissions"`
	RawPermissions                   []string `json:"rawPermissions"`
	DeferPermissioCheckToApplication bool     `json:"deferPermissioCheckToApplication"`
}

func parseRequires(payload interface{}) interface{} {
	requires, ok := payload.(*RequiredPermission)
	if !ok {
		return nil
	}

	rbacRequires := RbacRequires{
		DeferPermissioCheckToApplication: requires.DeferPermissionCheckToApplication,
	}

	if len(requires.Permissions) > 0 {
		for _, permission := range requires.Permissions {
			rbacRequires.Permissions = append(rbacRequires.Permissions, permission.String())
		}
	}
	rbacRequires.RawPermissions = []string{}
	if requires.RawPermissions != nil {
		rbacRequires.RawPermissions = append(rbacRequires.RawPermissions, requires.RawPermissions...)
	}
	return rbacRequires
}

func init() {
	extensions.SetTransformer("tetrateio.api.tsb.rbac.v2.requires", func(payload interface{}) interface{} {
		return parseRequires(payload)
	})

	extensions.SetTransformer("tetrateio.api.tsb.rbac.v2.default_requires", func(payload interface{}) interface{} {
		return parseRequires(payload)
	})

	extensions.SetTransformer("tetrateio.api.tsb.types.v2.spec", func(payload interface{}) interface{} {
		spec, ok := payload.(*IstioObjectSpec)
		if !ok {
			return nil
		}
		return IstioObjectSpecOptions{
			AtType: "type.googleapis.com/" + spec.Type,
		}
	})
}
