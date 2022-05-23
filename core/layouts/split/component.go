package split

// from file "react-core/src/layouts/Split/Split.tsx"

type Props struct {
	// HasGutter - Adds space between children. Optional.
	HasGutter bool
	// IsWrappable - Allows children to wrap. Optional.
	IsWrappable bool
	// Children - content rendered inside the Split layout. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Split layout. Optional.
	ClassName string
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any // React.ReactNode 
}
