package textinput

// from file "react-core/src/components/TextInput/TextInput.tsx"

type Props struct {
	// ClassName - Additional classes added to the TextInput. Optional.
	ClassName string
	// IsDisabled - Flag to show if the input is disabled. Optional.
	IsDisabled bool
	// IsReadOnly - Flag to show if the input is read only. Optional.
	IsReadOnly bool
	// IsRequired - Flag to show if the input is required. Optional.
	IsRequired bool
	// Validated - Value to indicate if the input is modified to show that validation state. If set to success,
	// input will be modified to indicate valid state. If set to error,  input will be modified to indicate error
	// state. Optional.
	Validated any /* "success" | "warning" | "error" | "default" */
	// OnChange - A callback for when the input value changes. Optional.
	OnChange func(value string, event any // React.FormEvent )
	// Type - Type that the input accepts. Optional.
	Type any /* "text" | "date" | "datetime-local" | "email" | "month" | "number" | "password" | "search" | "tel" | "time" | "url" */
	// Value - Value of the input. Optional.
	Value any /* string | int */
	// AriaLabel - Aria-label. The input requires an associated id or aria-label. Optional.
	AriaLabel string
	// InnerRef - A reference object to attach to the input box. Optional.
	InnerRef any // React.RefObject 
	// IsLeftTruncated - Trim text on left. Optional.
	IsLeftTruncated bool
	// OnFocus - Callback function when input is focused. Optional.
	OnFocus func(event any)
	// OnBlur - Callback function when input is blurred (focus leaves). Optional.
	OnBlur func(event any)
	// IconVariant - icon variant. Optional.
	IconVariant any /* "calendar" | "clock" | "search" */
	// IsIconSprite - Use the external file instead of a data URI. Optional.
	IsIconSprite bool
	// CustomIconUrl - Custom icon url to set as the input's background-image. Optional.
	CustomIconUrl string
	// CustomIconDimensions - Dimensions for the custom icon set as the input's background-size. Optional.
	CustomIconDimensions string
}

type State struct {
	// OuiaStateId - 
	OuiaStateId string
}
