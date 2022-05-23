package chip

// from file "react-core/src/components/Chip/Chip.tsx"

type Props struct {
	// Children - Content rendered inside the chip text. Optional.
	Children any // React.ReactNode 
	// CloseBtnAriaLabel - Aria Label for close button. Optional.
	CloseBtnAriaLabel string
	// ClassName - Additional classes added to the chip item. Optional.
	ClassName string
	// IsOverflowChip - Flag indicating if the chip is an overflow chip. Optional.
	IsOverflowChip bool
	// IsReadOnly - Flag indicating if chip is read only. Optional.
	IsReadOnly bool
	// OnClick - Function that is called when clicking on the chip close button. Optional.
	OnClick func(event any // React.MouseEvent )
	// Component - Component that will be used for chip. It is recommended that <button> or <li>  are used when
	// the chip is an overflow chip. Optional.
	Component any // React.ReactNode 
	// TooltipPosition - Position of the tooltip which is displayed if text is longer. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}

type State struct {
	// IsTooltipVisible - 
	IsTooltipVisible bool
	// OuiaStateId - 
	OuiaStateId string
}
