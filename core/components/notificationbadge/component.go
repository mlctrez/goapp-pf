package notificationbadge

type Props struct {
	// IsRead - Optional.
	IsRead bool
	// Variant - Determines the variant of the notification badge. Optional.
	Variant any //  /* any // NotificationBadgeVariant  | "read" | "unread" | "attention" */
	// Count - A number displayed in the badge alongside the icon. Optional.
	Count int
	// Children - content rendered inside the notification badge. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the notification badge. Optional.
	ClassName string
	// AriaLabel - Adds accessible text to the notification badge. Optional.
	AriaLabel string
	// AttentionIcon - Icon to display for attention variant. Optional.
	AttentionIcon any //  // React.ReactNode 
	// Icon - Icon do display in notification badge. Optional.
	Icon any //  // React.ReactNode 
}

// contents of react-core/src/components/NotificationBadge/NotificationBadge.tsx
// import * as React from 'react';
// import { Button, ButtonVariant, ButtonProps } from '../Button';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/NotificationBadge/notification-badge';
// import AttentionBellIcon from '@patternfly/react-icons/dist/esm/icons/attention-bell-icon';
// import BellIcon from '@patternfly/react-icons/dist/esm/icons/bell-icon';
// 
// export enum NotificationBadgeVariant {
//   read = 'read',
//   unread = 'unread',
//   attention = 'attention'
// }
// 
// export interface NotificationBadgeProps extends Omit<ButtonProps, 'variant'> {
//   /** @deprecated Use the variant prop instead - Adds styling to the notification badge to indicate it has been read */
//   isRead?: boolean;
//   /** Determines the variant of the notification badge */
//   variant?: NotificationBadgeVariant | 'read' | 'unread' | 'attention';
//   /** A number displayed in the badge alongside the icon */
//   count?: number;
//   /** content rendered inside the notification badge */
//   children?: React.ReactNode;
//   /** additional classes added to the notification badge */
//   className?: string;
//   /** Adds accessible text to the notification badge. */
//   'aria-label'?: string;
//   /** Icon to display for attention variant */
//   attentionIcon?: React.ReactNode;
//   /** Icon do display in notification badge */
//   icon?: React.ReactNode;
// }
// 
// export const NotificationBadge: React.FunctionComponent<NotificationBadgeProps> = ({
//   isRead,
//   children,
//   variant = isRead ? 'read' : 'unread',
//   count = 0,
//   attentionIcon = <AttentionBellIcon />,
//   icon = <BellIcon />,
//   className,
//   ...props
// }: NotificationBadgeProps) => {
//   let notificationChild = icon;
//   if (children !== undefined) {
//     notificationChild = children;
//   } else if (variant === NotificationBadgeVariant.attention) {
//     notificationChild = attentionIcon;
//   }
//   return (
//     <Button variant={ButtonVariant.plain} className={className} {...props}>
//       <span className={css(styles.notificationBadge, styles.modifiers[variant])}>
//         {notificationChild}
//         {count > 0 && <span className={css(styles.notificationBadgeCount)}>{count}</span>}
//       </span>
//     </Button>
//   );
// };
// NotificationBadge.displayName = 'NotificationBadge';
