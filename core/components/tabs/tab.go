package tabs

// from file "react-core/src/components/Tabs/Tab.tsx"

type TabProps struct {
	// Children - content rendered inside the Tab content area. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Tab. Optional.
	ClassName string
	// Href - URL associated with the Tab. A Tab with an href will render as an <a> instead of a <button>. A
	// Tab inside a <Tabs component="nav"> should have an href. Optional.
	Href string
	// Title - Content rendered in the tab title. Should be <TabTitleText> and/or <TabTitleIcon> for proper styling.
	Title any // React.ReactNode 
	// EventKey - uniquely identifies the tab.
	EventKey any /* int | string */
	// TabContentId - child id for case in which a TabContent section is defined outside of a Tabs component.
	// Optional.
	TabContentId any /* string | int */
	// TabContentRef - child reference for case in which a TabContent section is defined outside of a Tabs component.
	// Optional.
	TabContentRef any // React.RefObject 
	// IsHidden - whether to render the tab or not. Optional.
	IsHidden bool
	// IsDisabled - Adds disabled styling and disables the button using the disabled html attribute. Optional.
	IsDisabled bool
	// IsAriaDisabled - Adds disabled styling and communicates that the button is disabled using the aria-disabled
	// html attribute. Optional.
	IsAriaDisabled bool
	// InoperableEvents - Events to prevent when the button is in an aria-disabled state. Optional.
	InoperableEvents []string
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
	// Tooltip - Optional Tooltip rendered to a Tab. Should be <Tooltip> with appropriate props for proper rendering.
	// Optional.
	Tooltip any // React.ReactElement 
	// CloseButtonAriaLabel - Optional.
	CloseButtonAriaLabel string
	// IsCloseDisabled - Optional.
	IsCloseDisabled bool
}
