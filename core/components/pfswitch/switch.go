package pfswitch

// from file "react-core/src/components/Switch/Switch.tsx"

type SwitchProps struct {
	// Id - id for the label. Optional.
	Id string
	// ClassName - Additional classes added to the switch. Optional.
	ClassName string
	// Label - Text value for the label when on. Optional.
	Label any // React.ReactNode 
	// LabelOff - Text value for the label when off. Optional.
	LabelOff any // React.ReactNode 
	// IsChecked - Flag to show if the switch is checked. Optional.
	IsChecked bool
	// HasCheckIcon - Flag to show if the switch has a check icon. Optional.
	HasCheckIcon bool
	// IsDisabled - Flag to show if the switch is disabled. Optional.
	IsDisabled bool
	// OnChange - A callback for when the switch selection changes. (isChecked, event) => {}. Optional.
	OnChange func(checked bool, event any // React.FormEvent )
	// AriaLabel - Adds accessible text to the switch, and should describe the isChecked="true" state. When label
	// is defined, aria-label should be set to the text string that is visible when isChecked is true. Optional.
	AriaLabel string
	// IsReversed - Flag to reverse the layout of toggle and label (toggle on right). Optional.
	IsReversed bool
}
