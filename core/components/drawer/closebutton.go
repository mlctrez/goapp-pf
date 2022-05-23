package drawer

// from file "react-core/src/components/Drawer/DrawerCloseButton.tsx"

type CloseButtonProps struct {
	// ClassName - Additional classes added to the drawer close button outer <div>. Optional.
	ClassName string
	// OnClose - A callback for when the close button is clicked. Optional.
	OnClose func()
	// AriaLabel - Accessible label for the drawer close button. Optional.
	AriaLabel string
}
