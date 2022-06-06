package accordion

type ContentProps struct {
	// Children - Content rendered inside the Accordion. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the Accordion content. Optional.
	ClassName string
	// Id - Identify the AccordionContent item. Optional.
	Id string
	// IsHidden - Flag to show if the expanded content of the Accordion item is visible. Optional.
	IsHidden bool
	// IsFixed - Flag to indicate Accordion content is fixed. Optional.
	IsFixed bool
	// AriaLabel - Adds accessible text to the Accordion content. Optional.
	AriaLabel string
	// Component - Component to use as content container. Optional.
	Component any //  // React.ElementType 
	// IsCustomContent - Flag indicating content is custom. Expanded content Body wrapper will be removed from
	// children.  This allows multiple bodies to be rendered as content. Optional.
	IsCustomContent any //  // React.ReactNode 
}

// contents of react-core/src/components/Accordion/AccordionContent.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Accordion/accordion';
// import { AccordionContext } from './AccordionContext';
// import { AccordionExpandedContentBody } from './AccordionExpandedContentBody';
// 
// export interface AccordionContentProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the Accordion  */
//   children?: React.ReactNode;
//   /** Additional classes added to the Accordion content  */
//   className?: string;
//   /** Identify the AccordionContent item  */
//   id?: string;
//   /** Flag to show if the expanded content of the Accordion item is visible  */
//   isHidden?: boolean;
//   /** Flag to indicate Accordion content is fixed  */
//   isFixed?: boolean;
//   /** Adds accessible text to the Accordion content */
//   'aria-label'?: string;
//   /** Component to use as content container */
//   component?: React.ElementType;
//   /** Flag indicating content is custom. Expanded content Body wrapper will be removed from children.  This allows multiple bodies to be rendered as content. */
//   isCustomContent?: React.ReactNode;
// }
// 
// export const AccordionContent: React.FunctionComponent<AccordionContentProps> = ({
//   className = '',
//   children = null,
//   id = '',
//   isHidden = false,
//   isFixed = false,
//   isCustomContent = false,
//   'aria-label': ariaLabel = '',
//   component,
//   ...props
// }: AccordionContentProps) => (
//   <AccordionContext.Consumer>
//     {({ ContentContainer }) => {
//       const Container = component || ContentContainer;
//       return (
//         <Container
//           id={id}
//           className={css(
//             styles.accordionExpandedContent,
//             isFixed && styles.modifiers.fixed,
//             !isHidden && styles.modifiers.expanded,
//             className
//           )}
//           hidden={isHidden}
//           aria-label={ariaLabel}
//           {...props}
//         >
//           {isCustomContent ? children : <AccordionExpandedContentBody>{children}</AccordionExpandedContentBody>}
//         </Container>
//       );
//     }}
//   </AccordionContext.Consumer>
// );
// AccordionContent.displayName = 'AccordionContent';
