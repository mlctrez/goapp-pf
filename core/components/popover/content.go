package popover

type ContentProps struct {
	// ClassName - PopoverContent additional class. Optional.
	ClassName string
	// Children - PopoverContent content.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Popover/PopoverContent.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// import { css } from '@patternfly/react-styles';
// 
// export const PopoverContent: React.FunctionComponent<PopoverContentProps> = ({
//   className = null,
//   children,
//   ...props
// }: PopoverContentProps) => (
//   <div className={css(styles.popoverContent, className)} {...props}>
//     {children}
//   </div>
// );
// PopoverContent.displayName = 'PopoverContent';
// 
// export interface PopoverContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** PopoverContent additional class */
//   className?: string;
//   /** PopoverContent content */
//   children: React.ReactNode;
// }
