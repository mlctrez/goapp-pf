package dropdown

// from file "react-core/src/components/Dropdown/DropdownItem.tsx"

type ItemProps struct {
	// Children - Anything which can be rendered as dropdown item. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of dropdown item. Optional.
	ClassName string
	// ListItemClassName - Class to be applied to list item. Optional.
	ListItemClassName string
	// Component - A ReactElement to render, or a string to use as the component tag. Example: component={<Link
	// to="/components/alert/">Alert</Link>} Example: component="button" If React.isValidElement(component) the
	// className prop will be injected unless styleChildren="false". Optional.
	Component any // React.ReactNode 
	// ComponentID - ID for the component element. Optional.
	ComponentID string
	// StyleChildren - Whether to set className on component when React.isValidElement(component). Optional.
	StyleChildren bool
	// IsDisabled - Render dropdown item as disabled option. Optional.
	IsDisabled bool
	// IsAriaDisabled - Render dropdown item as aria-disabled option. Optional.
	IsAriaDisabled bool
	// IsPlainText - Render dropdown item as non-interactive item. Optional.
	IsPlainText bool
	// IsHovered - Forces display of the hover state of the element. Optional.
	IsHovered bool
	// Href - Default hyperlink location. Optional.
	Href string
	// Tooltip - Tooltip to display when hovered over the item. Optional.
	Tooltip any // React.ReactNode 
	// TooltipProps - Additional tooltip props forwarded to the Tooltip component. Optional.
	TooltipProps any
	// AdditionalChild - Additional node to include alongside item within the <li>. Optional.
	AdditionalChild any // React.ReactNode 
	// CustomChild - Custom item rendering that receives the DropdownContext. Optional.
	CustomChild any // React.ReactNode 
	// TabIndex - tabIndex to use, null to unset it. Optional.
	TabIndex any /* int | "" */
	// Icon - An image to display within the DropdownItem, appearing before any component children. Optional.
	Icon any // React.ReactNode 
	// AutoFocus - Initial focus on the item when the menu is opened (Note: Only applicable to one of the items).
	// Optional.
	AutoFocus bool
	// Description - A short description of the dropdown item, displayed under the dropdown item content. Optional.
	Description any // React.ReactNode 
}
