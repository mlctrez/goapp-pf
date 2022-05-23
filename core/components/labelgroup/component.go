package labelgroup

// from file "react-core/src/components/LabelGroup/LabelGroup.tsx"

type Props struct {
	// Children - Content rendered inside the label group. Should be <Label> elements. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the label item. Optional.
	ClassName string
	// DefaultIsOpen - Flag for having the label group default to expanded. Optional.
	DefaultIsOpen bool
	// ExpandedText - Customizable "Show Less" text string. Optional.
	ExpandedText string
	// CollapsedText - Customizeable template string. Use variable "${remaining}" for the overflow label count.
	// Optional.
	CollapsedText string
	// CategoryName - Category name text for the label group category.  If this prop is supplied the label group
	// with have a label and category styling applied. Optional.
	CategoryName string
	// AriaLabel - Aria label for label group that does not have a category name. Optional.
	AriaLabel string
	// NumLabels - Set number of labels to show before overflow. Optional.
	NumLabels int
	// IsClosable - Flag if label group can be closed. Optional.
	IsClosable bool
	// IsCompact - Flag indicating the labels in the group are compact. Optional.
	IsCompact bool
	// CloseBtnAriaLabel - Aria label for close button. Optional.
	CloseBtnAriaLabel string
	// OnClick - Function that is called when clicking on the label group close button. Optional.
	OnClick func(event any // React.MouseEvent )
	// TooltipPosition - Position of the tooltip which is displayed if the category name text is longer. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// IsVertical - Flag to implement a vertical layout. Optional.
	IsVertical bool
	// IsEditable - Optional.
	IsEditable bool
	// HasEditableTextArea - Optional.
	HasEditableTextArea bool
	// EditableTextAreaProps - Optional.
	EditableTextAreaProps any
}

type State struct {
	// IsOpen - 
	IsOpen bool
	// IsTooltipVisible - 
	IsTooltipVisible bool
}
