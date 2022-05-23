package overflowmenu

// from file "react-core/src/components/OverflowMenu/OverflowMenu.tsx"

type Props struct {
	// Children - Any elements that can be rendered in the menu. Optional.
	Children any
	// ClassName - Additional classes added to the OverflowMenu. Optional.
	ClassName string
	// Breakpoint - Indicates breakpoint at which to switch between horizontal menu and vertical dropdown.
	Breakpoint any /* "md" | "lg" | "xl" | "2xl" */
}

type State struct {
	// IsBelowBreakpoint - 
	IsBelowBreakpoint bool
}
