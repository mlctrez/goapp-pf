package numberinput

// from file "react-core/src/components/NumberInput/NumberInput.tsx"

type Props struct {
	// Value - Value of the number input. Optional.
	Value int
	// ClassName - Additional classes added to the number input. Optional.
	ClassName string
	// WidthChars - Sets the width of the number input to a number of characters. Optional.
	WidthChars int
	// IsDisabled - Indicates the whole number input should be disabled. Optional.
	IsDisabled bool
	// OnMinus - Callback for the minus button. Optional.
	OnMinus func(event any // React.MouseEvent , name string)
	// OnChange - Callback for the text input changing. Optional.
	OnChange func(event any // React.FormEvent )
	// OnBlur - Callback function when text input is blurred (focus leaves). Optional.
	OnBlur func(event any)
	// OnPlus - Callback for the plus button. Optional.
	OnPlus func(event any // React.MouseEvent , name string)
	// Unit - Adds the given unit to the number input. Optional.
	Unit any // React.ReactNode 
	// UnitPosition - Position of the number input unit in relation to the number input. Optional.
	UnitPosition any /* "before" | "after" */
	// Min - Minimum value of the number input, disabling the minus button when reached. Optional.
	Min int
	// Max - Maximum value of the number input, disabling the plus button when reached. Optional.
	Max int
	// InputName - Name of the input. Optional.
	InputName string
	// InputAriaLabel - Aria label of the input. Optional.
	InputAriaLabel string
	// MinusBtnAriaLabel - Aria label of the minus button. Optional.
	MinusBtnAriaLabel string
	// PlusBtnAriaLabel - Aria label of the plus button. Optional.
	PlusBtnAriaLabel string
	// InputProps - Additional properties added to the text input. Optional.
	InputProps any
	// MinusBtnProps - Additional properties added to the minus button. Optional.
	MinusBtnProps any // ButtonProps 
	// PlusBtnProps - Additional properties added to the plus button. Optional.
	PlusBtnProps any // ButtonProps 
}
