package duallistselector

// from file "react-core/src/components/DualListSelector/DualListSelectorListWrapper.tsx"

type ListWrapperProps struct {
	// ClassName - Additional classes applied to the dual list selector. Optional.
	ClassName string
	// Children - Anything that can be rendered inside of the list. Optional.
	Children any // React.ReactNode 
	// Id - Id of the dual list selector list. Optional.
	Id string
	// AriaLabelledby - Accessibly label for the list.
	AriaLabelledby string
	// InnerRef - Optional.
	InnerRef any // React.RefObject 
	// Options - Optional.
	Options []any // React.ReactNode 
	// SelectedOptions - Optional.
	SelectedOptions any /* []string | []int */
	// OnOptionSelect - Optional.
	OnOptionSelect func(e any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  */, index int, id string)
	// DisplayOption - Optional.
	DisplayOption func(option any // React.ReactNode ) bool
	// IsDisabled - Flag indicating whether the component is disabled. Optional.
	IsDisabled bool
}
