package alertgroup

// from file "react-core/src/components/AlertGroup/AlertGroup.tsx"

type Props struct {
	// ClassName - Additional classes added to the AlertGroup. Optional.
	ClassName string
	// Children - Alerts to be rendered in the AlertGroup. Optional.
	Children any // React.ReactNode 
	// IsToast - Toast notifications are positioned at the top right corner of the viewport. Optional.
	IsToast bool
	// IsLiveRegion - Turns the container into a live region so that changes to content within the AlertGroup,
	// such as appending an Alert, are reliably announced to assistive technology. Optional.
	IsLiveRegion bool
	// AppendTo - Determine where the alert is appended to. Optional.
	AppendTo any /* any // HTMLElement  | func() any // HTMLElement  */
	// OnOverflowClick - Function to call if user clicks on overflow message. Optional.
	OnOverflowClick func()
	// OverflowMessage - Custom text to show for the overflow message. Optional.
	OverflowMessage string
}

type State struct {
	// Container - 
	Container any // HTMLElement 
}
