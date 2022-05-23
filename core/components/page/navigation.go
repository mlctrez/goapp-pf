package page

// from file "react-core/src/components/Page/PageNavigation.tsx"

type NavigationProps struct {
	// ClassName - Additional classes to apply to the PageNavigation. Optional.
	ClassName string
	// Children - Content rendered inside of the PageNavigation. Optional.
	Children any // React.ReactNode 
	// IsWidthLimited - Limits the width of the PageNavigation. Optional.
	IsWidthLimited bool
	// Sticky - Modifier indicating if the PageNavigation is sticky to the top or bottom. Optional.
	Sticky any /* "top" | "bottom" */
	// HasShadowTop - Flag indicating if PageNavigation should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Flag indicating if PageNavigation should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageNavigation has a scrolling overflow. Optional.
	HasOverflowScroll bool
}
