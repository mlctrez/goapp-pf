package simplelist

// from file "react-core/src/components/SimpleList/SimpleList.tsx"

type Props struct {
	// Children - Content rendered inside the SimpleList. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the SimpleList container. Optional.
	ClassName string
	// OnSelect - Callback when an item is selected. Optional.
	OnSelect func(ref any /* any // React.RefObject  | any // React.RefObject  */, props any // SimpleListItemProps )
	// IsControlled - Indicates whether component is controlled by its internal state. Optional.
	IsControlled bool
}

type State struct {
	// CurrentRef - Ref of the current SimpleListItem.
	CurrentRef any /* any // React.RefObject  | any // React.RefObject  */
}

type ContextProps struct {
	// CurrentRef - 
	CurrentRef any /* any // React.RefObject  | any // React.RefObject  */
	// UpdateCurrentRef - 
	UpdateCurrentRef func(id any /* any // React.RefObject  | any // React.RefObject  */, props any // SimpleListItemProps )
	// IsControlled - 
	IsControlled bool
}
