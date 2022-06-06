package overflowmenu

type DropdownItemProps struct {
	// IsShared - Indicates when a dropdown item shows and hides the corresponding list item. Optional.
	IsShared bool
	// Index - Indicates the index of the list item. Optional.
	Index int
}

// contents of react-core/src/components/OverflowMenu/OverflowMenuDropdownItem.tsx
// import * as React from 'react';
// import { DropdownItem, DropdownItemProps } from '../Dropdown';
// import { OverflowMenuContext } from './OverflowMenuContext';
// 
// export interface OverflowMenuDropdownItemProps extends DropdownItemProps {
//   /** Indicates when a dropdown item shows and hides the corresponding list item */
//   isShared?: boolean;
//   /** Indicates the index of the list item */
//   index?: number;
// }
// 
// export const OverflowMenuDropdownItem: React.FunctionComponent<OverflowMenuDropdownItemProps> = ({
//   children,
//   isShared = false,
//   index,
//   ...additionalProps
// }: OverflowMenuDropdownItemProps) => (
//   <OverflowMenuContext.Consumer>
//     {value =>
//       (!isShared || value.isBelowBreakpoint) && (
//         <DropdownItem component="button" index={index} {...additionalProps}>
//           {children}
//         </DropdownItem>
//       )
//     }
//   </OverflowMenuContext.Consumer>
// );
// OverflowMenuDropdownItem.displayName = 'OverflowMenuDropdownItem';
