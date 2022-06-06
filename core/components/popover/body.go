package popover

type BodyProps struct {
	// Id - Popover body id.
	Id string
	// Children - Popover body content.
	Children any //  // React.ReactNode 
	// ClassName - Classes to be applied to the popover body. Optional.
	ClassName string
}

// contents of react-core/src/components/Popover/PopoverBody.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Popover/popover';
// import { css } from '@patternfly/react-styles';
// 
// export interface PopoverBodyProps extends React.HTMLProps<HTMLDivElement> {
//   /** Popover body id */
//   id: string;
//   /** Popover body content */
//   children: React.ReactNode;
//   /** Classes to be applied to the popover body. */
//   className?: string;
// }
// 
// export const PopoverBody: React.FunctionComponent<PopoverBodyProps> = ({
//   children,
//   id,
//   className,
//   ...props
// }: PopoverBodyProps) => (
//   <div className={css(styles.popoverBody, className)} id={id} {...props}>
//     {children}
//   </div>
// );
// PopoverBody.displayName = 'PopoverBody';
