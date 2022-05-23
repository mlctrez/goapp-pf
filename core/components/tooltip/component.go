package tooltip

// from file "react-core/src/components/Tooltip/Tooltip.tsx"

type Props struct {
	// AppendTo - The element to append the tooltip to, defaults to body. Optional.
	AppendTo any /* any // HTMLElement  | func(ref any // HTMLElement ) any // HTMLElement  */
	// Aria - aria-labelledby or aria-describedby for tooltip. The trigger will be cloned to add the aria attribute,
	// and the corresponding id in the form of 'pf-tooltip-#' is added to the content container. If you don't
	// want that or prefer to add the aria attribute yourself on the trigger, set aria to 'none'. Optional.
	Aria any /* "describedby" | "labelledby" | "none" */
	// AriaLive - Determines whether the tooltip is an aria-live region. If the reference prop is passed in the
	// default behavior is 'polite' in order to ensure the tooltip contents is announced to assistive technologies.
	// Otherwise the default behavior is 'off'. Optional.
	AriaLive any /* "off" | "polite" */
	// Children - The reference element to which the Tooltip is relatively placed to. If you cannot wrap the
	// reference with the Tooltip, you can use the reference prop instead. Usage: <Tooltip><Button>Reference</Button></Tooltip>.
	// Optional.
	Children any // ReactElement 
	// Reference - The reference element to which the Tooltip is relatively placed to. If you can wrap the reference
	// with the Tooltip, you can use the children prop instead. Usage: <Tooltip reference={() => document.getElementById('reference-element')}
	// />. Optional.
	Reference any /* any // HTMLElement  | func() any // HTMLElement  | any // React.RefObject  */
	// ClassName - Tooltip additional class. Optional.
	ClassName string
	// Content - Tooltip content.
	Content any // React.ReactNode 
	// Distance - Distance of the tooltip to its target, defaults to 15. Optional.
	Distance int
	// EnableFlip - If true, tries to keep the tooltip in view by flipping it if necessary. Optional.
	EnableFlip bool
	// EntryDelay - Delay in ms before the tooltip appears. Optional.
	EntryDelay int
	// ExitDelay - Delay in ms before the tooltip disappears. Optional.
	ExitDelay int
	// FlipBehavior - The desired position to flip the tooltip to if the initial position is not possible. By
	// setting this prop to 'flip' it attempts to flip the tooltip to the opposite side if there is no space.
	// You can also pass an array of positions that determines the flip order. It should contain the initial
	// position followed by alternative positions if that position is unavailable. Example: Initial position
	// is 'top'. Button with tooltip is in the top right corner. 'flipBehavior' is set to ['top', 'right', 'left'].
	// Since there is no space to the top, it checks if right is available. There's also no space to the right,
	// so it finally shows the tooltip on the left. Optional.
	FlipBehavior any /* "flip" | []( any /* "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */ ) */
	// MaxWidth - Maximum width of the tooltip (default 18.75rem). Optional.
	MaxWidth string
	// Position - Tooltip position. Note: With 'enableFlip' set to true, it will change the position if there
	// is not enough space for the starting position. The behavior of where it flips to can be controlled through
	// the flipBehavior prop. The 'auto' position chooses the side with the most space. The 'auto' position requires
	// the 'enableFlip' prop to be true. Optional.
	Position any /* any // TooltipPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// Trigger - Tooltip trigger: click, mouseenter, focus, manual Set to manual to trigger tooltip programmatically
	// (through the isVisible prop). Optional.
	Trigger string
	// IsContentLeftAligned - Flag to indicate that the text content is left aligned. Optional.
	IsContentLeftAligned bool
	// IsVisible - value for visibility when trigger is 'manual'. Optional.
	IsVisible bool
	// ZIndex - z-index of the tooltip. Optional.
	ZIndex int
	// Id - id of the tooltip. Optional.
	Id string
	// AnimationDuration - CSS fade transition animation duration. Optional.
	AnimationDuration int
	// Boundary - Optional.
	Boundary any /* "scrollParent" | "window" | "viewport" | any // HTMLElement  */
	// IsAppLauncher - Optional.
	IsAppLauncher bool
	// TippyProps - Optional.
	TippyProps any // Partial 
}
