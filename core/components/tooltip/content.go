package tooltip

type ContentProps struct {
	// ClassName - PopoverContent additional class. Optional.
	ClassName string
	// Children - PopoverContent content.
	Children any //  // React.ReactNode 
	// IsLeftAligned - Flag to align text to the left. Optional.
	IsLeftAligned bool
}

// contents of react-core/src/components/Tooltip/TooltipContent.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Tooltip/tooltip';
// import { css } from '@patternfly/react-styles';
// 
// export interface TooltipContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** PopoverContent additional class */
//   className?: string;
//   /** PopoverContent content */
//   children: React.ReactNode;
//   /** Flag to align text to the left */
//   isLeftAligned?: boolean;
// }
// 
// export const TooltipContent: React.FunctionComponent<TooltipContentProps> = ({
//   className,
//   children,
//   isLeftAligned,
//   ...props
// }: TooltipContentProps) => (
//   <div className={css(styles.tooltipContent, isLeftAligned && styles.modifiers.textAlignLeft, className)} {...props}>
//     {children}
//   </div>
// );
// TooltipContent.displayName = 'TooltipContent';
