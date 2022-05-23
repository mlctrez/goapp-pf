package backtotop

// from file "react-core/src/components/BackToTop/BackToTop.tsx"

type Props struct {
	// ClassName - Additional classes added to the back to top. Optional.
	ClassName string
	// Title - Title to appear in back to top button. Optional.
	Title string
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
	// ScrollableSelector - Selector for the scrollable element to spy on. Not passing a selector defaults to
	// spying on window scroll events. Optional.
	ScrollableSelector string
	// IsAlwaysVisible - Flag to always show back to top button, defaults to false. Optional.
	IsAlwaysVisible bool
}
