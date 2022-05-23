package page

// from file "react-core/src/components/Page/PageHeader.tsx"

type HeaderProps struct {
	// ClassName - Additional classes added to the page header. Optional.
	ClassName string
	// Logo - Component to render the logo/brand, use <Brand />. Optional.
	Logo any // React.ReactNode 
	// LogoProps - Additional props passed to the logo anchor container. Optional.
	LogoProps any /* TODO: how to pass additional properties */
	// LogoComponent - Component to use to wrap the passed <logo>. Optional.
	LogoComponent any // React.ReactNode 
	// HeaderTools - Component to render the header tools, use <PageHeaderTools />. Optional.
	HeaderTools any // React.ReactNode 
	// TopNav - Component to render navigation within the header, use <Nav />. Optional.
	TopNav any // React.ReactNode 
	// ShowNavToggle - True to show the nav toggle button (toggles side nav). Optional.
	ShowNavToggle bool
	// IsNavOpen - True if the side nav is shown. Optional.
	IsNavOpen bool
	// IsManagedSidebar - This prop is no longer managed through PageHeader but in the Page component. Optional.
	IsManagedSidebar bool
	// Role - Sets the value for role on the <main> element. Optional.
	Role string
	// OnNavToggle - Callback function to handle the side nav toggle button, managed by the Page component if
	// the Page isManagedSidebar prop is set to true. Optional.
	OnNavToggle func()
	// AriaLabel - Aria Label for the nav toggle button. Optional.
	AriaLabel string
}
