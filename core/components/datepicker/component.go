package datepicker

// from file "react-core/src/components/DatePicker/DatePicker.tsx"

type Props struct {
	// ClassName - Additional classes added to the date time picker. Optional.
	ClassName string
	// AriaLabel - Accessible label for the date picker. Optional.
	AriaLabel string
	// DateFormat - How to format the date in the TextInput. Optional.
	DateFormat func(date any // Date ) string
	// DateParse - How to format the date in the TextInput. Optional.
	DateParse func(value string) any // Date 
	// IsDisabled - Flag indicating the date picker is disabled. Optional.
	IsDisabled bool
	// Placeholder - String to display in the empty date picker field as a hint for the expected date format.
	// Optional.
	Placeholder string
	// Value - Value of TextInput. Optional.
	Value string
	// InvalidFormatText - Error message to display when the TextInput cannot be parsed. Optional.
	InvalidFormatText string
	// OnChange - Callback called every time the input value changes. Optional.
	OnChange func(value string, date any // Date )
	// OnBlur - Callback called every time the input loses focus. Optional.
	OnBlur func(value string, date any // Date )
	// HelperText - Text for label. Optional.
	HelperText any // React.ReactNode 
	// ButtonAriaLabel - Aria label for the button to open the date picker. Optional.
	ButtonAriaLabel string
	// AppendTo - The element to append the popover to. Optional.
	AppendTo any /* any // HTMLElement  | func(ref any // HTMLElement ) any // HTMLElement  */
	// PopoverProps - Props to pass to the Popover. Optional.
	PopoverProps any // Omit 
	// Validators - Functions that returns an error message if a date is invalid. Optional.
	Validators []( func(date any // Date ) string )
	// InputProps - Additional props for input field. Optional.
	InputProps any // TextInputProps 
}

type Ref struct {
	// SetCalendarOpen - Sets the calendar open status.
	SetCalendarOpen func(isOpen bool)
	// ToggleCalendar - Toggles the calendar open status.
	ToggleCalendar func()
}
