package card

// from file "react-core/src/components/Card/CardBody.tsx"

type BodyProps struct {
	// Children - Content rendered inside the Card Body. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Card Body. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* keyof any // JSX.IntrinsicElements  */
	// IsFilled - Enables the body Content to fill the height of the card. Optional.
	IsFilled bool
}
