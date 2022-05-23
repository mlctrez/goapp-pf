package popover

// from file "react-core/src/components/Popover/Popover.tsx"

type Props struct {
	// AriaLabel - Accessible label, required when header is not present. Optional.
	AriaLabel string
	// AppendTo - The element to append the popover to, defaults to body. Optional.
	AppendTo any /* any // HTMLElement  | func(ref any // HTMLElement ) any // HTMLElement  */
	// BodyContent - Body content If you want to close the popover after an action within the bodyContent, you
	// can use the isVisible prop for manual control, or you can provide a function which will receive a callback
	// as an argument to hide the popover i.e. bodyContent={hide => <Button onClick={() => hide()}>Close</Button>}.
	BodyContent any /* any // React.ReactNode  | func(hide func()) any // React.ReactNode  */
	// Children - The reference element to which the Popover is relatively placed to. If you cannot wrap the
	// reference with the Popover, you can use the reference prop instead. Usage: <Popover><Button>Reference</Button></Popover>.
	// Optional.
	Children any // ReactElement 
	// Reference - The reference element to which the Popover is relatively placed to. If you can wrap the reference
	// with the Popover, you can use the children prop instead. Usage: <Popover reference={() => document.getElementById('reference-element')}
	// />. Optional.
	Reference any /* any // HTMLElement  | func() any // HTMLElement  | any // React.RefObject  */
	// ClassName - Popover additional class. Optional.
	ClassName string
	// CloseBtnAriaLabel - Aria label for the Close button. Optional.
	CloseBtnAriaLabel string
	// ShowClose - Whether to show the close button. Optional.
	ShowClose bool
	// Distance - Distance of the popover to its target, defaults to 25. Optional.
	Distance int
	// EnableFlip - If true, tries to keep the popover in view by flipping it if necessary. If the position is
	// set to 'auto', this prop is ignored. Optional.
	EnableFlip bool
	// FlipBehavior - The desired position to flip the popover to if the initial position is not possible. By
	// setting this prop to 'flip' it attempts to flip the popover to the opposite side if there is no space.
	// You can also pass an array of positions that determines the flip order. It should contain the initial
	// position followed by alternative positions if that position is unavailable. Example: Initial position
	// is 'top'. Button with popover is in the top right corner. 'flipBehavior' is set to ['top', 'right', 'left'].
	// Since there is no space to the top, it checks if right is available. There's also no space to the right,
	// so it finally shows the popover on the left. Optional.
	FlipBehavior any /* "flip" | []( any /* "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */ ) */
	// FooterContent - Footer content If you want to close the popover after an action within the bodyContent,
	// you can use the isVisible prop for manual control, or you can provide a function which will receive a
	// callback as an argument to hide the popover i.e. footerContent={hide => <Button onClick={() => hide()}>Close</Button>}.
	// Optional.
	FooterContent any /* any // React.ReactNode  | func(hide func()) any // React.ReactNode  */
	// HeaderContent - Simple header content to be placed within a title. If you want to close the popover after
	// an action within the bodyContent, you can use the isVisible prop for manual control, or you can provide
	// a function which will receive a callback as an argument to hide the popover i.e. headerContent={hide =>
	// <Button onClick={() => hide()}>Close</Button>}. Optional.
	HeaderContent any /* any // React.ReactNode  | func(hide func()) any // React.ReactNode  */
	// HeaderComponent - Sets the heading level to use for the popover header. Default is h6. Optional.
	HeaderComponent any /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
	// HeaderIcon - Optional.
	HeaderIcon any // React.ReactNode 
	// AlertSeverityVariant - Optional.
	AlertSeverityVariant any /* "default" | "info" | "warning" | "success" | "danger" */
	// AlertSeverityScreenReaderText - Optional.
	AlertSeverityScreenReaderText string
	// HideOnOutsideClick - Hides the popover when a click occurs outside (only works if isVisible is not controlled
	// by the user). Optional.
	HideOnOutsideClick bool
	// IsVisible - True to show the popover programmatically. Used in conjunction with the shouldClose prop.
	// By default, the popover child element handles click events automatically. If you want to control this
	// programmatically, the popover will not auto-close if the Close button is clicked, ESC key is used, or
	// if a click occurs outside the popover. Instead, the consumer is responsible for closing the popover themselves
	// by adding a callback listener for the shouldClose prop. Optional.
	IsVisible bool
	// MinWidth - Minimum width of the popover (default 6.25rem). Optional.
	MinWidth string
	// MaxWidth - Maximum width of the popover (default 18.75rem). Optional.
	MaxWidth string
	// OnHidden - Lifecycle function invoked when the popover has fully transitioned out. Note: The tip argument
	// is no longer passed and has been deprecated. Optional.
	OnHidden func(tip any // TippyInstance )
	// OnHide - Lifecycle function invoked when the popover begins to transition out. Note: The tip argument
	// is no longer passed and has been deprecated. Optional.
	OnHide func(tip any // TippyInstance )
	// OnMount - Lifecycle function invoked when the popover has been mounted to the DOM. Note: The tip argument
	// is no longer passed and has been deprecated. Optional.
	OnMount func(tip any // TippyInstance )
	// OnShow - Lifecycle function invoked when the popover begins to transition in. Note: The tip argument is
	// no longer passed and has been deprecated. Optional.
	OnShow func(tip any // TippyInstance )
	// OnShown - Lifecycle function invoked when the popover has fully transitioned in. Note: The tip argument
	// is no longer passed and has been deprecated. Optional.
	OnShown func(tip any // TippyInstance )
	// Position - Popover position. Note: With 'enableFlip' set to true, it will change the position if there
	// is not enough space for the starting position. The behavior of where it flips to can be controlled through
	// the flipBehavior prop. Optional.
	Position any /* any // PopoverPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// ShouldClose - Callback function that is only invoked when isVisible is also controlled. Called when the
	// popover Close button is clicked, Enter key was used on it, or the ESC key is used. Note: The tip argument
	// is no longer passed and has been deprecated. Optional.
	ShouldClose func(tip any // TippyInstance , hideFunction func(), event any /* any // MouseEvent  | any // KeyboardEvent  */)
	// ShouldOpen - Callback function that is only invoked when isVisible is also controlled. Called when the
	// Enter key is used on the focused trigger. Optional.
	ShouldOpen func(showFunction func(), event any /* any // MouseEvent  | any // KeyboardEvent  */)
	// ZIndex - z-index of the popover. Optional.
	ZIndex int
	// AnimationDuration - CSS fade transition animation duration. Optional.
	AnimationDuration int
	// Id - id used as part of the various popover elements (popover-${id}-header/body/footer). Optional.
	Id string
	// WithFocusTrap - Whether to trap focus in the popover. Optional.
	WithFocusTrap bool
	// HasAutoWidth - Removes fixed-width and allows width to be defined by contents. Optional.
	HasAutoWidth bool
	// HasNoPadding - Allows content to touch edges of popover container. Optional.
	HasNoPadding bool
	// Boundary - Optional.
	Boundary any /* "scrollParent" | "window" | "viewport" | any // HTMLElement  */
	// TippyProps - Optional.
	TippyProps any // Partial 
}
