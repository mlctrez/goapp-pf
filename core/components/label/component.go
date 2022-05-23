package label

// from file "react-core/src/components/Label/Label.tsx"

type Props struct {
	// Children - Content rendered inside the label. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the label. Optional.
	ClassName string
	// Color - Color of the label. Optional.
	Color any /* "blue" | "cyan" | "green" | "orange" | "purple" | "red" | "grey" */
	// Variant - Variant of the label. Optional.
	Variant any /* "outline" | "filled" */
	// IsCompact - Flag indicating the label is compact. Optional.
	IsCompact bool
	// IsEditable - Optional.
	IsEditable bool
	// EditableProps - Optional.
	EditableProps any
	// OnEditComplete - Optional.
	OnEditComplete func(newText string)
	// OnEditCancel - Optional.
	OnEditCancel func(previousText string)
	// IsTruncated - Flag indicating the label text should be truncated. Optional.
	IsTruncated bool
	// TooltipPosition - Position of the tooltip which is displayed if text is truncated. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// Icon - Icon added to the left of the label text. Optional.
	Icon any // React.ReactNode 
	// OnClose - Close click callback for removable labels. If present, label will have a close button. Optional.
	OnClose func(event any // React.MouseEvent )
	// CloseBtn - Node for custom close button. Optional.
	CloseBtn any // React.ReactNode 
	// CloseBtnAriaLabel - Aria label for close button. Optional.
	CloseBtnAriaLabel string
	// CloseBtnProps - Additional properties for the default close button. Optional.
	CloseBtnProps any
	// Href - Href for a label that is a link. If present, the label will change to an anchor element. Optional.
	Href string
	// IsOverflowLabel - Flag indicating if the label is an overflow label. Optional.
	IsOverflowLabel bool
	// Render - Forwards the label content and className to rendered function.  Use this prop for react router
	// support. Optional.
	Render func( map[string]string /* { className:string, "content":any // React.ReactNode , "componentRef":any } */) any // React.ReactNode 
}
