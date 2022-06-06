package clipboardcopy

type ButtonProps struct {
	// OnClick - Callback for the copy when the button is clicked.
	OnClick func(event any // React.MouseEvent )
	// Children - Content of the copy button.
	Children any //  // React.ReactNode 
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
	Position any //  /* any // TooltipPosition  | any // PopoverPosition  | "auto" | "top" | "bottom" | "left" | "right" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "left-start" | "left-end" | "right-start" | "right-end" */
	// AriaLabel - Aria-label for the copy button. Optional.
	AriaLabel string
	// Variant - Variant of the copy button. Optional.
	Variant any //  /* "control" | "plain" */
}

// contents of react-core/src/components/ClipboardCopy/ClipboardCopyButton.tsx
// import * as React from 'react';
// import CopyIcon from '@patternfly/react-icons/dist/esm/icons/copy-icon';
// import { Button } from '../Button';
// import { Tooltip, TooltipPosition } from '../Tooltip';
// import { PopoverPosition } from '../Popover';
// 
// export interface ClipboardCopyButtonProps
//   extends Omit<React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>, 'ref'> {
//   /** Callback for the copy when the button is clicked */
//   onClick: (event: React.MouseEvent) => void;
//   /** Content of the copy button */
//   children: React.ReactNode;
//   /** ID of the copy button */
//   id: string;
//   /** ID of the content that is being copied */
//   textId: string;
//   /** Additional classes added to the copy button */
//   className?: string;
//   /** Exit delay on the copy button tooltip */
//   exitDelay?: number;
//   /** Entry delay on the copy button tooltip */
//   entryDelay?: number;
//   /** Max width of the copy button tooltip */
//   maxWidth?: string;
//   /** Position of the copy button tooltip */
//   position?:
//     | TooltipPosition
//     | PopoverPosition
//     | 'auto'
//     | 'top'
//     | 'bottom'
//     | 'left'
//     | 'right'
//     | 'top-start'
//     | 'top-end'
//     | 'bottom-start'
//     | 'bottom-end'
//     | 'left-start'
//     | 'left-end'
//     | 'right-start'
//     | 'right-end';
//   /** Aria-label for the copy button */
//   'aria-label'?: string;
//   /** Variant of the copy button */
//   variant?: 'control' | 'plain';
// }
// 
// export const ClipboardCopyButton: React.FunctionComponent<ClipboardCopyButtonProps> = ({
//   onClick,
//   exitDelay = 0,
//   entryDelay = 300,
//   maxWidth = '100px',
//   position = 'top',
//   'aria-label': ariaLabel = 'Copyable input',
//   id,
//   textId,
//   children,
//   variant = 'control',
//   ...props
// }: ClipboardCopyButtonProps) => (
//   <Tooltip
//     trigger="mouseenter focus click"
//     exitDelay={exitDelay}
//     entryDelay={entryDelay}
//     maxWidth={maxWidth}
//     position={position}
//     aria-live="polite"
//     aria="none"
//     content={<div>{children}</div>}
//   >
//     <Button
//       type="button"
//       variant={variant}
//       onClick={onClick}
//       aria-label={ariaLabel}
//       id={id}
//       aria-labelledby={`${id} ${textId}`}
//       {...props}
//     >
//       <CopyIcon />
//     </Button>
//   </Tooltip>
// );
// ClipboardCopyButton.displayName = 'ClipboardCopyButton';
