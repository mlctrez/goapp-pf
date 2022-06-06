package menu

type DrilldownMenuProps struct {
	// Children - Items within drilldown sub-menu. Optional.
	Children any //  // React.ReactNode 
	// Id - ID of the drilldown sub-menu. Optional.
	Id string
	// IsMenuDrilledIn - Flag indicating whether the menu is drilled in. Optional.
	IsMenuDrilledIn bool
	// GetHeight - Optional callback to get the height of the sub menu. Optional.
	GetHeight func(height string)
}

// contents of react-core/src/components/Menu/DrilldownMenu.tsx
// import React from 'react';
// import { Menu } from './Menu';
// import { MenuContent } from './MenuContent';
// import { MenuList } from './MenuList';
// import { MenuContext } from './MenuContext';
// 
// export interface DrilldownMenuProps extends Omit<React.HTMLAttributes<HTMLDivElement>, 'ref' | 'onSelect'> {
//   /** Items within drilldown sub-menu */
//   children?: React.ReactNode;
//   /** ID of the drilldown sub-menu */
//   id?: string;
//   /** Flag indicating whether the menu is drilled in */
//   isMenuDrilledIn?: boolean;
//   /** Optional callback to get the height of the sub menu */
//   getHeight?: (height: string) => void;
// }
// 
// export const DrilldownMenu: React.FunctionComponent<DrilldownMenuProps> = ({
//   children,
//   id,
//   isMenuDrilledIn = false,
//   getHeight,
//   ...props
// }: DrilldownMenuProps) => (
//   /* eslint-disable @typescript-eslint/no-unused-vars */
//   <MenuContext.Consumer>
//     {({ menuId, parentMenu, flyoutRef, setFlyoutRef, disableHover, ...context }) => (
//       <Menu
//         id={id}
//         parentMenu={menuId}
//         isMenuDrilledIn={isMenuDrilledIn}
//         isRootMenu={false}
//         ref={React.createRef()}
//         {...context}
//         {...props}
//       >
//         <MenuContent getHeight={getHeight}>
//           <MenuList>{children}</MenuList>
//         </MenuContent>
//       </Menu>
//     )}
//   </MenuContext.Consumer>
//   /* eslint-enable @typescript-eslint/no-unused-vars */
// );
// 
// DrilldownMenu.displayName = 'DrilldownMenu';
