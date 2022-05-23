package simplelist

// from file "react-core/src/components/SimpleList/SimpleListItem.tsx"

type ItemProps struct {
	// ItemId - id for the item. Optional.
	ItemId any /* int | string */
	// Children - Content rendered inside the SimpleList item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the SimpleList <li>. Optional.
	ClassName string
	// Component - Component type of the SimpleList item. Optional.
	Component any /* "button" | "a" */
	// ComponentClassName - Additional classes added to the SimpleList <a> or <button>. Optional.
	ComponentClassName string
	// ComponentProps - Additional props added to the SimpleList <a> or <button>. Optional.
	ComponentProps any
	// IsActive - Indicates if the link is current/highlighted. Optional.
	IsActive bool
	// IsCurrent - Optional.
	IsCurrent bool
	// OnClick - OnClick callback for the SimpleList item. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */)
	// Type - Type of button SimpleList item. Optional.
	Type any /* "button" | "submit" | "reset" */
	// Href - Default hyperlink location. Optional.
	Href string
}
