package notificationdrawer

type Props struct {
	// Children - Content rendered inside the notification drawer. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the notification drawer. Optional.
	ClassName string
	// InnerRef - Optional.
	InnerRef any //  // React.Ref 
}

// contents of react-core/src/components/NotificationDrawer/NotificationDrawer.tsx
// import * as React from 'react';
// 
// import styles from '@patternfly/react-styles/css/components/NotificationDrawer/notification-drawer';
// import { css } from '@patternfly/react-styles';
// 
// export interface NotificationDrawerProps extends React.HTMLProps<HTMLDivElement> {
//   /**  Content rendered inside the notification drawer */
//   children?: React.ReactNode;
//   /**  Additional classes added to the notification drawer */
//   className?: string;
//   /** @hide Forwarded ref */
//   innerRef?: React.Ref<any>;
// }
// 
// const NotificationDrawerBase: React.FunctionComponent<NotificationDrawerProps> = ({
//   children,
//   className = '',
//   innerRef,
//   ...props
// }: NotificationDrawerProps) => (
//   <div ref={innerRef} {...props} className={css(styles.notificationDrawer, className)}>
//     {children}
//   </div>
// );
// export const NotificationDrawer = React.forwardRef((props: NotificationDrawerProps, ref: React.Ref<any>) => (
//   <NotificationDrawerBase innerRef={ref} {...props} />
// ));
// NotificationDrawer.displayName = 'NotificationDrawer';
