package alert

// from file "react-core/src/components/Alert/AlertActionCloseButton.tsx"

type ActionCloseButtonProps struct {
	// ClassName - Additional classes added to the AlertActionCloseButton. Optional.
	ClassName string
	// OnClose - A callback for when the close button is clicked. Optional.
	OnClose func()
	// AriaLabel - Aria Label for the Close button. Optional.
	AriaLabel string
	// VariantLabel - Variant Label for the Close button. Optional.
	VariantLabel string
}
