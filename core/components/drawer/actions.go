package drawer

type ActionsProps struct {
	// ClassName - Additional classes added to the drawer actions button. Optional.
	ClassName string
	// Children - Actions to be rendered in the panel head. Optional.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/Drawer/DrawerActions.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Drawer/drawer';
// import { css } from '@patternfly/react-styles';
// 
// export interface DrawerActionsProps extends React.HTMLProps<HTMLDivElement> {
//   /** Additional classes added to the drawer actions button. */
//   className?: string;
//   /** Actions to be rendered in the panel head. */
//   children?: React.ReactNode;
// }
// 
// export const DrawerActions: React.FunctionComponent<DrawerActionsProps> = ({
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   className = '',
//   children,
//   ...props
// }: DrawerActionsProps) => (
//   <div className={css(styles.drawerActions, className)} {...props}>
//     {children}
//   </div>
// );
// DrawerActions.displayName = 'DrawerActions';
