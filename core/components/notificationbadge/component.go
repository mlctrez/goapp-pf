package notificationbadge

// from file "react-core/src/components/NotificationBadge/NotificationBadge.tsx"

type Props struct {
	// IsRead - Optional.
	IsRead bool
	// Variant - Determines the variant of the notification badge. Optional.
	Variant any /* any // NotificationBadgeVariant  | "read" | "unread" | "attention" */
	// Count - A number displayed in the badge alongside the icon. Optional.
	Count int
	// Children - content rendered inside the notification badge. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the notification badge. Optional.
	ClassName string
	// AriaLabel - Adds accessible text to the notification badge. Optional.
	AriaLabel string
	// AttentionIcon - Icon to display for attention variant. Optional.
	AttentionIcon any // React.ReactNode 
	// Icon - Icon do display in notification badge. Optional.
	Icon any // React.ReactNode 
}
