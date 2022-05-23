package form

// from file "react-core/src/components/Form/Form.tsx"

type Props struct {
	// Children - Anything that can be rendered as Form content. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Form. Optional.
	ClassName string
	// IsHorizontal - Sets the Form to horizontal. Optional.
	IsHorizontal bool
	// IsWidthLimited - Limits the max-width of the form. Optional.
	IsWidthLimited bool
	// MaxWidth - Sets a custom max-width for the form. Optional.
	MaxWidth string
}
