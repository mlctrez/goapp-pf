package masthead

type BrandProps struct {
	// Children - Content rendered inside of the masthead brand. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the masthead brand. Optional.
	ClassName string
	// Component - Component type of the masthead brand. Optional.
	Component any //  /* any // React.ElementType  | any // React.ComponentType  */
}

// contents of react-core/src/components/Masthead/MastheadBrand.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Masthead/masthead';
// import { css } from '@patternfly/react-styles';
// 
// export interface MastheadBrandProps
//   extends React.DetailedHTMLProps<React.HTMLProps<HTMLAnchorElement>, HTMLAnchorElement> {
//   /** Content rendered inside of the masthead brand. */
//   children?: React.ReactNode;
//   /** Additional classes added to the masthead brand. */
//   className?: string;
//   /** Component type of the masthead brand. */
//   component?: React.ElementType<any> | React.ComponentType<any>;
// }
// 
// export const MastheadBrand: React.FunctionComponent<MastheadBrandProps> = ({
//   children,
//   className,
//   component = 'a',
//   ...props
// }: MastheadBrandProps) => {
//   const Component = component as any;
//   return (
//     <Component className={css(styles.mastheadBrand, className)} tabIndex={0} {...props}>
//       {children}
//     </Component>
//   );
// };
// MastheadBrand.displayName = 'MastheadBrand';
