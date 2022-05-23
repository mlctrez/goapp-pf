package page

// from file "react-core/src/components/Page/PageBreadcrumb.tsx"

type BreadcrumbProps struct {
	// ClassName - Additional classes to apply to the PageBreadcrumb. Optional.
	ClassName string
	// Children - Content rendered inside of the PageBreadcrumb. Optional.
	Children any // React.ReactNode 
	// IsWidthLimited - Limits the width of the breadcrumb. Optional.
	IsWidthLimited bool
	// Sticky - Modifier indicating if the PageBreadcrumb is sticky to the top or bottom. Optional.
	Sticky any /* "top" | "bottom" */
	// HasShadowTop - Flag indicating if PageBreadcrumb should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Flag indicating if PageBreadcrumb should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageBreadcrumb has a scrolling overflow. Optional.
	HasOverflowScroll bool
}
