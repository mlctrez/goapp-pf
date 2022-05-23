package timepicker

// from file "react-core/src/components/TimePicker/TimePicker.tsx"

type Props struct {
	// ClassName - Additional classes added to the time picker. Optional.
	ClassName string
	// AriaLabel - Accessible label for the time picker. Optional.
	AriaLabel string
	// IsDisabled - Flag indicating the time picker is disabled. Optional.
	IsDisabled bool
	// Placeholder - String to display in the empty time picker field as a hint for the expected time format.
	// Optional.
	Placeholder string
	// Delimiter - Character to display between the hour and minute. Optional.
	Delimiter string
	// Time - A time string. The format could be  an ISO 8601 formatted date string or in 'HH{delimiter}MM' format.
	// Optional.
	Time any /* string | any // Date  */
	// InvalidFormatErrorMessage - Error message to display when the time is provided in an invalid format. Optional.
	InvalidFormatErrorMessage string
	// InvalidMinMaxErrorMessage - Error message to display when the time provided is not within the minTime/maxTime
	// constriants. Optional.
	InvalidMinMaxErrorMessage string
	// Is24Hour - True if the time is 24 hour time. False if the time is 12 hour time. Optional.
	Is24Hour bool
	// OnChange - Optional event handler called each time the value in the time picker input changes. Optional.
	OnChange func(time string, hour int, minute int, seconds int, isValid bool)
	// ValidateTime - Optional validator can be provided to override the internal time validator. Optional.
	ValidateTime func(time string) bool
	// Id - Id of the time picker. Optional.
	Id string
	// Width - Width of the time picker. Optional.
	Width string
	// MenuAppendTo - The container to append the menu to. Defaults to 'inline'. If your menu is being cut off
	// you can append it to an element higher up the DOM tree. Some examples: menuAppendTo="parent" menuAppendTo={()
	// => document.body} menuAppendTo={document.getElementById('target')}. Optional.
	MenuAppendTo any /* any // HTMLElement  | func() any // HTMLElement  | "inline" | "parent" */
	// StepMinutes - Size of step between time options in minutes. Optional.
	StepMinutes int
	// InputProps - Additional props for input field. Optional.
	InputProps any // TextInputProps 
	// MinTime - A time string indicating the minimum value allowed. The format could be an ISO 8601 formatted
	// date string or in 'HH{delimiter}MM' format. Optional.
	MinTime any /* string | any // Date  */
	// MaxTime - A time string indicating the maximum value allowed. The format could be an ISO 8601 formatted
	// date string or in 'HH{delimiter}MM' format. Optional.
	MaxTime any /* string | any // Date  */
	// IncludeSeconds - Includes number of seconds with the chosen time and allows users to manually edit the
	// seconds value. Optional.
	IncludeSeconds bool
}

type State struct {
	// IsInvalid - 
	IsInvalid bool
	// IsOpen - 
	IsOpen bool
	// TimeState - 
	TimeState string
	// FocusedIndex - 
	FocusedIndex int
	// ScrollIndex - 
	ScrollIndex int
	// TimeRegex - 
	TimeRegex any // RegExp 
	// MinTimeState - 
	MinTimeState string
	// MaxTimeState - 
	MaxTimeState string
}
