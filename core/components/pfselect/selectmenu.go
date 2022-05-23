package pfselect

// from file "react-core/src/components/Select/SelectMenu.tsx"

type SelectMenuProps struct {
	// Children - Content rendered inside the SelectMenu.
	Children any /* []any // React.ReactElement  | any // React.ReactNode  */
	// IsCustomContent - Flag indicating that the children is custom content to render inside the SelectMenu.
	// If true, variant prop is ignored. Optional.
	IsCustomContent bool
	// ClassName - Additional classes added to the SelectMenu control. Optional.
	ClassName string
	// IsExpanded - Flag indicating the Select is expanded. Optional.
	IsExpanded bool
	// IsGrouped - Flag indicating the Select options are grouped. Optional.
	IsGrouped bool
	// Selected - Currently selected option (for single, typeahead variants). Optional.
	Selected any /* string | any // SelectOptionObject  | []( any /* string | any // SelectOptionObject  */ ) */
	// Checked - Currently checked options (for checkbox variant). Optional.
	Checked []( any /* string | any // SelectOptionObject  */ )
	// OpenedOnEnter - Optional.
	OpenedOnEnter bool
	// MaxHeight - Flag to specify the  maximum height of the menu, as a string percentage or number of pixels.
	// Optional.
	MaxHeight any /* string | int */
	// Position - Indicates where menu will be alligned horizontally. Optional.
	Position any /* any // SelectPosition  | "right" | "left" */
	// NoResultsFoundText - Inner prop passed from parent. Optional.
	NoResultsFoundText string
	// CreateText - Inner prop passed from parent. Optional.
	CreateText string
	// SendRef - Optional.
	SendRef func(ref any // React.ReactNode , favoriteRef any // React.ReactNode , index int)
	// KeyHandler - Optional.
	KeyHandler func(index int, innerIndex int, position string)
	// HasInlineFilter - Flag indicating select has an inline text input for filtering. Optional.
	HasInlineFilter bool
	// InnerRef - Optional.
	InnerRef any
	// Footer - Content rendered in the footer of the select menu. Optional.
	Footer any // React.ReactNode 
	// FooterRef - The menu footer element. Optional.
	FooterRef any // React.RefObject 
	// IsLastOptionBeforeFooter - Optional.
	IsLastOptionBeforeFooter func(index int)
}
