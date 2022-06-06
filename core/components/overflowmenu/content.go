package overflowmenu

type ContentProps struct {
	// Children - Any elements that can be rendered in the menu. Optional.
	Children any // 
	// ClassName - Additional classes added to the OverflowMenuContent. Optional.
	ClassName string
	// IsPersistent - Modifies the overflow menu content visibility. Optional.
	IsPersistent bool
}

// contents of react-core/src/components/OverflowMenu/OverflowMenuContent.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/OverflowMenu/overflow-menu';
// import { OverflowMenuContext } from './OverflowMenuContext';
// 
// export interface OverflowMenuContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** Any elements that can be rendered in the menu */
//   children?: any;
//   /** Additional classes added to the OverflowMenuContent */
//   className?: string;
//   /** Modifies the overflow menu content visibility */
//   isPersistent?: boolean;
// }
// 
// export const OverflowMenuContent: React.FunctionComponent<OverflowMenuContentProps> = ({
//   className,
//   children,
//   isPersistent
// }: OverflowMenuContentProps) => (
//   <OverflowMenuContext.Consumer>
//     {value =>
//       (!value.isBelowBreakpoint || isPersistent) && (
//         <div className={css(styles.overflowMenuContent, className)}>{children}</div>
//       )
//     }
//   </OverflowMenuContext.Consumer>
// );
// OverflowMenuContent.displayName = 'OverflowMenuContent';
