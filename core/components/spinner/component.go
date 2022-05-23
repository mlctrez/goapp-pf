package spinner

// from file "react-core/src/components/Spinner/Spinner.tsx"

type Props struct {
	// ClassName - Additional classes added to the Spinner. Optional.
	ClassName string
	// Size - Size variant of progress. Optional.
	Size any /* "sm" | "md" | "lg" | "xl" */
	// AriaValuetext - Text describing that current loading status or progress. Optional.
	AriaValuetext string
	// IsSVG - Whether to use an SVG (new) rather than a span (old). Optional.
	IsSVG bool
	// Diameter - Diameter of spinner set as CSS variable. Optional.
	Diameter string
	// AriaLabel - Accessible label to describe what is loading. Optional.
	AriaLabel string
	// AriaLabelledBy - Id of element which describes what is being loaded. Optional.
	AriaLabelledBy string
}
