package textarea

// from file "react-core/src/components/TextArea/TextArea.tsx"

type Props struct {
	// ClassName - Additional classes added to the TextArea. Optional.
	ClassName string
	// IsRequired - Flag to show if the TextArea is required. Optional.
	IsRequired bool
	// IsDisabled - Flag to show if the TextArea is disabled. Optional.
	IsDisabled bool
	// IsReadOnly - Flag to show if the TextArea is read only. Optional.
	IsReadOnly bool
	// IsIconSprite - Use the external file instead of a data URI. Optional.
	IsIconSprite bool
	// AutoResize - Flag to modify height based on contents. Optional.
	AutoResize bool
	// Validated - Value to indicate if the textarea is modified to show that validation state. If set to success,
	// textarea will be modified to indicate valid state. If set to error, textarea will be modified to indicate
	// error state. Optional.
	Validated any /* "success" | "warning" | "error" | "default" */
	// Value - Value of the TextArea. Optional.
	Value any /* string | int */
	// OnChange - A callback for when the TextArea value changes. Optional.
	OnChange func(value string, event any // React.ChangeEvent )
	// ResizeOrientation - Sets the orientation to limit the resize to. Optional.
	ResizeOrientation any /* "horizontal" | "vertical" | "both" */
	// AriaLabel - Custom flag to show that the TextArea requires an associated id or aria-label. Optional.
	AriaLabel string
	// InnerRef - A reference object to attach to the textarea. Optional.
	InnerRef any // React.RefObject 
}
