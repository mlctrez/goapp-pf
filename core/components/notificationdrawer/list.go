package notificationdrawer

type ListProps struct {
	// Children - Content rendered inside the notification drawer list body. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the notification drawer list body. Optional.
	ClassName string
	// IsHidden - Adds styling to the notification drawer list to indicate expand/hide state. Optional.
	IsHidden bool
}

// contents of react-core/src/components/NotificationDrawer/NotificationDrawerList.tsx
// import * as React from 'react';
// 
// import { css } from '@patternfly/react-styles';
// 
// export interface NotificationDrawerListProps extends React.HTMLProps<HTMLUListElement> {
//   /**  Content rendered inside the notification drawer list body */
//   children?: React.ReactNode;
//   /**  Additional classes added to the notification drawer list body */
//   className?: string;
//   /**  Adds styling to the notification drawer list to indicate expand/hide state */
//   isHidden?: boolean;
// }
// 
// export const NotificationDrawerList: React.FunctionComponent<NotificationDrawerListProps> = ({
//   children,
//   className = '',
//   isHidden = false,
//   ...props
// }: NotificationDrawerListProps) => (
//   <ul {...props} className={css('pf-c-notification-drawer__list', className)} hidden={isHidden}>
//     {children}
//   </ul>
// );
// NotificationDrawerList.displayName = 'NotificationDrawerList';
