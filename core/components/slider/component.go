package slider

// from file "react-core/src/components/Slider/Slider.tsx"

type StepObject struct {
	// Value - Value of the step. This value is a percentage of the slider where the  tick is drawn.
	Value int
	// Label - The display label for the step value. This is also used for the aria-valuetext.
	Label string
	// IsLabelHidden - Flag to hide the label. Optional.
	IsLabelHidden bool
}

type Props struct {
	// ClassName - Additional classes added to the spinner. Optional.
	ClassName string
	// Value - Current value. Optional.
	Value int
	// AreCustomStepsContinuous - Flag indicating if the slider is is discrete for custom steps.  This will cause
	// the slider to snap to the closest value. Optional.
	AreCustomStepsContinuous bool
	// IsDisabled - Adds disabled styling and disables the slider and the input component is present. Optional.
	IsDisabled bool
	// Step - The step interval. Optional.
	Step int
	// Min - Minimum permitted value. Optional.
	Min int
	// Max - The maximum permitted value. Optional.
	Max int
	// ShowBoundaries - Flag to indicate if boundaries should be shown for slider that does not have custom steps.
	// Optional.
	ShowBoundaries bool
	// ShowTicks - Flag to indicate if ticks should be shown for slider that does not have custom steps. Optional.
	ShowTicks bool
	// CustomSteps - Array of custom slider step objects (value and label of each step) for the slider. Optional.
	CustomSteps []any // SliderStepObject 
	// IsInputVisible - Flag to show value input field. Optional.
	IsInputVisible bool
	// InputValue - Value displayed in the input field. Optional.
	InputValue int
	// InputAriaLabel - Aria label for the input field. Optional.
	InputAriaLabel string
	// ThumbAriaLabel - Optional.
	ThumbAriaLabel string
	// HasTooltipOverThumb - Optional.
	HasTooltipOverThumb bool
	// InputLabel - Label that is place after the input field. Optional.
	InputLabel any /* string | int */
	// InputPosition - Position of the input. Optional.
	InputPosition any /* "aboveThumb" | "right" */
	// OnChange - Value change callback. This is called when the slider value changes. Optional.
	OnChange func(value int, inputValue int, setLocalInputValue any // React.Dispatch )
	// LeftActions - Actions placed to the left of the slider. Optional.
	LeftActions any // React.ReactNode 
	// RightActions - Actions placed to the right of the slider. Optional.
	RightActions any // React.ReactNode 
}
