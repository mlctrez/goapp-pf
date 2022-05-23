package bullseye

// from file "react-core/src/layouts/Bullseye/Bullseye.tsx"

type Props struct {
	// Children - content rendered inside the Bullseye layout. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Bullseye layout. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* keyof any // JSX.IntrinsicElements  */
}
