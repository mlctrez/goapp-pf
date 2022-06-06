package popover

type DialogProps struct {
	// Position - PopoverDialog position. Optional.
	Position any //  /* "top" | "bottom" | "left" | "right" */
	// ClassName - PopoverDialog additional class. Optional.
	ClassName string
	// Children - PopoverDialog body.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Popover/PopoverDialog.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// import { css } from '@patternfly/react-styles';
// 
// export const PopoverPosition = {
//   top: 'top',
//   bottom: 'bottom',
//   left: 'left',
//   right: 'right'
// };
// 
// export const PopoverDialog: React.FunctionComponent<PopoverDialogProps> = ({
//   position = 'top',
//   children = null,
//   className = null,
//   ...props
// }: PopoverDialogProps) => (
//   <div
//     className={css(styles.popover, styles.modifiers[position] || styles.modifiers.top, className)}
//     role="dialog"
//     aria-modal="true"
//     {...props}
//   >
//     {children}
//   </div>
// );
// PopoverDialog.displayName = 'PopoverDialog';
// 
// export interface PopoverDialogProps extends React.HTMLProps<HTMLDivElement> {
//   /** PopoverDialog position */
//   position?: 'top' | 'bottom' | 'left' | 'right';
//   /** PopoverDialog additional class */
//   className?: string;
//   /** PopoverDialog body */
//   children: React.ReactNode;
// }
