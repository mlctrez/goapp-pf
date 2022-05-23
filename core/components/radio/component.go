package radio

// from file "react-core/src/components/Radio/Radio.tsx"

type Props struct {
	// ClassName - Additional classes added to the radio. Optional.
	ClassName string
	// Id - Id of the radio.
	Id string
	// IsLabelWrapped - Flag to show if the radio label is wrapped on small screen. Optional.
	IsLabelWrapped bool
	// IsLabelBeforeButton - Flag to show if the radio label is shown before the radio button. Optional.
	IsLabelBeforeButton bool
	// Checked - Flag to show if the radio is checked. Optional.
	Checked bool
	// IsChecked - Flag to show if the radio is checked. Optional.
	IsChecked bool
	// IsDisabled - Flag to show if the radio is disabled. Optional.
	IsDisabled bool
	// IsValid - Flag to show if the radio selection is valid or invalid. Optional.
	IsValid bool
	// Label - Label text of the radio. Optional.
	Label any // React.ReactNode 
	// Name - Name for group of radios.
	Name string
	// OnChange - A callback for when the radio selection changes. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// AriaLabel - Aria label for the radio. Optional.
	AriaLabel string
	// Description - Description text of the radio. Optional.
	Description any // React.ReactNode 
	// Body - Body of the radio. Optional.
	Body any // React.ReactNode 
}
