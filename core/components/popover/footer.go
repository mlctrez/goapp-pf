package popover

type FooterProps struct {
	// ClassName - Additional classes added to the Popover footer. Optional.
	ClassName string
	// Children - Footer node.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Popover/PopoverFooter.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// import { css } from '@patternfly/react-styles';
// 
// export interface PopoverFooterProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the Popover footer */
//   className?: string;
//   /** Footer node */
//   children: React.ReactNode;
// }
// 
// export const PopoverFooter: React.FunctionComponent<PopoverFooterProps> = ({
//   children,
//   className = '',
//   ...props
// }: PopoverFooterProps) => (
//   <footer className={css(styles.popoverFooter, className)} {...props}>
//     {children}
//   </footer>
// );
// PopoverFooter.displayName = 'PopoverFooter';
