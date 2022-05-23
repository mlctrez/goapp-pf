package page

// from file "react-core/src/components/Page/PageSection.tsx"

type SectionProps struct {
	// Children - Content rendered inside the section. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the section. Optional.
	ClassName string
	// Variant - Section background color variant. Optional.
	Variant any /* "default" | "light" | "dark" | "darker" */
	// Type - Section type variant. Optional.
	Type any /* "default" | "nav" | "subnav" | "breadcrumb" | "tabs" | "wizard" */
	// IsFilled - Enables the page section to fill the available vertical space. Optional.
	IsFilled bool
	// IsWidthLimited - Limits the width of the section. Optional.
	IsWidthLimited bool
	// IsCenterAligned - Flag indicating if the section content is center aligned. isWidthLimited must be set
	// for this to work. Optional.
	IsCenterAligned bool
	// Padding - Padding at various breakpoints. Optional.
	Padding map[string]string /* { default:{ padding, noPadding }, sm:{ padding, noPadding }, md:{ padding, noPadding }, lg:{ padding, noPadding }, xl:{ padding, noPadding }, 2xl:{ padding, noPadding } } */
	// Sticky - Modifier indicating if PageSection is sticky to the top or bottom. Optional.
	Sticky any /* "top" | "bottom" */
	// HasShadowTop - Modifier indicating if PageSection should have a shadow at the top. Optional.
	HasShadowTop bool
	// HasShadowBottom - Modifier indicating if PageSection should have a shadow at the bottom. Optional.
	HasShadowBottom bool
	// HasOverflowScroll - Flag indicating if the PageSection has a scrolling overflow. Optional.
	HasOverflowScroll bool
}
