package form

// from file "react-core/src/components/Form/InternalFormFieldGroup.tsx"

type InternalFormFieldGroupProps struct {
	// Children - Anything that can be rendered as form field group content. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the form field group. Optional.
	ClassName string
	// Header - Form field group header. Optional.
	Header any
	// IsExpandable - Flag indicating if the field group is expandable. Optional.
	IsExpandable bool
	// IsExpanded - Flag indicate if the form field group is expanded. Modifies the card to be expandable. Optional.
	IsExpanded bool
	// OnToggle - Function callback called when user clicks toggle button. Optional.
	OnToggle func()
	// ToggleAriaLabel - Aria-label to use on the form field group toggle button. Optional.
	ToggleAriaLabel string
}
