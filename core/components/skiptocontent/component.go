package skiptocontent

// from file "react-core/src/components/SkipToContent/SkipToContent.tsx"

type Props struct {
	// Href - The skip to content link.
	Href string
	// Children - Content to display within the skip to content component, typically a string. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional styles to apply to the skip to content component. Optional.
	ClassName string
	// Show - Forces the skip to content to display. This is primarily for demonstration purposes and would not
	// normally be used. Optional.
	Show bool
}
