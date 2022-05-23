package page

// from file "react-core/src/components/Page/PageGroup.tsx"

type GroupProps struct {
	// ClassName - Additional classes to apply to the PageGroup. Optional.
	ClassName string
	// Children - Content rendered inside of the PageGroup. Optional.
	Children any // React.ReactNode 
	// Sticky - Modifier indicating if PageGroup is sticky to the top or bottom. Optional.
	Sticky any /* "top" | "bottom" */
	// HasShadowTop - Modifier indicating if PageGroup should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Modifier indicating if PageGroup should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageGroup has a scrolling overflow. Optional.
	HasOverflowScroll bool
}
