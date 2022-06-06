package avatar

type Props struct {
	// ClassName - Additional classes added to the Avatar. Optional.
	ClassName string
	// Src - Attribute that specifies the URL of the image for the Avatar. Optional.
	Src string
	// Alt - Attribute that specifies the alternate text of the image for the Avatar.
	Alt string
	// Border - Border to add. Optional.
	Border any //  /* "light" | "dark" */
	// Size - Size variant of avatar. Optional.
	Size any //  /* "sm" | "md" | "lg" | "xl" */
}

// contents of react-core/src/components/Avatar/Avatar.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Avatar/avatar';
// import { css } from '@patternfly/react-styles';
// 
// export interface AvatarProps
//   extends React.DetailedHTMLProps<React.ImgHTMLAttributes<HTMLImageElement>, HTMLImageElement> {
//   /** Additional classes added to the Avatar. */
//   className?: string;
//   /** Attribute that specifies the URL of the image for the Avatar. */
//   src?: string;
//   /** Attribute that specifies the alternate text of the image for the Avatar. */
//   alt: string;
//   /** Border to add */
//   border?: 'light' | 'dark';
//   /** Size variant of avatar. */
//   size?: 'sm' | 'md' | 'lg' | 'xl';
// }
// 
// export const Avatar: React.FunctionComponent<AvatarProps> = ({
//   className = '',
//   src = '',
//   alt,
//   border,
//   size,
//   ...props
// }: AvatarProps) => (
//   <img
//     src={src}
//     alt={alt}
//     className={css(
//       styles.avatar,
//       styles.modifiers[size],
//       border === 'light' && styles.modifiers.light,
//       border === 'dark' && styles.modifiers.dark,
//       className
//     )}
//     {...props}
//   />
// );
// Avatar.displayName = 'Avatar';
