package alert

type ActionLinkProps struct {
	// Children - Content rendered inside the AlertLinkAction. Optional.
	Children string
	// ClassName - Additional classes added to the AlertActionLink. Optional.
	ClassName string
}

// contents of react-core/src/components/Alert/AlertActionLink.tsx
// import * as React from 'react';
// import { Button, ButtonVariant, ButtonProps } from '../Button';
// 
// export interface AlertActionLinkProps extends ButtonProps {
//   /** Content rendered inside the AlertLinkAction  */
//   children?: string;
//   /** Additional classes added to the AlertActionLink  */
//   className?: string;
// }
// 
// export const AlertActionLink: React.FunctionComponent<AlertActionLinkProps> = ({
//   className = '',
//   children,
//   ...props
// }: AlertActionLinkProps) => (
//   <Button variant={ButtonVariant.link} isInline className={className} {...props}>
//     {children}
//   </Button>
// );
// AlertActionLink.displayName = 'AlertActionLink';
