package notificationdrawer

// from file "react-core/src/components/NotificationDrawer/NotificationDrawerListItem.tsx"

type ListItemProps struct {
	// Children - Content rendered inside the list item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the list item. Optional.
	ClassName string
	// IsHoverable - Modifies the list item to include hover styles on :hover. Optional.
	IsHoverable bool
	// IsRead - Adds styling to the list item to indicate it has been read. Optional.
	IsRead bool
	// OnClick - Callback for when a list item is clicked. Optional.
	OnClick func(event any)
	// TabIndex - Tab index for the list item. Optional.
	TabIndex int
	// Variant - Variant indicates the severity level. Optional.
	Variant any /* "default" | "success" | "danger" | "warning" | "info" */
}
