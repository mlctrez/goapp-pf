package form

// from file "react-core/src/components/Form/FormFieldGroupToggle.tsx"

type FieldGroupToggleProps struct {
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// OnToggle - Callback for onClick.
	OnToggle func()
	// IsExpanded - Flag indicating if the toggle is expanded.
	IsExpanded bool
	// AriaLabel - Aria label of the toggle button. Optional.
	AriaLabel string
	// AriaLabelledby - Sets the aria-labelledby attribute on the toggle button element. Optional.
	AriaLabelledby string
	// ToggleId - The id applied to the toggle button. Optional.
	ToggleId string
}
