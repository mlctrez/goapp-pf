package clipboardcopy

// from file "react-core/src/components/ClipboardCopy/ClipboardCopy.tsx"

type State struct {
	// Text - 
	Text any /* string | int */
	// Expanded - 
	Expanded bool
	// Copied - 
	Copied bool
}

type Props struct {
	// ClassName - Additional classes added to the clipboard copy container. Optional.
	ClassName string
	// HoverTip - Tooltip message to display when hover the copy button. Optional.
	HoverTip string
	// ClickTip - Tooltip message to display when clicking the copy button. Optional.
	ClickTip string
	// TextAriaLabel - Aria-label to use on the TextInput. Optional.
	TextAriaLabel string
	// ToggleAriaLabel - Aria-label to use on the ClipboardCopyToggle. Optional.
	ToggleAriaLabel string
	// IsReadOnly - Flag to show if the input is read only. Optional.
	IsReadOnly bool
	// IsExpanded - Flag to determine if clipboard copy is in the expanded state initially. Optional.
	IsExpanded bool
	// IsCode - Flag to determine if clipboard copy content includes code. Optional.
	IsCode bool
	// IsBlock - Flag to determine if inline clipboard copy should be block styling. Optional.
	IsBlock bool
	// Variant - Adds Clipboard Copy variant styles. Optional.
	Variant any /* typeof ClipboardCopyVariant | "inline" | "expansion" | "inline-compact" */
	// Position - Copy button popover position. Optional.
	Position any /* any // PopoverPosition  | any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// MaxWidth - Maximum width of the tooltip (default 150px). Optional.
	MaxWidth string
	// ExitDelay - Delay in ms before the tooltip disappears. Optional.
	ExitDelay int
	// EntryDelay - Delay in ms before the tooltip appears. Optional.
	EntryDelay int
	// SwitchDelay - Delay in ms before the tooltip message switch to hover tip. Optional.
	SwitchDelay int
	// OnCopy - A function that is triggered on clicking the copy button. Optional.
	OnCopy func(event any // React.ClipboardEvent , text any // React.ReactNode )
	// OnChange - A function that is triggered on changing the text. Optional.
	OnChange func(text any /* string | int */)
	// Children - The text which is copied.
	Children any // React.ReactNode 
	// AdditionalActions - Additional actions for inline clipboard copy. Should be wrapped with ClipboardCopyAction.
	// Optional.
	AdditionalActions any // React.ReactNode 
}
