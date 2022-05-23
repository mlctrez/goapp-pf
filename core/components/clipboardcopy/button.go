package clipboardcopy

// from file "react-core/src/components/ClipboardCopy/ClipboardCopyButton.tsx"

type ButtonProps struct {
	// OnClick - Callback for the copy when the button is clicked.
	OnClick func(event any // React.MouseEvent )
	// Children - Content of the copy button.
	Children any // React.ReactNode 
	// Id - ID of the copy button.
	Id string
	// TextId - ID of the content that is being copied.
	TextId string
	// ClassName - Additional classes added to the copy button. Optional.
	ClassName string
	// ExitDelay - Exit delay on the copy button tooltip. Optional.
	ExitDelay int
	// EntryDelay - Entry delay on the copy button tooltip. Optional.
	EntryDelay int
	// MaxWidth - Max width of the copy button tooltip. Optional.
	MaxWidth string
	// Position - Position of the copy button tooltip. Optional.
	Position any /* any // TooltipPosition  | any // PopoverPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// AriaLabel - Aria-label for the copy button. Optional.
	AriaLabel string
	// Variant - Variant of the copy button. Optional.
	Variant any /* "control" | "plain" */
}
