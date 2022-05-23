package card

// from file "react-core/src/components/Card/CardTitle.tsx"

type TitleProps struct {
	// Children - Content rendered inside the CardTitle. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the CardTitle. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* keyof any // JSX.IntrinsicElements  */
}
