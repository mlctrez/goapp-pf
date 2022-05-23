package sidebar

// from file "react-core/src/components/Sidebar/SidebarPanel.tsx"

type PanelProps struct {
	// Children - 
	Children any // React.ReactNode 
	// Variant - Indicates whether the panel is positioned statically or sticky to the top. Default is sticky
	// on small screens when the orientation is stack, and static on medium and above screens when the orientation
	// is split. Optional.
	Variant any /* "default" | "sticky" | "static" */
	// HasNoBackground - Removes the background color. Optional.
	HasNoBackground bool
	// Width - Sets the panel width at various breakpoints. Default is 250px when the orientation is split. Optional.
	Width map[string]string /* { default:{ default, width_25, width_33, width_50, width_66, width_75, width_100 }, sm:{ default, width_25, width_33, width_50, width_66, width_75, width_100 }, md:{ default, width_25, width_33, width_50, width_66, width_75, width_100 }, lg:{ default, width_25, width_33, width_50, width_66, width_75, width_100 }, xl:{ default, width_25, width_33, width_50, width_66, width_75, width_100 }, 2xl:{ default, width_25, width_33, width_50, width_66, width_75, width_100 } } */
}
