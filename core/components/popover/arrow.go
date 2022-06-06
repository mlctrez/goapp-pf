package popover

type ArrowProps struct {
	// ClassName - Popover arrow additional className. Optional.
	ClassName string
}

// contents of react-core/src/components/Popover/PopoverArrow.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// import { css } from '@patternfly/react-styles';
// 
// export const PopoverArrow: React.FunctionComponent<PopoverArrowProps> = ({
//   className = '',
//   ...props
// }: PopoverArrowProps) => <div className={css(styles.popoverArrow, className)} {...props} />;
// PopoverArrow.displayName = 'PopoverArrow';
// 
// export interface PopoverArrowProps extends React.HTMLProps<HTMLDivElement> {
//   /** Popover arrow additional className */
//   className?: string;
// }
