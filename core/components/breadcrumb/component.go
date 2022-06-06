package breadcrumb

type Props struct {
	// Children - Children nodes be rendered to the BreadCrumb. Should be of type BreadCrumbItem. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the breadcrumb nav. Optional.
	ClassName string
	// AriaLabel - Aria label added to the breadcrumb nav. Optional.
	AriaLabel string
}

// contents of react-core/src/components/Breadcrumb/Breadcrumb.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Breadcrumb/breadcrumb';
// import { css } from '@patternfly/react-styles';
// import { useOUIAProps, OUIAProps } from '../../helpers';
// 
// export interface BreadcrumbProps extends React.HTMLProps<HTMLElement>, OUIAProps {
//   /** Children nodes be rendered to the BreadCrumb. Should be of type BreadCrumbItem. */
//   children?: React.ReactNode;
//   /** Additional classes added to the breadcrumb nav. */
//   className?: string;
//   /** Aria label added to the breadcrumb nav. */
//   'aria-label'?: string;
// }
// 
// export const Breadcrumb: React.FunctionComponent<BreadcrumbProps> = ({
//   children = null,
//   className = '',
//   'aria-label': ariaLabel = 'Breadcrumb',
//   ouiaId,
//   ouiaSafe = true,
//   ...props
// }: BreadcrumbProps) => {
//   const ouiaProps = useOUIAProps(Breadcrumb.displayName, ouiaId, ouiaSafe);
//   return (
//     <nav {...props} aria-label={ariaLabel} className={css(styles.breadcrumb, className)} {...ouiaProps}>
//       <ol className={styles.breadcrumbList}>
//         {React.Children.map(children, (child, index) => {
//           const showDivider = index > 0;
//           if (React.isValidElement(child)) {
//             return React.cloneElement(child, { showDivider });
//           }
// 
//           return child;
//         })}
//       </ol>
//     </nav>
//   );
// };
// Breadcrumb.displayName = 'Breadcrumb';
