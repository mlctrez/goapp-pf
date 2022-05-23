package progress

// from file "react-core/src/components/Progress/Progress.tsx"

type Props struct {
	// ClassName - Classname for progress component. Optional.
	ClassName string
	// Size - Size variant of progress. Optional.
	Size any /* "sm" | "md" | "lg" */
	// MeasureLocation - Where the measure percent will be located. Optional.
	MeasureLocation any /* "outside" | "inside" | "top" | "none" */
	// Variant - Status variant of progress. Optional.
	Variant any /* "danger" | "success" | "warning" */
	// Title - Title above progress. The isTitleTruncated property will only affect string titles. Node title
	// truncation must be handled manually. Optional.
	Title any // React.ReactNode 
	// Label - Text description of current progress value to display instead of percentage. Optional.
	Label any // React.ReactNode 
	// Value - Actual value of progress. Optional.
	Value int
	// Id - DOM id for progress component. Optional.
	Id string
	// Min - Minimal value of progress. Optional.
	Min int
	// Max - Maximum value of progress. Optional.
	Max int
	// ValueText - Accessible text description of current progress value, for when value is not a percentage.
	// Use with label. Optional.
	ValueText string
	// IsTitleTruncated - Indicate whether to truncate the string title. Optional.
	IsTitleTruncated bool
	// TooltipPosition - Position of the tooltip which is displayed if title is truncated. Optional.
	TooltipPosition any /* "auto" | "top" | "bottom" | "left" | "right" */
	// AriaLabel - Adds accessible text to the ProgressBar. Required when title not used and there is not any
	// label associated with the progress bar. Optional.
	AriaLabel string
	// AriaLabelledby - Associates the ProgressBar with it's label for accessibility purposes. Required when
	// title not used. Optional.
	AriaLabelledby string
}
