package banner

type Props struct {
	// Children - Content rendered inside the banner. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the banner. Optional.
	ClassName string
	// Variant - Variant styles for the banner. Optional.
	Variant any //  /* "default" | "info" | "danger" | "success" | "warning" */
	// IsSticky - If set to true, the banner sticks to the top of its container. Optional.
	IsSticky bool
}

// contents of react-core/src/components/Banner/Banner.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Banner/banner';
// import { css } from '@patternfly/react-styles';
// 
// export interface BannerProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the banner */
//   children?: React.ReactNode;
//   /** Additional classes added to the banner */
//   className?: string;
//   /** Variant styles for the banner */
//   variant?: 'default' | 'info' | 'danger' | 'success' | 'warning';
//   /** If set to true, the banner sticks to the top of its container */
//   isSticky?: boolean;
// }
// 
// export const Banner: React.FunctionComponent<BannerProps> = ({
//   children,
//   className,
//   variant = 'default',
//   isSticky = false,
//   ...props
// }: BannerProps) => (
//   <div
//     className={css(
//       styles.banner,
//       styles.modifiers[variant as 'success' | 'danger' | 'warning' | 'info'],
//       isSticky && styles.modifiers.sticky,
//       className
//     )}
//     {...props}
//   >
//     {children}
//   </div>
// );
// Banner.displayName = 'Banner';
