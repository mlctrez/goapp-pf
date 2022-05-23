package textinputgroup

// from file "react-core/src/components/TextInputGroup/TextInputGroupMain.tsx"

type MainProps struct {
	// Children - Content rendered inside the text input group main div. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the text input group main container. Optional.
	ClassName string
	// Icon - Icon to be shown on the left side of the text input group main container. Optional.
	Icon any // React.ReactNode 
	// Type - Type that the input accepts. Optional.
	Type any /* "text" | "date" | "datetime-local" | "email" | "month" | "number" | "password" | "search" | "tel" | "time" | "url" */
	// Hint - Suggestion that will show up like a placeholder even with text in the input. Optional.
	Hint string
	// OnChange - Callback for when there is a change in the input field. Optional.
	OnChange func(value string, event any // React.FormEvent )
	// OnFocus - Callback for when the input field is focused. Optional.
	OnFocus func(event any)
	// OnBlur - Callback for when focus is lost on the input field. Optional.
	OnBlur func(event any)
	// AriaLabel - Accessibility label for the input. Optional.
	AriaLabel string
	// Value - Value for the input. Optional.
	Value any /* string | int */
	// Placeholder - Placeholder value for the input. Optional.
	Placeholder string
}
