package drawer

// from file "react-core/src/components/Drawer/DrawerPanelContent.tsx"

type PanelContentProps struct {
	// ClassName - Additional classes added to the drawer. Optional.
	ClassName string
	// Id - ID of the drawer panel. Optional.
	Id string
	// Children - Content to be rendered in the drawer panel. Optional.
	Children any // React.ReactNode 
	// HasNoBorder - Flag indicating that the drawer panel should not have a border. Optional.
	HasNoBorder bool
	// IsResizable - Flag indicating that the drawer panel should be resizable. Optional.
	IsResizable bool
	// OnResize - Callback for resize end. Optional.
	OnResize func(width int, id string)
	// MinSize - The minimum size of a drawer, in either pixels or percentage. Optional.
	MinSize string
	// DefaultSize - The starting size of a resizable drawer, in either pixels or percentage. Optional.
	DefaultSize string
	// MaxSize - The maximum size of a drawer, in either pixels or percentage. Optional.
	MaxSize string
	// Increment - The increment amount for keyboard drawer resizing, in pixels. Optional.
	Increment int
	// ResizeAriaLabel - Aria label for the resizable drawer splitter. Optional.
	ResizeAriaLabel string
	// Widths - Width for drawer panel at various breakpoints. Overriden by resizable drawer minSize and defaultSize.
	// Optional.
	Widths map[string]string /* { default:{ width_25, width_33, width_50, width_66, width_75, width_100 }, lg:{ width_25, width_33, width_50, width_66, width_75, width_100 }, xl:{ width_25, width_33, width_50, width_66, width_75, width_100 }, 2xl:{ width_25, width_33, width_50, width_66, width_75, width_100 } } */
	// ColorVariant - Color variant of the background of the drawer panel. Optional.
	ColorVariant any /* any // DrawerColorVariant  | "light-200" | "default" */
	// DrawerContentRef - Optional.
	DrawerContentRef any // React.RefObject 
}
