package menu

type ListProps struct {
	// Children - Anything that can be rendered inside of menu list.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the menu list. Optional.
	ClassName string
}

// contents of react-core/src/components/Menu/MenuList.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Menu/menu';
// import { css } from '@patternfly/react-styles';
// 
// export interface MenuListProps extends React.HTMLProps<HTMLUListElement> {
//   /** Anything that can be rendered inside of menu list */
//   children: React.ReactNode;
//   /** Additional classes added to the menu list */
//   className?: string;
// }
// 
// export const MenuList: React.FunctionComponent<MenuListProps> = ({
//   children = null,
//   className,
//   ...props
// }: MenuListProps) => (
//   <ul role="menu" className={css(styles.menuList, className)} {...props}>
//     {children}
//   </ul>
// );
// MenuList.displayName = 'MenuList';
