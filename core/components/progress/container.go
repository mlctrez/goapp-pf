package progress

// from file "react-core/src/components/Progress/ProgressContainer.tsx"

type ContainerProps struct {
	// ProgressBarAriaProps - Properties needed for aria support. Optional.
	ProgressBarAriaProps any // AriaProps 
	// ParentId - Progress component DOM ID.
	ParentId string
	// Title - Progress title. The isTitleTruncated property will only affect string titles. Node title truncation
	// must be handled manually. Optional.
	Title any // React.ReactNode 
	// Label - Label to indicate what progress is showing. Optional.
	Label any // React.ReactNode 
	// Variant - Type of progress status. Optional.
	Variant any /* "danger" | "success" | "warning" */
	// MeasureLocation - Location of progress value. Optional.
	MeasureLocation any /* "outside" | "inside" | "top" | "none" */
	// Value - Actual progress value.
	Value int
	// IsTitleTruncated - Whether string title should be truncated. Optional.
	IsTitleTruncated bool
	// TooltipPosition - Position of the tooltip which is displayed if title is truncated. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}
