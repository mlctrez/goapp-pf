package page

// from file "react-core/src/components/Page/PageToggleButton.tsx"

type ToggleButtonProps struct {
	// Children - Content of the page toggle button. Optional.
	Children any // React.ReactNode 
	// IsNavOpen - True if the side nav is shown. Optional.
	IsNavOpen bool
	// OnNavToggle - Callback function to handle the side nav toggle button, managed by the Page component if
	// the Page isManagedSidebar prop is set to true. Optional.
	OnNavToggle func()
}
