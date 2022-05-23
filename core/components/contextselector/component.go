package contextselector

// from file "react-core/src/components/ContextSelector/ContextSelector.tsx"

type Props struct {
	// Children - content rendered inside the Context Selector. Optional.
	Children any // React.ReactNode 
	// ClassName - Classes applied to root element of Context Selector. Optional.
	ClassName string
	// IsOpen - Flag to indicate if Context Selector is opened. Optional.
	IsOpen bool
	// OnToggle - Function callback called when user clicks toggle button. Optional.
	OnToggle func(event any, value bool)
	// OnSelect - Function callback called when user selects item. Optional.
	OnSelect func(event any, value any // React.ReactNode )
	// IsFullHeight - Flag indicating that the context selector should expand to full height. Optional.
	IsFullHeight bool
	// ScreenReaderLabel - Labels the Context Selector for Screen Readers. Optional.
	ScreenReaderLabel string
	// ToggleText - Text that appears in the Context Selector Toggle. Optional.
	ToggleText string
	// SearchButtonAriaLabel - Aria-label for the Context Selector Search Button. Optional.
	SearchButtonAriaLabel string
	// SearchInputValue - Value in the Search field. Optional.
	SearchInputValue string
	// OnSearchInputChange - Function callback called when user changes the Search Input. Optional.
	OnSearchInputChange func(value string)
	// SearchInputPlaceholder - Search Input placeholder. Optional.
	SearchInputPlaceholder string
	// OnSearchButtonClick - Function callback for when Search Button is clicked. Optional.
	OnSearchButtonClick func(event any // React.SyntheticEvent )
	// Footer - Footer of the context selector. Optional.
	Footer any // React.ReactNode 
	// IsPlain - Flag to indicate the toggle has no border or background. Optional.
	IsPlain bool
	// IsText - Flag to indicate if toggle is textual toggle. Optional.
	IsText bool
	// DisableFocusTrap - Flag to disable focus trap. Optional.
	DisableFocusTrap bool
}
