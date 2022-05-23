package checkbox

// from file "react-core/src/components/Checkbox/Checkbox.tsx"

type Props struct {
	// ClassName - Additional classes added to the Checkbox. Optional.
	ClassName string
	// IsValid - Flag to show if the Checkbox selection is valid or invalid. Optional.
	IsValid bool
	// IsDisabled - Flag to show if the Checkbox is disabled. Optional.
	IsDisabled bool
	// IsChecked - Flag to show if the Checkbox is checked. If null, the checkbox will be indeterminate (partially
	// checked). Optional.
	IsChecked any /* bool | "" */
	// Checked - Optional.
	Checked bool
	// OnChange - A callback for when the Checkbox selection changes. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// Label - Label text of the checkbox. Optional.
	Label any // React.ReactNode 
	// Id - Id of the checkbox.
	Id string
	// AriaLabel - Aria-label of the checkbox. Optional.
	AriaLabel string
	// Description - Description text of the checkbox. Optional.
	Description any // React.ReactNode 
	// Body - Body text of the checkbox. Optional.
	Body any // React.ReactNode 
}

type State struct {
	// OuiaStateId - 
	OuiaStateId string
}
