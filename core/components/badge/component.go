package badge

type Props struct {
	// IsRead - Adds styling to the badge to indicate it has been read. Optional.
	IsRead bool
	// Children - content rendered inside the Badge. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the Badge. Optional.
	ClassName string
}

// contents of react-core/src/components/Badge/Badge.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Badge/badge';
// 
// export interface BadgeProps extends React.HTMLProps<HTMLSpanElement> {
//   /**  Adds styling to the badge to indicate it has been read */
//   isRead?: boolean;
//   /** content rendered inside the Badge */
//   children?: React.ReactNode;
//   /** additional classes added to the Badge */
//   className?: string;
// }
// 
// export const Badge: React.FunctionComponent<BadgeProps> = ({
//   isRead = false,
//   className = '',
//   children = '',
//   ...props
// }: BadgeProps) => (
//   <span
//     {...props}
//     className={css(styles.badge, (isRead ? styles.modifiers.read : styles.modifiers.unread) as any, className)}
//   >
//     {children}
//   </span>
// );
// Badge.displayName = 'Badge';
