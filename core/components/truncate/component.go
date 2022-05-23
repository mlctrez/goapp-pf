package truncate

// from file "react-core/src/components/Truncate/Truncate.tsx"

type Props struct {
	// ClassName - Class to add to outer span. Optional.
	ClassName string
	// Content - Text to truncate.
	Content string
	// TrailingNumChars - The number of characters displayed in the second half of the truncation. Optional.
	TrailingNumChars int
	// Position - Where the text will be truncated. Optional.
	Position any /* "start" | "middle" | "end" */
	// TooltipPosition - Tooltip position. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}
