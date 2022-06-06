package gallery

type Props struct {
	// Children - content rendered inside the Gallery layout. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the Gallery layout. Optional.
	ClassName string
	// HasGutter - Adds space between children. Optional.
	HasGutter bool
	// MinWidths - Minimum widths at various breakpoints. Optional.
	MinWidths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// MaxWidths - Maximum widths at various breakpoints. Optional.
	MaxWidths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any //  /* any // React.ElementType  | any // React.ComponentType  */
}

// contents of react-core/src/layouts/Gallery/Gallery.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/layouts/Gallery/gallery';
// 
// export interface GalleryProps extends React.HTMLProps<HTMLDivElement> {
//   /** content rendered inside the Gallery layout */
//   children?: React.ReactNode;
//   /** additional classes added to the Gallery layout */
//   className?: string;
//   /** Adds space between children. */
//   hasGutter?: boolean;
//   /** Minimum widths at various breakpoints. */
//   minWidths?: {
//     default?: string;
//     sm?: string;
//     md?: string;
//     lg?: string;
//     xl?: string;
//     '2xl'?: string;
//   };
//   /** Maximum widths at various breakpoints. */
//   maxWidths?: {
//     default?: string;
//     sm?: string;
//     md?: string;
//     lg?: string;
//     xl?: string;
//     '2xl'?: string;
//   };
//   /** Sets the base component to render. defaults to div */
//   component?: React.ElementType<any> | React.ComponentType<any>;
// }
// export const Gallery: React.FunctionComponent<GalleryProps> = ({
//   children = null,
//   className = '',
//   component = 'div',
//   hasGutter = false,
//   minWidths,
//   maxWidths,
//   ...props
// }: GalleryProps) => {
//   const minWidthStyles: any = {};
//   const Component: any = component;
// 
//   if (minWidths) {
//     Object.entries(minWidths || {}).map(
//       ([breakpoint, value]) =>
//         (minWidthStyles[
//           `--pf-l-gallery--GridTemplateColumns--min${breakpoint !== 'default' ? `-on-${breakpoint}` : ''}`
//         ] = value)
//     );
//   }
//   const maxWidthStyles: any = {};
//   if (maxWidths) {
//     Object.entries(maxWidths || {}).map(
//       ([breakpoint, value]) =>
//         (maxWidthStyles[
//           `--pf-l-gallery--GridTemplateColumns--max${breakpoint !== 'default' ? `-on-${breakpoint}` : ''}`
//         ] = value)
//     );
//   }
//   const widthStyles = { ...minWidthStyles, ...maxWidthStyles };
// 
//   return (
//     <Component
//       className={css(styles.gallery, hasGutter && styles.modifiers.gutter, className)}
//       {...props}
//       {...((minWidths || maxWidths) && { style: { ...widthStyles, ...props.style } as React.CSSProperties })}
//     >
//       {children}
//     </Component>
//   );
// };
// Gallery.displayName = 'Gallery';
