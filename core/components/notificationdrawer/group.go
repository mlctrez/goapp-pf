package notificationdrawer

// from file "react-core/src/components/NotificationDrawer/NotificationDrawerGroup.tsx"

type GroupProps struct {
	// Children - Content rendered inside the group. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the group. Optional.
	ClassName string
	// Count - Notification drawer group count.
	Count int
	// IsExpanded - Adds styling to the group to indicate expanded state.
	IsExpanded bool
	// IsRead - Adds styling to the group to indicate whether it has been read. Optional.
	IsRead bool
	// OnExpand - Callback for when group button is clicked to expand. Optional.
	OnExpand func(event any, value bool)
	// Title - Notification drawer group title.
	Title any /* string | any // React.ReactNode  */
	// TruncateTitle - Truncate title to number of lines. Optional.
	TruncateTitle int
	// TooltipPosition - Position of the tooltip which is displayed if text is truncated. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}
