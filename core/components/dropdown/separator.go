package dropdown

type SeparatorProps struct {
	// ClassName - Classes applied to root element of dropdown item. Optional.
	ClassName string
	// OnClick - Click event to pass to InternalDropdownItem. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
}

// contents of react-core/src/components/Dropdown/DropdownSeparator.tsx
// import * as React from 'react';
// import { DropdownArrowContext } from './dropdownConstants';
// import { InternalDropdownItem } from './InternalDropdownItem';
// import { Divider, DividerVariant } from '../Divider';
// import { useOUIAProps, OUIAProps } from '../../helpers';
// 
// export interface SeparatorProps extends React.HTMLProps<HTMLAnchorElement>, OUIAProps {
//   /** Classes applied to root element of dropdown item */
//   className?: string;
//   /** Click event to pass to InternalDropdownItem */
//   onClick?: (event: React.MouseEvent<HTMLAnchorElement> | React.KeyboardEvent | MouseEvent) => void;
// }
// 
// export const DropdownSeparator: React.FunctionComponent<SeparatorProps> = ({
//   className = '',
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   ref, // Types of Ref are different for React.FunctionComponent vs React.Component
//   ouiaId,
//   ouiaSafe,
//   ...props
// }: SeparatorProps) => {
//   const ouiaProps = useOUIAProps(DropdownSeparator.displayName, ouiaId, ouiaSafe);
//   return (
//     <DropdownArrowContext.Consumer>
//       {context => (
//         <InternalDropdownItem
//           {...props}
//           context={context}
//           component={<Divider component={DividerVariant.div} />}
//           className={className}
//           role="separator"
//           {...ouiaProps}
//         />
//       )}
//     </DropdownArrowContext.Consumer>
//   );
// };
// DropdownSeparator.displayName = 'DropdownSeparator';
