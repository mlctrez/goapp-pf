package tabs

// from file "react-core/src/components/Tabs/Tabs.tsx"

type Props struct {
	// Children - Content rendered inside the tabs component. Must be React.ReactElement<TabProps>[].
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the tabs. Optional.
	ClassName string
	// Variant - Tabs background color variant. Optional.
	Variant any /* "default" | "light300" */
	// ActiveKey - The index of the active tab. Optional.
	ActiveKey any /* int | string */
	// DefaultActiveKey - The index of the default active tab. Set this for uncontrolled Tabs. Optional.
	DefaultActiveKey any /* int | string */
	// OnSelect - Callback to handle tab selection. Optional.
	OnSelect func(event any // React.MouseEvent , eventKey any /* int | string */)
	// OnClose - Optional.
	OnClose func(event any // React.MouseEvent , eventKey any /* int | string */)
	// OnAdd - Optional.
	OnAdd func()
	// AddButtonAriaLabel - Optional.
	AddButtonAriaLabel string
	// Id - Uniquely identifies the tabs. Optional.
	Id string
	// IsFilled - Enables the filled tab list layout. Optional.
	IsFilled bool
	// IsSecondary - Enables secondary tab styling. Optional.
	IsSecondary bool
	// IsBox - Enables box styling to the tab component. Optional.
	IsBox bool
	// IsVertical - Enables vertical tab styling. Optional.
	IsVertical bool
	// HasBorderBottom - Enables border bottom tab styling on tabs. Defaults to true. To remove the bottom border,
	// set this prop to false. Optional.
	HasBorderBottom bool
	// HasSecondaryBorderBottom - Enables border bottom styling for secondary tabs. Optional.
	HasSecondaryBorderBottom bool
	// LeftScrollAriaLabel - Aria-label for the left scroll button. Optional.
	LeftScrollAriaLabel string
	// RightScrollAriaLabel - Aria-label for the right scroll button. Optional.
	RightScrollAriaLabel string
	// Component - Determines what tag is used around the tabs. Use "nav" to define the tabs inside a navigation
	// region. Optional.
	Component any /* "div" | "nav" */
	// AriaLabel - Provides an accessible label for the tabs. Labels should be unique for each set of tabs that
	// are present on a page. When component is set to nav, this prop should be defined to differentiate the
	// tabs from other navigation regions on the page. Optional.
	AriaLabel string
	// MountOnEnter - Waits until the first "enter" transition to mount tab children (add them to the DOM). Optional.
	MountOnEnter bool
	// UnmountOnExit - Unmounts tab children (removes them from the DOM) when they are no longer visible. Optional.
	UnmountOnExit bool
	// UsePageInsets - Flag indicates that the tabs should use page insets. Optional.
	UsePageInsets bool
	// Inset - Insets at various breakpoints. Optional.
	Inset map[string]string /* { default:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, sm:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, md:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, lg:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, xl:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl }, 2xl:{ insetNone, insetSm, insetMd, insetLg, insetXl, inset2xl } } */
	// Expandable - Enable expandable vertical tabs at various breakpoints. (isVertical should be set to true
	// for this to work). Optional.
	Expandable map[string]string /* { default:{ expandable, nonExpandable }, sm:{ expandable, nonExpandable }, md:{ expandable, nonExpandable }, lg:{ expandable, nonExpandable }, xl:{ expandable, nonExpandable }, 2xl:{ expandable, nonExpandable } } */
	// IsExpanded - Flag to indicate if the vertical tabs are expanded. Optional.
	IsExpanded bool
	// DefaultIsExpanded - Flag indicating the default expanded state for uncontrolled expand/collapse of. Optional.
	DefaultIsExpanded bool
	// ToggleText - Text that appears in the expandable toggle. Optional.
	ToggleText string
	// ToggleAriaLabel - Aria-label for the expandable toggle. Optional.
	ToggleAriaLabel string
	// OnToggle - Callback function to toggle the expandable tabs. Optional.
	OnToggle func(isExpanded bool)
}

type State struct {
	// ShowScrollButtons - 
	ShowScrollButtons bool
	// DisableLeftScrollButton - 
	DisableLeftScrollButton bool
	// DisableRightScrollButton - 
	DisableRightScrollButton bool
	// ShownKeys - 
	ShownKeys []( any /* string | int */ )
	// UncontrolledActiveKey - 
	UncontrolledActiveKey any /* int | string */
	// UncontrolledIsExpandedLocal - 
	UncontrolledIsExpandedLocal bool
	// OuiaStateId - 
	OuiaStateId string
}
