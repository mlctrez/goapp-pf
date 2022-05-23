package menu

// from file "react-core/src/components/Menu/MenuItemAction.tsx"

type ItemActionProps struct {
	// ClassName - Additional classes added to the action button. Optional.
	ClassName string
	// Icon - The action icon to use. Optional.
	Icon any /* "favorites" | any // React.ReactNode  */
	// OnClick - Callback on action click, can also specify onActionClick on the Menu instead. Optional.
	OnClick func(event any)
	// AriaLabel - Accessibility label. Optional.
	AriaLabel string
	// IsFavorited - Flag indicating if the item is favorited. Optional.
	IsFavorited bool
	// IsDisabled - Disables action, can also be specified on the MenuItem instead. Optional.
	IsDisabled bool
	// ActionId - Identifies the action item in the onActionClick on the Menu. Optional.
	ActionId any
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
}
