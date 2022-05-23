package datalist

// from file "react-core/src/components/DataList/DataListCheck.tsx"

type CheckProps struct {
	// ClassName - Additional classes added to the DataList item checkbox. Optional.
	ClassName string
	// IsValid - Flag to show if the DataList checkbox selection is valid or invalid. Optional.
	IsValid bool
	// IsDisabled - Flag to show if the DataList checkbox is disabled. Optional.
	IsDisabled bool
	// IsChecked - Flag to show if the DataList checkbox is checked when it is controlled by React state. Both
	// isChecked and checked are valid, but only use one. To make the DataList checkbox uncontrolled, instead
	// use the defaultChecked prop, but do not use both. Optional.
	IsChecked bool
	// Checked - Flag to show if the DataList checkbox is checked when it is controlled by React state. Both
	// isChecked and checked are valid, but only use one. To make the DataList checkbox uncontrolled, instead
	// use the defaultChecked prop, but do not use both. Optional.
	Checked bool
	// DefaultChecked - Flag to set default value of DataList checkbox when it is uncontrolled by React state.
	// To make the DataList checkbox controlled, instead use the isChecked prop, but do not use both. Optional.
	DefaultChecked bool
	// OnChange - A callback for when the DataList checkbox selection changes. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// AriaLabelledby - Aria-labelledby of the DataList checkbox.
	AriaLabelledby string
	// OtherControls - Flag to indicate if other controls are used in the DataListItem. Optional.
	OtherControls bool
}
