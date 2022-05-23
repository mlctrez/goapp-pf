package jumplinks

// from file "react-core/src/components/JumpLinks/JumpLinks.tsx"

type Props struct {
	// IsCentered - Whether to center children. Optional.
	IsCentered bool
	// IsVertical - Whether the layout of children is vertical or horizontal. Optional.
	IsVertical bool
	// Label - Label to add to nav element. Optional.
	Label any // React.ReactNode 
	// AlwaysShowLabel - Flag to always show the label when using `expandable`. Optional.
	AlwaysShowLabel bool
	// AriaLabel - Aria-label to add to nav element. Defaults to label. Optional.
	AriaLabel string
	// ScrollableSelector - Selector for the scrollable element to spy on. Not passing a selector disables spying.
	// Optional.
	ScrollableSelector string
	// ActiveIndex - The index of the child Jump link to make active. Optional.
	ActiveIndex int
	// Children - Children nodes. Optional.
	Children any // React.ReactNode 
	// Offset - Offset to add to `scrollPosition`, potentially for a masthead which content scrolls under. Optional.
	Offset int
	// Expandable - When to collapse/expand at different breakpoints. Optional.
	Expandable map[string]string /* { default:{ expandable, nonExpandable }, sm:{ expandable, nonExpandable }, md:{ expandable, nonExpandable }, lg:{ expandable, nonExpandable }, xl:{ expandable, nonExpandable }, 2xl:{ expandable, nonExpandable } } */
	// IsExpanded - On mobile whether or not the JumpLinks starts out expanded. Optional.
	IsExpanded bool
	// ToggleAriaLabel - Aria label for expandable toggle. Optional.
	ToggleAriaLabel string
	// ClassName - Class for nav. Optional.
	ClassName string
}
