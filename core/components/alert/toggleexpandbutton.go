package alert

// from file "react-core/src/components/Alert/AlertToggleExpandButton.tsx"

type ToggleExpandButtonProps struct {
	// AriaLabel - Aria label for the toggle button. Optional.
	AriaLabel string
	// OnToggleExpand - A callback for when the toggle button is clicked. Optional.
	OnToggleExpand func()
	// IsExpanded - Flag to indicate if the content is expanded. Optional.
	IsExpanded bool
	// VariantLabel - Variant label for the toggle button. Optional.
	VariantLabel string
}
