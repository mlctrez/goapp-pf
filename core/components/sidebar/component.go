package sidebar

// from file "react-core/src/components/Sidebar/Sidebar.tsx"

type Props struct {
	// Children - Optional.
	Children any // React.ReactNode 
	// Orientation - Indicates the direction of the layout. Default orientation is stack on small screens, and
	// split on medium screens and above. Optional.
	Orientation any /* "stack" | "split" */
	// IsPanelRight - Indicates that the panel is displayed to the right of the content when the oritentation
	// is split. Optional.
	IsPanelRight bool
	// HasGutter - Adds space between the panel and content. Optional.
	HasGutter bool
	// HasNoBackground - Removes the background color. Optional.
	HasNoBackground bool
}
