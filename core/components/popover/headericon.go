package popover

type HeaderIconProps struct {
	// Children - Content of the header icon.
	Children any //  // React.ReactNode 
	// ClassName - Class to be applied to the header icon. Optional.
	ClassName string
}

// contents of react-core/src/components/Popover/PopoverHeaderIcon.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// 
// export interface PopoverHeaderIconProps extends React.HTMLProps<HTMLSpanElement> {
//   /** Content of the header icon */
//   children: React.ReactNode;
//   /** Class to be applied to the header icon */
//   className?: string;
// }
// 
// export const PopoverHeaderIcon: React.FunctionComponent<PopoverHeaderIconProps> = ({
//   children,
//   className,
//   ...props
// }: PopoverHeaderIconProps) => (
//   <span className={css(styles.popoverTitleIcon, className)} {...props}>
//     {children}
//   </span>
// );
// PopoverHeaderIcon.displayName = 'PopoverHeaderIcon';
