package dropdown

// from file "react-core/src/components/Dropdown/DropdownToggleCheckbox.tsx"

type ToggleCheckboxProps struct {
	// ClassName - Additional classes added to the DropdownToggleCheckbox. Optional.
	ClassName string
	// IsValid - Flag to show if the checkbox selection is valid or invalid. Optional.
	IsValid bool
	// IsDisabled - Flag to show if the checkbox is disabled. Optional.
	IsDisabled bool
	// IsChecked - Flag to show if the checkbox is checked. Optional.
	IsChecked any /* bool | "" */
	// Checked - Alternate Flag to show if the checkbox is checked. Optional.
	Checked any /* bool | "" */
	// OnChange - A callback for when the checkbox selection changes. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// Children - Element to be rendered inside the <span>. Optional.
	Children any // React.ReactNode 
	// Id - Id of the checkbox.
	Id string
	// AriaLabel - Aria-label of the checkbox.
	AriaLabel string
}
