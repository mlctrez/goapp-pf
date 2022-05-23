package form

// from file "react-core/src/components/Form/FormGroup.tsx"

type GroupProps struct {
	// Children - Anything that can be rendered as FormGroup content. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the FormGroup. Optional.
	ClassName string
	// Label - Label text before the field. Optional.
	Label any // React.ReactNode 
	// LabelInfo - Additional label information displayed after the label. Optional.
	LabelInfo any // React.ReactNode 
	// LabelIcon - Sets an icon for the label. For providing additional context. Host element for Popover. Optional.
	LabelIcon any // React.ReactElement 
	// IsRequired - Sets the FormGroup required. Optional.
	IsRequired bool
	// Validated - Sets the FormGroup validated. If you set to success, text color of helper text will be modified
	// to indicate valid state. If set to error, text color of helper text will be modified to indicate error
	// state. If set to warning, text color of helper text will be modified to indicate warning state. Optional.
	Validated any /* "success" | "warning" | "error" | "default" */
	// IsInline - Sets the FormGroup isInline. Optional.
	IsInline bool
	// IsStack - Sets the FormGroupControl to be stacked. Optional.
	IsStack bool
	// HasNoPaddingTop - Removes top spacer from label. Optional.
	HasNoPaddingTop bool
	// HelperText - Helper text regarding the field. It can be a simple text or an object. Optional.
	HelperText any // React.ReactNode 
	// IsHelperTextBeforeField - Flag to position the helper text before the field. False by default. Optional.
	IsHelperTextBeforeField bool
	// HelperTextInvalid - Helper text after the field when the field is invalid. It can be a simple text or
	// an object. Optional.
	HelperTextInvalid any // React.ReactNode 
	// HelperTextIcon - Icon displayed to the left of the helper text. Optional.
	HelperTextIcon any // React.ReactNode 
	// HelperTextInvalidIcon - Icon displayed to the left of the helper text when the field is invalid. Optional.
	HelperTextInvalidIcon any // React.ReactNode 
	// FieldId - ID of the included field. It has to be the same for proper working.
	FieldId string
}
