package helpertext

// from file "react-core/src/components/HelperText/HelperTextItem.tsx"

type ItemProps struct {
	// Children - Content rendered inside the helper text item. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the helper text item. Optional.
	ClassName string
	// Component - Sets the component type of the helper text item. Optional.
	Component any /* "div" | "li" */
	// Variant - Variant styling of the helper text item. Optional.
	Variant any /* "default" | "indeterminate" | "warning" | "success" | "error" */
	// Icon - Custom icon prefixing the helper text. This property will override the default icon paired with
	// each helper text variant. Optional.
	Icon any // React.ReactNode 
	// IsDynamic - Flag indicating the helper text item is dynamic. This prop should be used when the text content
	// of the helper text item will never change, but the icon and styling will be dynamically updated via the
	// `variant` prop. Optional.
	IsDynamic bool
	// HasIcon - Flag indicating the helper text should have an icon. Dynamic helper texts include icons by default
	// while static helper texts do not. Optional.
	HasIcon bool
	// Id - ID for the helper text item. The value of this prop can be passed into a form component's aria-describedby
	// prop when you intend for only specific helper text items to be announced to assistive technologies. Optional.
	Id string
}
