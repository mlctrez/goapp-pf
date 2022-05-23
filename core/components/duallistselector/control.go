package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorControl.tsx"

type ControlProps struct {
	// Children - Content to be rendered in the dual list selector control. Optional.
	Children any // React.ReactNode 
	// InnerRef - Optional.
	InnerRef any // React.Ref 
	// IsDisabled - Flag indicating the control is disabled. Optional.
	IsDisabled bool
	// ClassName - Additional classes applied to the dual list selector control. Optional.
	ClassName string
	// OnClick - Callback fired when dual list selector control is selected. Optional.
	OnClick func(event any // React.MouseEvent )
	// AriaLabel - Accessible label for the dual list selector control. Optional.
	AriaLabel string
	// TooltipContent - Content to be displayed in a tooltip on hover of control. Optional.
	TooltipContent any // React.ReactNode 
	// TooltipProps - Additional tooltip properties passed to the tooltip. Optional.
	TooltipProps any
}
