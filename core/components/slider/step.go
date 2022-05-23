package slider

// from file "react-core/src/components/Slider/SliderStep.tsx"

type StepProps struct {
	// ClassName - Additional classes added to the slider steps. Optional.
	ClassName string
	// Value - Step value *. Optional.
	Value int
	// Label - Step label *. Optional.
	Label string
	// IsTickHidden - Flag indicating that the tick should be hidden. Optional.
	IsTickHidden bool
	// IsLabelHidden - Flag indicating that the label should be hidden. Optional.
	IsLabelHidden bool
	// IsActive - Flag indicating the step is active. Optional.
	IsActive bool
}
