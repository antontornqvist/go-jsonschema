// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

type YamlMultilineDescriptions struct {
	// I'm a multiline description in a literal block. Literal blocks, on the other
	// hand,
	// should have hard line breaks and may end in a line break, too.
	//
	// This may look funky when Go code is generated to a specific line width, though.
	//
	Bar *string `json:"bar,omitempty" yaml:"bar,omitempty" mapstructure:"bar,omitempty"`

	// I'm a multiline description in a folded block. Folded blocks should not have
	// hard line breaks after parsing. They should also not end in a line break.
	Foo *string `json:"foo,omitempty" yaml:"foo,omitempty" mapstructure:"foo,omitempty"`
}