package menutoggle

// from file "react-core/src/components/MenuToggle/MenuToggle.tsx"

type Props struct {
	// Children - Content rendered inside the toggle. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the toggle. Optional.
	ClassName string
	// IsExpanded - Flag indicating the toggle has expanded styling. Optional.
	IsExpanded bool
	// IsDisabled - Flag indicating the toggle is disabled. Optional.
	IsDisabled bool
	// IsFullHeight - Flag indicating the toggle is full height. Optional.
	IsFullHeight bool
	// IsFullWidth - Flag indicating the toggle takes up the full width of its parent. Optional.
	IsFullWidth bool
	// Variant - Variant styles of the menu toggle. Optional.
	Variant any /* "default" | "plain" | "primary" | "plainText" | "secondary" */
	// Icon - Optional icon rendered inside the toggle, before the children content. Optional.
	Icon any // React.ReactNode 
	// Badge - Optional badge rendered inside the toggle, after the children content. Optional.
	Badge any /* any // BadgeProps  | any // React.ReactNode  */
	// InnerRef - Forwarded ref. Optional.
	InnerRef any // React.Ref 
}
