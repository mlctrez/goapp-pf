package popover

type HeaderTextProps struct {
	// Children - Content of the header text.
	Children any //  // React.ReactNode 
	// ClassName - Class to be applied to the header text. Optional.
	ClassName string
}

// contents of react-core/src/components/Popover/PopoverHeaderText.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// 
// export interface PopoverHeaderTextProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content of the header text */
//   children: React.ReactNode;
//   /** Class to be applied to the header text */
//   className?: string;
// }
// 
// export const PopoverHeaderText: React.FunctionComponent<PopoverHeaderTextProps> = ({
//   children,
//   className,
//   ...props
// }: PopoverHeaderTextProps) => (
//   <span className={css(styles.popoverTitleText, className)} {...props}>
//     {children}
//   </span>
// );
// PopoverHeaderText.displayName = 'PopoverHeaderText';
