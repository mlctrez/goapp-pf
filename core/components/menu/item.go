package menu

// from file "react-core/src/components/Menu/MenuItem.tsx"

type ItemProps struct {
	// Children - Content rendered inside the menu list item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the menu list item. Optional.
	ClassName string
	// ItemId - Identifies the component in the Menu onSelect or onActionClick callback. Optional.
	ItemId any
	// To - Target navigation link. Optional.
	To string
	// IsActive - Flag indicating whether the item is active. Optional.
	IsActive bool
	// IsFavorited - Flag indicating if the item is favorited. Optional.
	IsFavorited bool
	// IsLoadButton - Flag indicating if the item causes a load. Optional.
	IsLoadButton bool
	// IsLoading - Flag indicating a loading state. Optional.
	IsLoading bool
	// OnClick - Callback for item click. Optional.
	OnClick func(event any)
	// Component - Component used to render the menu item. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
	// IsDisabled - Render item as disabled option. Optional.
	IsDisabled bool
	// Icon - Render item with icon. Optional.
	Icon any // React.ReactNode 
	// Actions - Render item with one or more actions. Optional.
	Actions any // React.ReactNode 
	// Description - Description of the menu item. Optional.
	Description any // React.ReactNode 
	// IsExternalLink - Render external link icon. Optional.
	IsExternalLink bool
	// IsSelected - Flag indicating if the option is selected. Optional.
	IsSelected bool
	// FlyoutMenu - Optional.
	FlyoutMenu any // React.ReactElement 
	// OnShowFlyout - Optional.
	OnShowFlyout func(event any)
	// DrilldownMenu - Optional.
	DrilldownMenu any /* any // React.ReactNode  | func() any // React.ReactNode  */
	// Direction - Optional.
	Direction any /* "down" | "up" */
	// IsOnPath - Optional.
	IsOnPath bool
	// AriaLabel - Accessibility label. Optional.
	AriaLabel string
	// InnerRef - Optional.
	InnerRef any // React.Ref 
}
