package modal

// from file "react-core/src/components/Modal/ModalBox.tsx"

type BoxProps struct {
	// Children - Content rendered inside the ModalBox.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the ModalBox. Optional.
	ClassName string
	// Variant - Variant of the modal. Optional.
	Variant any /* "small" | "medium" | "large" | "default" */
	// Position - Alternate position of the modal. Optional.
	Position "top"
	// PositionOffset - Offset from alternate position. Can be any valid CSS length/percentage. Optional.
	PositionOffset string
	// AriaLabelledby - Id to use for Modal Box label. Optional.
	AriaLabelledby string
	// AriaLabel - Accessible descriptor of modal. Optional.
	AriaLabel string
	// AriaDescribedby - Id to use for Modal Box description.
	AriaDescribedby string
}
