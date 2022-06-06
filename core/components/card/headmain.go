package card

type HeadMainProps struct {
	// Children - Content rendered inside the Card Head Main. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Card Head Main. Optional.
	ClassName string
}

// contents of react-core/src/components/Card/CardHeadMain.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// 
// export interface CardHeadMainProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the Card Head Main */
//   children?: React.ReactNode;
//   /** Additional classes added to the Card Head Main */
//   className?: string;
// }
// 
// export const CardHeadMain: React.FunctionComponent<CardHeadMainProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: CardHeadMainProps) => (
//   <div className={css('pf-c-card__head-main', className)} {...props}>
//     {children}
//   </div>
// );
// CardHeadMain.displayName = 'CardHeadMain';
