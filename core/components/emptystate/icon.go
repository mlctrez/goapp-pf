package emptystate

type IconProps struct {
	// Color - Changes the color of the icon. Optional.
	Color string
}

type IconProps struct {
	// ClassName - Additional classes added to the EmptyState. Optional.
	ClassName string
	// Icon - Icon component to be rendered inside the EmptyState on icon variant Usually a CheckCircleIcon,
	// ExclamationCircleIcon, LockIcon, PlusCircleIcon, RocketIcon SearchIcon, or WrenchIcon. Optional.
	Icon any //  // React.ComponentType 
	// Component - Component to be rendered inside the EmptyState on container variant. Optional.
	Component any //  // React.ComponentType 
	// Variant - Adds empty state icon variant styles. Optional.
	Variant any //  /* "icon" | "container" */
}

// contents of react-core/src/components/EmptyState/EmptyStateIcon.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/EmptyState/empty-state';
// 
// export interface IconProps extends Omit<React.HTMLProps<SVGElement>, 'size'> {
//   /** Changes the color of the icon.  */
//   color?: string;
// }
// 
// export interface EmptyStateIconProps extends IconProps {
//   /** Additional classes added to the EmptyState */
//   className?: string;
//   /** Icon component to be rendered inside the EmptyState on icon variant
//    * Usually a CheckCircleIcon, ExclamationCircleIcon, LockIcon, PlusCircleIcon, RocketIcon
//    * SearchIcon, or WrenchIcon */
//   icon?: React.ComponentType<any>;
//   /** Component to be rendered inside the EmptyState on container variant */
//   component?: React.ComponentType<any>;
//   /** Adds empty state icon variant styles  */
//   variant?: 'icon' | 'container';
// }
// 
// export const EmptyStateIcon: React.FunctionComponent<EmptyStateIconProps> = ({
//   className = '',
//   icon: IconComponent,
//   component: AnyComponent,
//   variant = 'icon',
//   ...props
// }: EmptyStateIconProps) => {
//   const classNames = css(styles.emptyStateIcon, className);
//   return variant === 'icon' ? (
//     <IconComponent className={classNames} {...props} aria-hidden="true" />
//   ) : (
//     <div className={classNames}>
//       <AnyComponent />
//     </div>
//   );
// };
// EmptyStateIcon.displayName = 'EmptyStateIcon';
