package notificationdrawer

// from file "react-core/src/components/NotificationDrawer/NotificationDrawerHeader.tsx"

type HeaderProps struct {
	// Children - Content rendered inside the drawer. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes for notification drawer header. Optional.
	ClassName string
	// CloseButtonAriaLabel - Adds custom accessible text to the notification drawer close button. Optional.
	CloseButtonAriaLabel string
	// Count - Notification drawer heading count. Optional.
	Count int
	// CustomText - Notification drawer heading custom text which can be used instead of providing count/unreadText.
	// Optional.
	CustomText string
	// OnClose - Callback for when close button is clicked. Optional.
	OnClose func()
	// Title - Notification drawer heading title. Optional.
	Title string
	// UnreadText - Notification drawer heading unread text used in combination with a count. Optional.
	UnreadText string
}
