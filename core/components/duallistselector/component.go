package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelector.tsx"

type Props struct {
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// Id - Id of the dual list selector. Optional.
	Id string
	// IsTree - Flag indicating if the dual list selector uses trees instead of simple lists. Optional.
	IsTree bool
	// IsDisabled - Flag indicating if the dual list selector is in a disabled state. Optional.
	IsDisabled bool
	// Children - Content to be rendered in the dual list selector. Panes & controls will not be built dynamically
	// when children are provided. Optional.
	Children any // React.ReactNode 
	// AvailableOptionsTitle - Title applied to the dynamically built available options pane. Optional.
	AvailableOptionsTitle string
	// AvailableOptions - Options to display in the dynamically built available options pane. When using trees,
	// the options should be in the DualListSelectorTreeItemData[] format. Optional.
	AvailableOptions any /* []any // React.ReactNode  | []any // DualListSelectorTreeItemData  */
	// AvailableOptionsStatus - Status message to display above the dynamically built available options pane.
	// Optional.
	AvailableOptionsStatus string
	// AvailableOptionsActions - Actions to be displayed above the dynamically built available options pane.
	// Optional.
	AvailableOptionsActions []any // React.ReactNode 
	// ChosenOptionsTitle - Title applied to the dynamically built chosen options pane. Optional.
	ChosenOptionsTitle string
	// ChosenOptions - Options to display in the dynamically built chosen options pane. When using trees, the
	// options should be in the DualListSelectorTreeItemData[] format. Optional.
	ChosenOptions any /* []any // React.ReactNode  | []any // DualListSelectorTreeItemData  */
	// ChosenOptionsStatus - Status message to display above the dynamically built chosen options pane. Optional.
	ChosenOptionsStatus string
	// ChosenOptionsActions - Actions to be displayed above the dynamically built chosen options pane. Optional.
	ChosenOptionsActions []any // React.ReactNode 
	// ControlsAriaLabel - Accessible label for the dynamically built controls between the two panes. Optional.
	ControlsAriaLabel string
	// AddSelected - Optional callback for the dynamically built add selected button. Optional.
	AddSelected func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddSelectedAriaLabel - Accessible label for the dynamically built add selected button. Optional.
	AddSelectedAriaLabel string
	// AddSelectedTooltip - Tooltip content for the dynamically built add selected button. Optional.
	AddSelectedTooltip any // React.ReactNode 
	// AddSelectedTooltipProps - Additonal tooltip properties for the dynamically built add selected tooltip.
	// Optional.
	AddSelectedTooltipProps any
	// OnListChange - Callback fired every time dynamically built options are chosen or removed. Optional.
	OnListChange func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddAll - Optional callback for the dynamically built add all button. Optional.
	AddAll func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// AddAllAriaLabel - Accessible label for the dynamically built add all button. Optional.
	AddAllAriaLabel string
	// AddAllTooltip - Tooltip content for the dynamically built add all button. Optional.
	AddAllTooltip any // React.ReactNode 
	// AddAllTooltipProps - Additonal tooltip properties for the dynamically built add all tooltip. Optional.
	AddAllTooltipProps any
	// RemoveSelected - Optional callback for the dynamically built remove selected button. Optional.
	RemoveSelected func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// RemoveSelectedAriaLabel - Accessible label for the dynamically built remove selected button. Optional.
	RemoveSelectedAriaLabel string
	// RemoveSelectedTooltip - Tooltip content for the dynamically built remove selected button. Optional.
	RemoveSelectedTooltip any // React.ReactNode 
	// RemoveSelectedTooltipProps - Additonal tooltip properties for the dynamically built remove selected tooltip.
	// Optional.
	RemoveSelectedTooltipProps any
	// RemoveAll - Optional callback for the dynamically built remove all button. Optional.
	RemoveAll func(newAvailableOptions []any // React.ReactNode , newChosenOptions []any // React.ReactNode )
	// RemoveAllAriaLabel - Accessible label for the dynamically built remove all button. Optional.
	RemoveAllAriaLabel string
	// RemoveAllTooltip - Tooltip content for the dynamically built remove all button. Optional.
	RemoveAllTooltip any // React.ReactNode 
	// RemoveAllTooltipProps - Additonal tooltip properties for the dynamically built remove all tooltip. Optional.
	RemoveAllTooltipProps any
	// OnOptionSelect - Optional callback fired when a dynamically built option is selected. Optional.
	OnOptionSelect func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, index int, isChosen bool, id string, itemData any, parentData any)
	// OnOptionCheck - Optional callback fired when a dynamically built option is checked. Optional.
	OnOptionCheck func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, checked bool, checkedId string, newCheckedItems []string)
	// IsSearchable - Flag indicating a search bar should be included above both the dynamically built available
	// and chosen panes. Optional.
	IsSearchable bool
	// AvailableOptionsSearchAriaLabel - Accessible label for the search input on the dynamically built available
	// options pane. Optional.
	AvailableOptionsSearchAriaLabel string
	// OnAvailableOptionsSearchInputChanged - A callback for when the search input value for the dynamically
	// built available options changes. Optional.
	OnAvailableOptionsSearchInputChanged func(value string, event any // React.FormEvent )
	// ChosenOptionsSearchAriaLabel - Accessible label for the search input on the dynamically built chosen options
	// pane. Optional.
	ChosenOptionsSearchAriaLabel string
	// OnChosenOptionsSearchInputChanged - A callback for when the search input value for the dynamically built
	// chosen options changes. Optional.
	OnChosenOptionsSearchInputChanged func(value string, event any // React.FormEvent )
	// FilterOption - Optional filter function for custom filtering based on search string. Used with a dynamically
	// built search input. Optional.
	FilterOption func(option any // React.ReactNode , input string) bool
}

type State struct {
	// AvailableOptions - 
	AvailableOptions []any // React.ReactNode 
	// AvailableOptionsSelected - 
	AvailableOptionsSelected []int
	// AvailableFilteredOptions - 
	AvailableFilteredOptions []any // React.ReactNode 
	// ChosenOptions - 
	ChosenOptions []any // React.ReactNode 
	// ChosenOptionsSelected - 
	ChosenOptionsSelected []int
	// ChosenFilteredOptions - 
	ChosenFilteredOptions []any // React.ReactNode 
	// AvailableTreeFilteredOptions - 
	AvailableTreeFilteredOptions []string
	// AvailableTreeOptionsChecked - 
	AvailableTreeOptionsChecked []string
	// ChosenTreeOptionsChecked - 
	ChosenTreeOptionsChecked []string
	// ChosenTreeFilteredOptions - 
	ChosenTreeFilteredOptions []string
}
