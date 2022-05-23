package progressstepper

// from file "react-core/src/components/ProgressStepper/ProgressStep.tsx"

type ProgressStepProps struct {
	// Children - Content rendered inside the progress step. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes applied to the progress step container. Optional.
	ClassName string
	// Variant - Variant of the progress step. Each variant has a default icon. Optional.
	Variant any /* "default" | "success" | "info" | "pending" | "warning" | "danger" */
	// IsCurrent - Flag indicating the progress step is the current step. Optional.
	IsCurrent bool
	// Icon - Custom icon of a progress step. Will override default icons provided by the variant. Optional.
	Icon any // React.ReactNode 
	// Description - Description text of a progress step. Optional.
	Description string
	// TitleId - ID of the title of the progress step. Optional.
	TitleId string
	// AriaLabel - Accessible label for the progress step. Should communicate all information being communicated
	// by the progress step's icon, including the variant and the completed status. Optional.
	AriaLabel string
	// PopoverRender - Forwards the step ref to rendered function.  Use this prop to add a popover to the step.
	// Optional.
	PopoverRender func(stepRef any // React.RefObject ) any // React.ReactNode 
}
