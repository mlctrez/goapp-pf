package form

// from file "react-core/src/components/Form/FormFieldGroupExpandable.tsx"

type FieldGroupExpandableProps struct {
	// Children - Anything that can be rendered as form field group content. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the form field group. Optional.
	ClassName string
	// Header - Form field group header. Optional.
	Header any // React.ReactNode 
	// IsExpanded - Flag indicating if the form field group is initially expanded. Optional.
	IsExpanded bool
	// ToggleAriaLabel - Aria-label to use on the form field group toggle button. Optional.
	ToggleAriaLabel string
}
