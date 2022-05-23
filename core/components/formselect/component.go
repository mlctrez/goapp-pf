package formselect

// from file "react-core/src/components/FormSelect/FormSelect.tsx"

type Props struct {
	// Children - content rendered inside the FormSelect.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the FormSelect control. Optional.
	ClassName string
	// Value - value of selected option. Optional.
	Value any
	// Validated - Value to indicate if the select is modified to show that validation state. If set to success,
	// select will be modified to indicate valid state. If set to error, select will be modified to indicate
	// error state. Optional.
	Validated any /* "success" | "warning" | "error" | "default" */
	// IsDisabled - Flag indicating the FormSelect is disabled. Optional.
	IsDisabled bool
	// IsRequired - Sets the FormSelect required. Optional.
	IsRequired bool
	// IsIconSprite - Use the external file instead of a data URI. Optional.
	IsIconSprite bool
	// OnBlur - Optional callback for updating when selection loses focus. Optional.
	OnBlur func(event any // React.FormEvent )
	// OnFocus - Optional callback for updating when selection gets focus. Optional.
	OnFocus func(event any // React.FormEvent )
	// OnChange - Optional callback for updating when selection changes. Optional.
	OnChange func(value string, event any // React.FormEvent )
	// AriaLabel - Custom flag to show that the FormSelect requires an associated id or aria-label. Optional.
	AriaLabel string
}
