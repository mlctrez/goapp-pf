package drawer

// from file "react-core/src/components/Drawer/DrawerContent.tsx"

type ContentProps struct {
	// ClassName - Additional classes added to the Drawer. Optional.
	ClassName string
	// Children - Content to be rendered in the drawer. Optional.
	Children any // React.ReactNode 
	// PanelContent - Content rendered in the drawer panel.
	PanelContent any // React.ReactNode 
	// ColorVariant - Color variant of the background of the drawer panel. Optional.
	ColorVariant any /* any // DrawerColorVariant  | "light-200" | "default" */
}
