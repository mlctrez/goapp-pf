package menu

// from file "react-core/src/components/Menu/Menu.tsx"

type Props struct {
	// Children - Anything that can be rendered inside of the Menu. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the Menu. Optional.
	ClassName string
	// Id - ID of the menu. Optional.
	Id string
	// OnSelect - Callback for updating when item selection changes. You can also specify onClick on the MenuItem.
	// Optional.
	OnSelect func(event any // React.MouseEvent , itemId any /* string | int */)
	// Selected - Single itemId for single select menus, or array of itemIds for multi select. You can also specify
	// isSelected on the MenuItem. Optional.
	Selected any /* any | []any */
	// OnActionClick - Callback called when an MenuItems's action button is clicked. You can also specify it
	// within a MenuItemAction. Optional.
	OnActionClick func(event any, itemId any, actionId any)
	// HasSearchInput - Search input of menu. Optional.
	HasSearchInput bool
	// OnSearchInputChange - A callback for when the input value changes. Optional.
	OnSearchInputChange func(event any /* any // React.FormEvent  | any // React.SyntheticEvent  */, value string)
	// AriaLabel - Accessibility label. Optional.
	AriaLabel string
	// ContainsFlyout - Optional.
	ContainsFlyout bool
	// IsNavFlyout - Optional.
	IsNavFlyout bool
	// ContainsDrilldown - Optional.
	ContainsDrilldown bool
	// IsMenuDrilledIn - Optional.
	IsMenuDrilledIn bool
	// DrilldownItemPath - Optional.
	DrilldownItemPath []string
	// DrilledInMenus - Optional.
	DrilledInMenus []string
	// OnDrillIn - Optional.
	OnDrillIn func(fromItemId string, toItemId string, itemId string)
	// OnDrillOut - Optional.
	OnDrillOut func(toItemId string, itemId string)
	// OnGetMenuHeight - Optional.
	OnGetMenuHeight func(menuId string, height int)
	// ParentMenu - Optional.
	ParentMenu string
	// ActiveMenu - Optional.
	ActiveMenu string
	// ActiveItemId - Optional.
	ActiveItemId any /* string | int */
	// InnerRef - Optional.
	InnerRef any // React.Ref 
	// IsRootMenu - Internal flag indicating if the Menu is the root of a menu tree. Optional.
	IsRootMenu bool
	// IsPlain - Indicates if the menu should be without the outer box-shadow. Optional.
	IsPlain bool
	// IsScrollable - Indicates if the menu should be srollable. Optional.
	IsScrollable bool
}

type State struct {
	// SearchInputValue - 
	SearchInputValue any /* string | "" */
	// OuiaStateId - 
	OuiaStateId string
	// TransitionMoveTarget - 
	TransitionMoveTarget any // HTMLElement 
	// FlyoutRef - 
	FlyoutRef any /* any // React.Ref  | "" */
	// DisableHover - 
	DisableHover bool
}
