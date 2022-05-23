package pfselect

// from file "react-core/src/components/Select/SelectToggle.tsx"

type SelectToggleProps struct {
	// Id - HTML ID of dropdown toggle.
	Id string
	// Children - Anything which can be rendered as dropdown toggle.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of dropdown toggle. Optional.
	ClassName string
	// IsOpen - Flag to indicate if select is open. Optional.
	IsOpen bool
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(isExpanded bool, event any /* any // React.MouseEvent  | any // React.ChangeEvent  | any // React.KeyboardEvent  | any // Event  */)
	// OnEnter - Callback for toggle open on keyboard entry. Optional.
	OnEnter func()
	// OnClose - Callback for toggle close. Optional.
	OnClose func()
	// OnBlur - Callback for toggle blur. Optional.
	OnBlur func(event any)
	// HandleTypeaheadKeys - Optional.
	HandleTypeaheadKeys func(position string, shiftKey bool)
	// MoveFocusToLastMenuItem - Optional.
	MoveFocusToLastMenuItem func()
	// ParentRef - Element which wraps toggle.
	ParentRef any // React.RefObject 
	// MenuRef - The menu element. Optional.
	MenuRef any // React.RefObject 
	// FooterRef - The menu footer element. Optional.
	FooterRef any // React.RefObject 
	// IsActive - Forces active state. Optional.
	IsActive bool
	// IsPlain - Display the toggle with no border or background. Optional.
	IsPlain bool
	// IsDisabled - Flag indicating if select is disabled. Optional.
	IsDisabled bool
	// HasPlaceholderStyle - Flag indicating if placeholder styles should be applied. Optional.
	HasPlaceholderStyle bool
	// Type - Type of the toggle button, defaults to 'button'. Optional.
	Type any /* "reset" | "button" | "submit" | undefined */
	// AriaLabelledby - Id of label for the Select aria-labelledby. Optional.
	AriaLabelledby string
	// AriaLabel - Label for toggle of select variants. Optional.
	AriaLabel string
	// Variant - Flag for variant, determines toggle rules and interaction. Optional.
	Variant any /* "single" | "checkbox" | "typeahead" | "typeaheadmulti" */
	// HasClearButton - Flag indicating if select toggle has an clear button. Optional.
	HasClearButton bool
	// HasFooter - Flag indicating if select menu has a footer. Optional.
	HasFooter bool
	// OnClickTypeaheadToggleButton - Optional.
	OnClickTypeaheadToggleButton func()
}
