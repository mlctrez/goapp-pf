package avatar

// from file "react-core/src/components/Avatar/Avatar.tsx"

type Props struct {
	// ClassName - Additional classes added to the Avatar. Optional.
	ClassName string
	// Src - Attribute that specifies the URL of the image for the Avatar. Optional.
	Src string
	// Alt - Attribute that specifies the alternate text of the image for the Avatar.
	Alt string
	// Border - Border to add. Optional.
	Border any /* "light" | "dark" */
	// Size - Size variant of avatar. Optional.
	Size any /* "sm" | "md" | "lg" | "xl" */
}
