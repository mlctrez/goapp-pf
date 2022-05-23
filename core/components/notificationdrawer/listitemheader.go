package notificationdrawer

// from file "react-core/src/components/NotificationDrawer/NotificationDrawerListItemHeader.tsx"

type ListItemHeaderProps struct {
	// Children - Actions rendered inside the notification drawer list item header. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes for notification drawer list item header. Optional.
	ClassName string
	// Icon - Add custom icon for notification drawer list item header. Optional.
	Icon any // React.ReactNode 
	// SrTitle - Notification drawer list item header screen reader title. Optional.
	SrTitle string
	// Title - Notification drawer list item title.
	Title string
	// Variant - Variant indicates the severity level. Optional.
	Variant any /* "success" | "danger" | "warning" | "info" | "default" */
	// TruncateTitle - Truncate title to number of lines. Optional.
	TruncateTitle int
	// TooltipPosition - Position of the tooltip which is displayed if text is truncated. Optional.
	TooltipPosition any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
}
