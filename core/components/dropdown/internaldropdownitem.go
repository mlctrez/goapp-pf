package dropdown

// from file "react-core/src/components/Dropdown/InternalDropdownItem.tsx"

type InternalDropdownItemProps struct {
	// Children - Anything which can be rendered as dropdown item. Optional.
	Children any // React.ReactNode 
	// StyleChildren - Whether to set className on component when React.isValidElement(component). Optional.
	StyleChildren bool
	// ClassName - Classes applied to root element of dropdown item. Optional.
	ClassName string
	// ListItemClassName - Class applied to list element. Optional.
	ListItemClassName string
	// Component - Indicates which component will be used as dropdown item. Will have className injected if React.isValidElement(component).
	// Optional.
	Component any // React.ReactNode 
	// Role - Role for the item. Optional.
	Role string
	// IsDisabled - Render dropdown item as disabled option. Optional.
	IsDisabled bool
	// IsAriaDisabled - Render dropdown item as aria disabled option. Optional.
	IsAriaDisabled bool
	// IsPlainText - Render dropdown item as a non-interactive item. Optional.
	IsPlainText bool
	// IsHovered - Forces display of the hover state of the element. Optional.
	IsHovered bool
	// Href - Default hyperlink location. Optional.
	Href string
	// Tooltip - Tooltip to display when hovered over the item. Optional.
	Tooltip any // React.ReactNode 
	// TooltipProps - Additional tooltip props forwarded to the Tooltip component. Optional.
	TooltipProps any
	// Index - Optional.
	Index int
	// Context - Optional.
	Context map[string]string /* { "keyHandler":func(index int, innerIndex int, direction string), "sendRef":func(index int, ref any, isDisabled bool, isSeparator bool) } */
	// OnClick - Callback for click event. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
	// Id - ID for the list element. Optional.
	Id string
	// ComponentID - ID for the component element. Optional.
	ComponentID string
	// AdditionalChild - Additional content to include alongside item within the <li>. Optional.
	AdditionalChild any // React.ReactNode 
	// CustomChild - Custom item rendering that receives the DropdownContext. Optional.
	CustomChild any // React.ReactNode 
	// EnterTriggersArrowDown - Flag indicating if hitting enter on an item also triggers an arrow down key press.
	// Optional.
	EnterTriggersArrowDown bool
	// Icon - An image to display within the InternalDropdownItem, appearing before any component children. Optional.
	Icon any // React.ReactNode 
	// AutoFocus - Initial focus on the item when the menu is opened (Note: Only applicable to one of the items).
	// Optional.
	AutoFocus bool
	// Description - A short description of the dropdown item, displayed under the dropdown item content. Optional.
	Description any // React.ReactNode 
	// InoperableEvents - Events to prevent when the item is disabled. Optional.
	InoperableEvents []string
}
