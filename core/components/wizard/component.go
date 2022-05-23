package wizard

// from file "react-core/src/components/Wizard/Wizard.tsx"

type Step struct {
	// Id - Optional identifier. Optional.
	Id any /* string | int */
	// Name - The name of the step.
	Name any // React.ReactNode 
	// Component - The component to render in the main body. Optional.
	Component any
	// IsFinishedStep - Setting to true hides the side nav and footer. Optional.
	IsFinishedStep bool
	// CanJumpTo - Enables or disables the step in the navigation. Enabled by default. Optional.
	CanJumpTo bool
	// Steps - Sub steps. Optional.
	Steps []any // WizardStep 
	// StepNavItemProps - Props to pass to the WizardNavItem. Optional.
	StepNavItemProps any /* any // React.HTMLProps  | any // WizardNavItemProps  */
	// NextButtonText - (Unused if footer is controlled) Can change the Next button text. If nextButtonText is
	// also set for the Wizard, this step specific one overrides it. Optional.
	NextButtonText any // React.ReactNode 
	// EnableNext - (Unused if footer is controlled) The condition needed to enable the Next button. Optional.
	EnableNext bool
	// HideCancelButton - (Unused if footer is controlled) True to hide the Cancel button. Optional.
	HideCancelButton bool
	// HideBackButton - (Unused if footer is controlled) True to hide the Back button. Optional.
	HideBackButton bool
}

type Props struct {
	// Width - Custom width of the wizard. Optional.
	Width any /* int | string */
	// Height - Custom height of the wizard. Optional.
	Height any /* int | string */
	// Title - The wizard title to display if header is desired. Optional.
	Title string
	// TitleId - An optional id for the title. Optional.
	TitleId string
	// DescriptionId - An optional id for the description. Optional.
	DescriptionId string
	// Description - The wizard description. Optional.
	Description any // React.ReactNode 
	// DescriptionComponent - Component type of the description. Optional.
	DescriptionComponent any /* "div" | "p" */
	// HideClose - Flag indicating whether the close button should be in the header. Optional.
	HideClose bool
	// OnClose - Callback function to close the wizard. Optional.
	OnClose func()
	// OnGoToStep - Callback function when a step in the nav is clicked. Optional.
	OnGoToStep any // WizardStepFunctionType 
	// ClassName - Additional classes spread to the Wizard. Optional.
	ClassName string
	// Steps - The wizard steps configuration object.
	Steps []any // WizardStep 
	// StartAtStep - The current step the wizard is on (1 or higher). Optional.
	StartAtStep int
	// NavAriaLabel - Aria-label for the Nav. Optional.
	NavAriaLabel string
	// NavAriaLabelledBy - Sets aria-labelledby on nav element. Optional.
	NavAriaLabelledBy string
	// MainAriaLabel - Aria-label for the main element. Optional.
	MainAriaLabel string
	// MainAriaLabelledBy - Sets aria-labelledby on the main element. Optional.
	MainAriaLabelledBy string
	// HasNoBodyPadding - Can remove the default padding around the main body content by setting this to true.
	// Optional.
	HasNoBodyPadding bool
	// Footer - (Use to control the footer) Passing in a footer component lets you control the buttons yourself.
	// Optional.
	Footer any // React.ReactNode 
	// OnSave - (Unused if footer is controlled) Callback function to save at the end of the wizard, if not specified
	// uses onClose. Optional.
	OnSave func()
	// OnNext - (Unused if footer is controlled) Callback function after Next button is clicked. Optional.
	OnNext any // WizardStepFunctionType 
	// OnBack - (Unused if footer is controlled) Callback function after Back button is clicked. Optional.
	OnBack any // WizardStepFunctionType 
	// NextButtonText - (Unused if footer is controlled) The Next button text. Optional.
	NextButtonText any // React.ReactNode 
	// BackButtonText - (Unused if footer is controlled) The Back button text. Optional.
	BackButtonText any // React.ReactNode 
	// CancelButtonText - (Unused if footer is controlled) The Cancel button text. Optional.
	CancelButtonText any // React.ReactNode 
	// CloseButtonAriaLabel - (Unused if footer is controlled) aria-label for the close button. Optional.
	CloseButtonAriaLabel string
	// AppendTo - The parent container to append the modal to. Defaults to document.body. Optional.
	AppendTo any /* any // HTMLElement  | func() any // HTMLElement  */
	// IsOpen - Flag indicating Wizard modal is open. Wizard will be placed into a modal if this prop is provided.
	// Optional.
	IsOpen bool
	// IsNavExpandable - Flag indicating nav items with sub steps are expandable. Optional.
	IsNavExpandable bool
}

type State struct {
	// CurrentStep - 
	CurrentStep int
	// IsNavOpen - 
	IsNavOpen bool
}
