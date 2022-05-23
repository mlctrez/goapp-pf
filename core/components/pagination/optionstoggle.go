package pagination

// from file "react-core/src/components/Pagination/OptionsToggle.tsx"

type OptionsToggleProps struct {
	// ItemsTitle - The type or title of the items being paginated. Optional.
	ItemsTitle string
	// OptionsToggle - Accessible label for the options toggle. Optional.
	OptionsToggle string
	// ItemsPerPageTitle - The title of the pagination options menu. Optional.
	ItemsPerPageTitle string
	// FirstIndex - The first index of the items being paginated. Optional.
	FirstIndex int
	// LastIndex - The last index of the items being paginated. Optional.
	LastIndex int
	// ItemCount - The total number of items being paginated. Optional.
	ItemCount int
	// WidgetId - Id added to the title of the pagination options menu. Optional.
	WidgetId string
	// ShowToggle - showToggle. Optional.
	ShowToggle bool
	// OnToggle - Event function that fires when user clicks the options menu toggle. Optional.
	OnToggle func(isOpen bool)
	// IsOpen - Flag indicating if the options menu dropdown is open or not. Optional.
	IsOpen bool
	// IsDisabled - Flag indicating if the options menu is disabled. Optional.
	IsDisabled bool
	// ParentRef - Optional.
	ParentRef any // HTMLElement 
	// ToggleTemplate - This will be shown in pagination toggle span. You can use firstIndex, lastIndex, itemCount,
	// itemsTitle props. Optional.
	ToggleTemplate any /* func(props any // ToggleTemplateProps ) any // React.ReactElement  | string */
	// OnEnter - Callback for toggle open on keyboard entry. Optional.
	OnEnter func()
	// OfWord - Label for the English word "of". Optional.
	OfWord string
	// PerPageComponent - Component to be used for wrapping the toggle contents. Use 'button' when you want all
	// of the toggle text to be clickable. Optional.
	PerPageComponent any /* "div" | "button" */
}
