package drawer

// from file "react-core/src/components/Drawer/Drawer.tsx"

type Props struct {
	// ClassName - Additional classes added to the Drawer. Optional.
	ClassName string
	// Children - Content rendered in the left hand panel. Optional.
	Children any // React.ReactNode 
	// IsExpanded - Indicates if the drawer is expanded. Optional.
	IsExpanded bool
	// IsInline - Indicates if the content element and panel element are displayed side by side. Optional.
	IsInline bool
	// IsStatic - Indicates if the drawer will always show both content and panel. Optional.
	IsStatic bool
	// Position - Position of the drawer panel. Optional.
	Position any /* "left" | "right" | "bottom" */
	// OnExpand - Callback when drawer panel is expanded after waiting 250ms for animation to complete. Optional.
	OnExpand func()
}

type ContextProps struct {
	// IsExpanded - 
	IsExpanded bool
	// IsStatic - 
	IsStatic bool
	// OnExpand - Optional.
	OnExpand func()
	// Position - Optional.
	Position string
	// DrawerRef - Optional.
	DrawerRef any // React.RefObject 
	// IsInline - 
	IsInline bool
}
