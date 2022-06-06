package descriptionlist

type BreakpointModifiers struct {
	// Default - Optional.
	Default string
	// Md - Optional.
	Md string
	// Lg - Optional.
	Lg string
	// Xl - Optional.
	Xl string
	// 2xl - Optional.
	2xl string
}

type Props struct {
	// Children - Anything that can be rendered inside of the list. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the list. Optional.
	ClassName string
	// IsAutoFit - Sets the description list to auto fit. Optional.
	IsAutoFit bool
	// IsHorizontal - Sets the description list component term and description pair to a horizontal layout. Optional.
	IsHorizontal bool
	// IsAutoColumnWidths - Sets the description list to format automatically. Optional.
	IsAutoColumnWidths bool
	// IsInlineGrid - Modifies the description list display to inline-grid. Optional.
	IsInlineGrid bool
	// IsCompact - Sets the description list to compact styling. Optional.
	IsCompact bool
	// IsFluid - Sets a horizontal description list to have fluid styling. Optional.
	IsFluid bool
	// IsFillColumns - Sets the the default placement of description list groups to fill from top to bottom.
	// Optional.
	IsFillColumns bool
	// ColumnModifier - Sets the number of columns on the description list at various breakpoints. Optional.
	ColumnModifier map[string]string /* { default:{ 1Col, 2Col, 3Col }, sm:{ 1Col, 2Col, 3Col }, md:{ 1Col, 2Col, 3Col }, lg:{ 1Col, 2Col, 3Col }, xl:{ 1Col, 2Col, 3Col }, 2xl:{ 1Col, 2Col, 3Col } } */
	// Orientation - Indicates how the menu will align at various breakpoints. Optional.
	Orientation map[string]string /* { sm:{ vertical, horizontal }, md:{ vertical, horizontal }, lg:{ vertical, horizontal }, xl:{ vertical, horizontal }, 2xl:{ vertical, horizontal } } */
	// AutoFitMinModifier - Sets the minimum column size for the auto-fit (isAutoFit) layout at various breakpoints.
	// Optional.
	AutoFitMinModifier map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// HorizontalTermWidthModifier - Sets the horizontal description list's term column width at various breakpoints.
	// Optional.
	HorizontalTermWidthModifier map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
}

// contents of react-core/src/components/DescriptionList/DescriptionList.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/DescriptionList/description-list';
// import { formatBreakpointMods } from '../../helpers';
// 
// export interface BreakpointModifiers {
//   default?: string;
//   md?: string;
//   lg?: string;
//   xl?: string;
//   '2xl'?: string;
// }
// 
// export interface DescriptionListProps extends Omit<React.HTMLProps<HTMLDListElement>, 'type'> {
//   /** Anything that can be rendered inside of the list */
//   children?: React.ReactNode;
//   /** Additional classes added to the list */
//   className?: string;
//   /** Sets the description list to auto fit. */
//   isAutoFit?: boolean;
//   /** Sets the description list component term and description pair to a horizontal layout. */
//   isHorizontal?: boolean;
//   /** Sets the description list to format automatically. */
//   isAutoColumnWidths?: boolean;
//   /** Modifies the description list display to inline-grid. */
//   isInlineGrid?: boolean;
//   /** Sets the description list to compact styling. */
//   isCompact?: boolean;
//   /** Sets a horizontal description list to have fluid styling. */
//   isFluid?: boolean;
//   /** Sets the the default placement of description list groups to fill from top to bottom. */
//   isFillColumns?: boolean;
//   /** Sets the number of columns on the description list at various breakpoints */
//   columnModifier?: {
//     default?: '1Col' | '2Col' | '3Col';
//     sm?: '1Col' | '2Col' | '3Col';
//     md?: '1Col' | '2Col' | '3Col';
//     lg?: '1Col' | '2Col' | '3Col';
//     xl?: '1Col' | '2Col' | '3Col';
//     '2xl'?: '1Col' | '2Col' | '3Col';
//   };
//   /** Indicates how the menu will align at various breakpoints. */
//   orientation?: {
//     sm?: 'vertical' | 'horizontal';
//     md?: 'vertical' | 'horizontal';
//     lg?: 'vertical' | 'horizontal';
//     xl?: 'vertical' | 'horizontal';
//     '2xl'?: 'vertical' | 'horizontal';
//   };
//   /** Sets the minimum column size for the auto-fit (isAutoFit) layout at various breakpoints. */
//   autoFitMinModifier?: {
//     default?: string;
//     sm?: string;
//     md?: string;
//     lg?: string;
//     xl?: string;
//     '2xl'?: string;
//   };
//   /** Sets the horizontal description list's term column width at various breakpoints. */
//   horizontalTermWidthModifier?: {
//     default?: string;
//     sm?: string;
//     md?: string;
//     lg?: string;
//     xl?: string;
//     '2xl'?: string;
//   };
// }
// 
// const setBreakpointModifiers = (prefix: string, modifiers: BreakpointModifiers) => {
//   const mods = modifiers as Partial<{ [k: string]: string }>;
//   return Object.keys(mods || {}).reduce(
//     (acc, curr) =>
//       curr === 'default' ? { ...acc, [prefix]: mods[curr] } : { ...acc, [`${prefix}-on-${curr}`]: mods[curr] },
//     {}
//   );
// };
// 
// export const DescriptionList: React.FunctionComponent<DescriptionListProps> = ({
//   className = '',
//   children = null,
//   isHorizontal = false,
//   isAutoColumnWidths,
//   isAutoFit,
//   isInlineGrid,
//   isCompact,
//   isFluid,
//   isFillColumns,
//   columnModifier,
//   autoFitMinModifier,
//   horizontalTermWidthModifier,
//   orientation,
//   style,
//   ...props
// }: DescriptionListProps) => {
//   if (isAutoFit && autoFitMinModifier) {
//     style = {
//       ...style,
//       ...setBreakpointModifiers('--pf-c-description-list--GridTemplateColumns--min', autoFitMinModifier)
//     };
//   }
//   if (isHorizontal && horizontalTermWidthModifier) {
//     style = {
//       ...style,
//       ...setBreakpointModifiers('--pf-c-description-list--m-horizontal__term--width', horizontalTermWidthModifier)
//     };
//   }
// 
//   return (
//     <dl
//       className={css(
//         styles.descriptionList,
//         (isHorizontal || isFluid) && styles.modifiers.horizontal,
//         isAutoColumnWidths && styles.modifiers.autoColumnWidths,
//         isAutoFit && styles.modifiers.autoFit,
//         formatBreakpointMods(columnModifier, styles),
//         formatBreakpointMods(orientation, styles),
//         isInlineGrid && styles.modifiers.inlineGrid,
//         isCompact && styles.modifiers.compact,
//         isFluid && styles.modifiers.fluid,
//         isFillColumns && styles.modifiers.fillColumns,
//         className
//       )}
//       style={style}
//       {...props}
//     >
//       {children}
//     </dl>
//   );
// };
// 
// DescriptionList.displayName = 'DescriptionList';
