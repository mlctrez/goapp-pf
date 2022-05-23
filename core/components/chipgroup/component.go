package chipgroup

// from file "react-core/src/components/ChipGroup/ChipGroup.tsx"

type Props struct {
	// Children - Content rendered inside the chip group. Should be <Chip> elements. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the chip item. Optional.
	ClassName string
	// DefaultIsOpen - Flag for having the chip group default to expanded. Optional.
	DefaultIsOpen bool
	// ExpandedText - Customizable "Show Less" text string. Optional.
	ExpandedText string
	// CollapsedText - Customizeable template string. Use variable "${remaining}" for the overflow chip count.
	// Optional.
	CollapsedText string
	// CategoryName - Category name text for the chip group category.  If this prop is supplied the chip group
	// with have a label and category styling applied. Optional.
	CategoryName string
	// AriaLabel - Aria label for chip group that does not have a category name. Optional.
	AriaLabel string
	// NumChips - Set number of chips to show before overflow. Optional.
	NumChips int
	// IsClosable - Flag if chip group can be closed. Optional.
	IsClosable bool
	// CloseBtnAriaLabel - Aria label for close button. Optional.
	CloseBtnAriaLabel string
	// OnClick - Function that is called when clicking on the chip group close button. Optional.
	OnClick func(event any // React.MouseEvent )
	// OnOverflowChipClick - Function that is called when clicking on the overflow (expand/collapse) chip button.
	// Optional.
	OnOverflowChipClick func(event any // React.MouseEvent )
	// TooltipPosition - Position of the tooltip which is displayed if the category name text is longer. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}

type State struct {
	// IsOpen - 
	IsOpen bool
	// IsTooltipVisible - 
	IsTooltipVisible bool
}
