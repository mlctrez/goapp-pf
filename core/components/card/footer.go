package card

// from file "react-core/src/components/Card/CardFooter.tsx"

type FooterProps struct {
	// Children - Content rendered inside the Card Footer. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Footer. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* keyof any // JSX.IntrinsicElements  */
}
