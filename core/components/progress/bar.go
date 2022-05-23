package progress

// from file "react-core/src/components/Progress/ProgressBar.tsx"

type AriaProps struct {
	// AriaLabelledby - Optional.
	AriaLabelledby string
	// AriaLabel - Optional.
	AriaLabel string
	// AriaValuemin - Optional.
	AriaValuemin int
	// AriaValuenow - Optional.
	AriaValuenow int
	// AriaValuemax - Optional.
	AriaValuemax int
	// AriaValuetext - Optional.
	AriaValuetext string
}

type BarProps struct {
	// Children - What should be rendered inside progress bar. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes for Progres bar. Optional.
	ClassName string
	// Value - Actual progress value.
	Value int
	// ProgressBarAriaProps - Minimal value of progress. Optional.
	ProgressBarAriaProps any // AriaProps 
}
