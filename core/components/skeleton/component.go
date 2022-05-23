package skeleton

// from file "react-core/src/components/Skeleton/Skeleton.tsx"

type Props struct {
	// ClassName - Additional classes added to the Skeleton. Optional.
	ClassName string
	// Width - The width of the Skeleton. Must specify pixels or percentage. Optional.
	Width string
	// Height - The height of the Skeleton. Must specify pixels or percentage. Optional.
	Height string
	// FontSize - The font size height of the Skeleton. Optional.
	FontSize any /* "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "4xl" */
	// Shape - The shape of the Skeleton. Optional.
	Shape any /* "circle" | "square" */
	// ScreenreaderText - Text read just to screen reader users. Optional.
	ScreenreaderText string
}
