package page

// from file "react-core/src/components/Page/Page.tsx"

type ContextProps struct {
	// IsManagedSidebar - 
	IsManagedSidebar bool
	// OnNavToggle - 
	OnNavToggle func()
	// IsNavOpen - 
	IsNavOpen bool
	// Width - 
	Width int
	// GetBreakpoint - 
	GetBreakpoint func(width any /* int | "" */) any /* "default" | "sm" | "md" | "lg" | "xl" | "2xl" */
}

type Props struct {
	// Children - Content rendered inside the main section of the page layout (e.g. <PageSection />). Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the page layout. Optional.
	ClassName string
	// Header - Header component (e.g. <PageHeader />). Optional.
	Header any // React.ReactNode 
	// Sidebar - Sidebar component for a side nav (e.g. <PageSidebar />). Optional.
	Sidebar any // React.ReactNode 
	// NotificationDrawer - Notification drawer component for an optional notification drawer (e.g. <NotificationDrawer
	// />). Optional.
	NotificationDrawer any // React.ReactNode 
	// IsNotificationDrawerExpanded - Flag indicating Notification drawer in expanded. Optional.
	IsNotificationDrawerExpanded bool
	// IsBreadcrumbWidthLimited - Flag indicating if breadcrumb width should be limited. Optional.
	IsBreadcrumbWidthLimited bool
	// OnNotificationDrawerExpand - Callback when notification drawer panel is finished expanding. Optional.
	OnNotificationDrawerExpand func()
	// SkipToContent - Skip to content component for the page. Optional.
	SkipToContent any // React.ReactElement 
	// Role - Sets the value for role on the <main> element. Optional.
	Role string
	// MainContainerId - an id to use for the [role="main"] element. Optional.
	MainContainerId string
	// MainTabIndex - tabIndex to use for the [role="main"] element, null to unset it. Optional.
	MainTabIndex any /* int | "" */
	// IsManagedSidebar - If true, manages the sidebar open/close state and there is no need to pass the isNavOpen
	// boolean into the sidebar component or add a callback onNavToggle function into the PageHeader component.
	// Optional.
	IsManagedSidebar bool
	// IsTertiaryNavWidthLimited - Flag indicating if tertiary nav width should be limited. Optional.
	IsTertiaryNavWidthLimited bool
	// DefaultManagedSidebarIsOpen - If true, the managed sidebar is initially open for desktop view. Optional.
	DefaultManagedSidebarIsOpen bool
	// OnPageResize - Can add callback to be notified when resize occurs, for example to set the sidebar isNav
	// prop to false for a width < 768px Returns object { mobileView: boolean, windowSize: number }. Optional.
	OnPageResize func(object any)
	// GetBreakpoint - The page resize observer uses the breakpoints returned from this function when adding
	// the pf-m-breakpoint-[default|sm|md|lg|xl|2xl] class You can override the default getBreakpoint function
	// to return breakpoints at different sizes than the default You can view the default getBreakpoint function
	// here: https://github.com/patternfly/patternfly-react/blob/main/packages/react-core/src/helpers/util.ts.
	// Optional.
	GetBreakpoint func(width any /* int | "" */) any /* "default" | "sm" | "md" | "lg" | "xl" | "2xl" */
	// Breadcrumb - Breadcrumb component for the page. Optional.
	Breadcrumb any // React.ReactNode 
	// TertiaryNav - Tertiary nav component for the page. Optional.
	TertiaryNav any // React.ReactNode 
	// MainAriaLabel - Accessible label, can be used to name main section. Optional.
	MainAriaLabel string
	// IsTertiaryNavGrouped - Flag indicating if the tertiaryNav should be in a group. Optional.
	IsTertiaryNavGrouped bool
	// IsBreadcrumbGrouped - Flag indicating if the breadcrumb should be in a group. Optional.
	IsBreadcrumbGrouped bool
	// AdditionalGroupedContent - Additional content of the group. Optional.
	AdditionalGroupedContent any // React.ReactNode 
	// GroupProps - Additional props of the group. Optional.
	GroupProps any // PageGroupProps 
}

type State struct {
	// DesktopIsNavOpen - 
	DesktopIsNavOpen bool
	// MobileIsNavOpen - 
	MobileIsNavOpen bool
	// MobileView - 
	MobileView bool
	// Width - 
	Width int
}
