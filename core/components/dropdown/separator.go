package dropdown

// from file "react-core/src/components/Dropdown/DropdownSeparator.tsx"

type SeparatorProps struct {
	// ClassName - Classes applied to root element of dropdown item. Optional.
	ClassName string
	// OnClick - Click event to pass to InternalDropdownItem. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
}
