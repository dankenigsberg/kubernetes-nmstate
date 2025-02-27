package shared

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type RawState []byte

// State contains the namestatectl yaml [1] as string instead of golang struct
// so we don't need to be in sync with the schema.
//
// [1] https://github.com/nmstate/nmstate/blob/base/libnmstate/schemas/operational-state.yaml
// +kubebuilder:validation:Type=object
type State struct {
	Raw RawState `json:"-"`
}

func NewState(raw string) State {
	return State{Raw: RawState(raw)}
}

// This overrides the State type [1] so we can do a custom marshaling of
// nmstate yaml without the need to have golang code representing the
// nmstate schema

// [1] https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (_ State) OpenAPISchemaType() []string { return []string{"object"} }
