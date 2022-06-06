package overflowmenu

type GroupProps struct {
	// Children - Any elements that can be rendered in the menu. Optional.
	Children any // 
	// ClassName - Additional classes added to the OverflowMenuGroup. Optional.
	ClassName string
	// IsPersistent - Modifies the overflow menu group visibility. Optional.
	IsPersistent bool
	// GroupType - Indicates a button or icon group. Optional.
	GroupType any //  /* "button" | "icon" */
}

// contents of react-core/src/components/OverflowMenu/OverflowMenuGroup.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/OverflowMenu/overflow-menu';
// import { OverflowMenuContext } from './OverflowMenuContext';
// 
// export interface OverflowMenuGroupProps extends React.HTMLProps<HTMLDivElement> {
//   /** Any elements that can be rendered in the menu */
//   children?: any;
//   /** Additional classes added to the OverflowMenuGroup */
//   className?: string;
//   /** Modifies the overflow menu group visibility */
//   isPersistent?: boolean;
//   /** Indicates a button or icon group */
//   groupType?: 'button' | 'icon';
// }
// 
// export const OverflowMenuGroup: React.FunctionComponent<OverflowMenuGroupProps> = ({
//   className,
//   children,
//   isPersistent = false,
//   groupType,
//   ...props
// }: OverflowMenuGroupProps) => (
//   <OverflowMenuContext.Consumer>
//     {value =>
//       (isPersistent || !value.isBelowBreakpoint) && (
//         <div
//           className={css(
//             styles.overflowMenuGroup,
//             groupType === 'button' && styles.modifiers.buttonGroup,
//             groupType === 'icon' && styles.modifiers.iconButtonGroup,
//             className
//           )}
//           {...props}
//         >
//           {children}
//         </div>
//       )
//     }
//   </OverflowMenuContext.Consumer>
// );
// OverflowMenuGroup.displayName = 'OverflowMenuGroup';
