package page

// from file "react-core/src/components/Page/PageSidebar.tsx"

type SidebarProps struct {
	// ClassName - Additional classes added to the page sidebar. Optional.
	ClassName string
	// Nav - Component to render the side navigation (e.g. <Nav />. Optional.
	Nav any // React.ReactNode 
	// IsManagedSidebar - If true, manages the sidebar open/close state and there is no need to pass the isNavOpen
	// boolean into the sidebar component or add a callback onNavToggle function into the PageHeader component.
	// Optional.
	IsManagedSidebar bool
	// IsNavOpen - Programmatically manage if the side nav is shown, if isManagedSidebar is set to true in the
	// Page component, this prop is managed. Optional.
	IsNavOpen bool
	// Theme - Indicates the color scheme of the sidebar. Optional.
	Theme any /* "dark" | "light" */
}

type SidebarContextProps struct {
	// IsNavOpen - 
	IsNavOpen bool
}
