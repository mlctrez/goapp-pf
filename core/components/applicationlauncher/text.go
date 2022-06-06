package applicationlauncher

type TextProps struct {
	// Children - content rendered inside the text container.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the text container. Optional.
	ClassName string
}

// contents of react-core/src/components/ApplicationLauncher/ApplicationLauncherText.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// 
// export interface ApplicationLauncherTextProps extends React.HTMLProps<HTMLSpanElement> {
//   /** content rendered inside the text container */
//   children: React.ReactNode;
//   /** Additional classes added to the text container */
//   className?: string;
// }
// 
// export const ApplicationLauncherText: React.FunctionComponent<ApplicationLauncherTextProps> = ({
//   className = '',
//   children,
//   ...props
// }: ApplicationLauncherTextProps) => (
//   <span className={css('pf-c-app-launcher__menu-item-text', className)} {...props}>
//     {children}
//   </span>
// );
// ApplicationLauncherText.displayName = 'ApplicationLauncherText';
