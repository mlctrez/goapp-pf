package nav

// from file "react-core/src/components/Nav/NavItem.tsx"

type ItemProps struct {
	// Children - Content rendered inside the nav item. Optional.
	Children any // React.ReactNode 
	// StyleChildren - Whether to set className on children when React.isValidElement(children). Optional.
	StyleChildren bool
	// ClassName - Additional classes added to the nav item. Optional.
	ClassName string
	// To - Target navigation link. Optional.
	To string
	// IsActive - Flag indicating whether the item is active. Optional.
	IsActive bool
	// GroupId - Group identifier, will be returned with the onToggle and onSelect callback passed to the Nav
	// component. Optional.
	GroupId any /* string | int | "" */
	// ItemId - Item identifier, will be returned with the onToggle and onSelect callback passed to the Nav component.
	// Optional.
	ItemId any /* string | int | "" */
	// PreventDefault - If true prevents the default anchor link action to occur. Set to true if you want to
	// handle navigation yourself. Optional.
	PreventDefault bool
	// OnClick - Callback for item click. Optional.
	OnClick any // NavSelectClickHandler 
	// Component - Component used to render NavItems if  React.isValidElement(children) is false. Optional.
	Component any // React.ReactNode 
	// Flyout - Flyout of a nav item. This should be a Menu component. Optional.
	Flyout any // React.ReactElement 
	// OnShowFlyout - Callback when flyout is opened or closed. Optional.
	OnShowFlyout func()
}
