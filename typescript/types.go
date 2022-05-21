package typescript

import (
	"github.com/mlctrez/goapp-pf/typescript/kind"
)

type AmdDependency struct {
	Node
}

type Argument struct {
	Node
	ArgumentExpression       *ArgumentExpression     `json:"argumentExpression,omitempty"`
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	ColonToken               *ColonToken             `json:"colonToken,omitempty"`
	Condition                *Condition              `json:"condition,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Head                     *Head                   `json:"head,omitempty"`
	Left                     *Left                   `json:"left,omitempty"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	Operand                  *Operand                `json:"operand,omitempty"`
	Operator                 int                     `json:"operator"`
	OperatorToken            *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                     `json:"originalKeywordKind"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	QuestionToken            *QuestionToken          `json:"questionToken,omitempty"`
	Right                    *Right                  `json:"right,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan         `json:"templateSpans,omitempty"`
	Text                     string                  `json:"text"`
	Type                     *Type                   `json:"type,omitempty"`
	WhenFalse                *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue                 *WhenTrue               `json:"whenTrue,omitempty"`
}

type ArgumentExpression struct {
	Node
	ArgumentExpression       *ArgumentExpression `json:"argumentExpression,omitempty"`
	Arguments                []*Argument         `json:"arguments,omitempty"`
	EscapedText              string              `json:"escapedText"`
	Expression               *Expression         `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                `json:"hasExtendedUnicodeEscape"`
	Head                     *Head               `json:"head,omitempty"`
	Left                     *Left               `json:"left,omitempty"`
	Name                     *Name               `json:"name,omitempty"`
	NumericLiteralFlags      int                 `json:"numericLiteralFlags"`
	OperatorToken            *OperatorToken      `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                 `json:"originalKeywordKind"`
	Right                    *Right              `json:"right,omitempty"`
	TemplateSpans            []*TemplateSpan     `json:"templateSpans,omitempty"`
	Text                     string              `json:"text"`
	Type                     *Type               `json:"type,omitempty"`
}

type Attributes struct {
	Node
	Properties []*Property `json:"properties,omitempty"`
}

type BindDiagnostic struct {
	Node
}

type Block struct {
	Node
	MultiLine  bool         `json:"multiLine"`
	Statements []*Statement `json:"statements,omitempty"`
}

type Body struct {
	Node
	Arguments              []*Argument             `json:"arguments,omitempty"`
	Attributes             *Attributes             `json:"attributes,omitempty"`
	Body                   *Body                   `json:"body,omitempty"`
	Children               []*Children             `json:"children,omitempty"`
	ClosingElement         *ClosingElement         `json:"closingElement,omitempty"`
	ColonToken             *ColonToken             `json:"colonToken,omitempty"`
	Condition              *Condition              `json:"condition,omitempty"`
	Elements               []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText            string                  `json:"escapedText"`
	Expression             *Expression             `json:"expression,omitempty"`
	Head                   *Head                   `json:"head,omitempty"`
	Left                   *Left                   `json:"left,omitempty"`
	MultiLine              bool                    `json:"multiLine"`
	Name                   *Name                   `json:"name,omitempty"`
	OpeningElement         *OpeningElement         `json:"openingElement,omitempty"`
	Operand                *Operand                `json:"operand,omitempty"`
	Operator               int                     `json:"operator"`
	OperatorToken          *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind    int                     `json:"originalKeywordKind"`
	Parameters             []*Parameter            `json:"parameters,omitempty"`
	QuestionToken          *QuestionToken          `json:"questionToken,omitempty"`
	Right                  *Right                  `json:"right,omitempty"`
	Statements             []*Statement            `json:"statements,omitempty"`
	TagName                *TagName                `json:"tagName,omitempty"`
	TemplateSpans          []*TemplateSpan         `json:"templateSpans,omitempty"`
	Type                   *Type                   `json:"type,omitempty"`
	WhenFalse              *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue               *WhenTrue               `json:"whenTrue,omitempty"`
}

type CaseBlock struct {
	Node
	Clauses []*Clause `json:"clauses,omitempty"`
}

type CatchClause struct {
	Node
	Block               *Block               `json:"block,omitempty"`
	VariableDeclaration *VariableDeclaration `json:"variableDeclaration,omitempty"`
}

type Children struct {
	Node
	Attributes                    *Attributes     `json:"attributes,omitempty"`
	Children                      []*Children     `json:"children,omitempty"`
	ClosingElement                *ClosingElement `json:"closingElement,omitempty"`
	ContainsOnlyTriviaWhiteSpaces bool            `json:"containsOnlyTriviaWhiteSpaces"`
	Expression                    *Expression     `json:"expression,omitempty"`
	OpeningElement                *OpeningElement `json:"openingElement,omitempty"`
	TagName                       *TagName        `json:"tagName,omitempty"`
	Text                          string          `json:"text"`
}

type Clause struct {
	Node
	Expression *Expression  `json:"expression,omitempty"`
	Statements []*Statement `json:"statements,omitempty"`
}

type ClosingElement struct {
	Node
	TagName *TagName `json:"tagName,omitempty"`
}

type ClosingFragment struct {
	Node
}

type ColonToken struct {
	Node
}

type Condition struct {
	Node
	ArgumentExpression *ArgumentExpression `json:"argumentExpression,omitempty"`
	Arguments          []*Argument         `json:"arguments,omitempty"`
	EscapedText        string              `json:"escapedText"`
	Expression         *Expression         `json:"expression,omitempty"`
	Left               *Left               `json:"left,omitempty"`
	Name               *Name               `json:"name,omitempty"`
	Operand            *Operand            `json:"operand,omitempty"`
	Operator           int                 `json:"operator"`
	OperatorToken      *OperatorToken      `json:"operatorToken,omitempty"`
	QuestionDotToken   *QuestionDotToken   `json:"questionDotToken,omitempty"`
	Right              *Right              `json:"right,omitempty"`
}

type Constraint struct {
	Node
	IndexType  *IndexType  `json:"indexType,omitempty"`
	ObjectType *ObjectType `json:"objectType,omitempty"`
}

type Declaration struct {
	Node
	Initializer *Initializer `json:"initializer,omitempty"`
	Name        *Name        `json:"name,omitempty"`
	Type        *Type        `json:"type,omitempty"`
}

type DeclarationList struct {
	Node
	Declarations []*Declaration `json:"declarations,omitempty"`
}

type DotDotDotToken struct {
	Node
}

type Element struct {
	Node
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	ColonToken               *ColonToken             `json:"colonToken,omitempty"`
	Condition                *Condition              `json:"condition,omitempty"`
	DotDotDotToken           *DotDotDotToken         `json:"dotDotDotToken,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Initializer              *Initializer            `json:"initializer,omitempty"`
	IsTypeOnly               bool                    `json:"isTypeOnly"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	PropertyName             *PropertyName           `json:"propertyName,omitempty"`
	QuestionToken            *QuestionToken          `json:"questionToken,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	Text                     string                  `json:"text"`
	Type                     *Type                   `json:"type,omitempty"`
	WhenFalse                *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue                 *WhenTrue               `json:"whenTrue,omitempty"`
}

type ElementType struct {
	Node
	ElementType   *ElementType    `json:"elementType,omitempty"`
	Type          *Type           `json:"type,omitempty"`
	TypeArguments []*TypeArgument `json:"typeArguments,omitempty"`
	TypeName      *TypeName       `json:"typeName,omitempty"`
}

type ElseStatement struct {
	Node
	ElseStatement *ElseStatement `json:"elseStatement,omitempty"`
	Expression    *Expression    `json:"expression,omitempty"`
	MultiLine     bool           `json:"multiLine"`
	Statements    []*Statement   `json:"statements,omitempty"`
	ThenStatement *ThenStatement `json:"thenStatement,omitempty"`
}

type EndOfFileToken struct {
	Node
}

type EqualsGreaterThanToken struct {
	Node
}

type ExclamationToken struct {
	Node
}

type ExprName struct {
	Node
	EscapedText string `json:"escapedText"`
}

type Expression struct {
	Node
	ArgumentExpression       *ArgumentExpression     `json:"argumentExpression,omitempty"`
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	ClosingFragment          *ClosingFragment        `json:"closingFragment,omitempty"`
	ColonToken               *ColonToken             `json:"colonToken,omitempty"`
	Condition                *Condition              `json:"condition,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Head                     *Head                   `json:"head,omitempty"`
	Left                     *Left                   `json:"left,omitempty"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	OpeningFragment          *OpeningFragment        `json:"openingFragment,omitempty"`
	Operand                  *Operand                `json:"operand,omitempty"`
	Operator                 int                     `json:"operator"`
	OperatorToken            *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                     `json:"originalKeywordKind"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	QuestionDotToken         *QuestionDotToken       `json:"questionDotToken,omitempty"`
	QuestionToken            *QuestionToken          `json:"questionToken,omitempty"`
	Right                    *Right                  `json:"right,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan         `json:"templateSpans,omitempty"`
	Text                     string                  `json:"text"`
	Type                     *Type                   `json:"type,omitempty"`
	TypeArguments            []*TypeArgument         `json:"typeArguments,omitempty"`
	WhenFalse                *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue                 *WhenTrue               `json:"whenTrue,omitempty"`
}

type ExternalModuleIndicator struct {
	Node
	DeclarationList *DeclarationList `json:"declarationList,omitempty"`
	ImportClause    *ImportClause    `json:"importClause,omitempty"`
	Modifiers       []*Modifier      `json:"modifiers,omitempty"`
	ModuleSpecifier *ModuleSpecifier `json:"moduleSpecifier,omitempty"`
}

type Head struct {
	Node
	RawText       string `json:"rawText"`
	TemplateFlags int    `json:"templateFlags"`
	Text          string `json:"text"`
}

type HeritageClause struct {
	Node
	Token int     `json:"token"`
	Types []*Type `json:"types,omitempty"`
}

type Identifiers struct {
	Node
}

type ImportClause struct {
	Node
	IsTypeOnly    bool           `json:"isTypeOnly"`
	Name          *Name          `json:"name,omitempty"`
	NamedBindings *NamedBindings `json:"namedBindings,omitempty"`
}

type Incrementor struct {
	Node
	Left          *Left          `json:"left,omitempty"`
	Operand       *Operand       `json:"operand,omitempty"`
	Operator      int            `json:"operator"`
	OperatorToken *OperatorToken `json:"operatorToken,omitempty"`
	Right         *Right         `json:"right,omitempty"`
}

type IndexType struct {
	Node
	Literal *Literal `json:"literal,omitempty"`
}

type Initializer struct {
	Node
	ArgumentExpression       *ArgumentExpression     `json:"argumentExpression,omitempty"`
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	ColonToken               *ColonToken             `json:"colonToken,omitempty"`
	Condition                *Condition              `json:"condition,omitempty"`
	Declarations             []*Declaration          `json:"declarations,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Head                     *Head                   `json:"head,omitempty"`
	Left                     *Left                   `json:"left,omitempty"`
	Modifiers                []*Modifier             `json:"modifiers,omitempty"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	Operand                  *Operand                `json:"operand,omitempty"`
	Operator                 int                     `json:"operator"`
	OperatorToken            *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                     `json:"originalKeywordKind"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	QuestionDotToken         *QuestionDotToken       `json:"questionDotToken,omitempty"`
	QuestionToken            *QuestionToken          `json:"questionToken,omitempty"`
	Right                    *Right                  `json:"right,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan         `json:"templateSpans,omitempty"`
	Text                     string                  `json:"text"`
	Type                     *Type                   `json:"type,omitempty"`
	TypeArguments            []*TypeArgument         `json:"typeArguments,omitempty"`
	WhenFalse                *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue                 *WhenTrue               `json:"whenTrue,omitempty"`
}

type JsDoc struct {
	Node
	Comment any    `json:"comment,omitempty"`
	Tags    []*Tag `json:"tags,omitempty"`
}

type Left struct {
	Node
	ArgumentExpression       *ArgumentExpression `json:"argumentExpression,omitempty"`
	Arguments                []*Argument         `json:"arguments,omitempty"`
	EscapedText              string              `json:"escapedText"`
	Expression               *Expression         `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                `json:"hasExtendedUnicodeEscape"`
	Left                     *Left               `json:"left,omitempty"`
	Name                     *Name               `json:"name,omitempty"`
	NumericLiteralFlags      int                 `json:"numericLiteralFlags"`
	Operand                  *Operand            `json:"operand,omitempty"`
	Operator                 int                 `json:"operator"`
	OperatorToken            *OperatorToken      `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                 `json:"originalKeywordKind"`
	QuestionDotToken         *QuestionDotToken   `json:"questionDotToken,omitempty"`
	Right                    *Right              `json:"right,omitempty"`
	Text                     string              `json:"text"`
}

type LibReferenceDirective struct {
	Node
}

type Literal struct {
	Node
	HasExtendedUnicodeEscape bool     `json:"hasExtendedUnicodeEscape"`
	NumericLiteralFlags      int      `json:"numericLiteralFlags"`
	Operand                  *Operand `json:"operand,omitempty"`
	Operator                 int      `json:"operator"`
	RawText                  string   `json:"rawText"`
	TemplateFlags            int      `json:"templateFlags"`
	Text                     string   `json:"text"`
}

type Member struct {
	Node
	Body             *Body             `json:"body,omitempty"`
	ExclamationToken *ExclamationToken `json:"exclamationToken,omitempty"`
	Initializer      *Initializer      `json:"initializer,omitempty"`
	JsDoc            []*JsDoc          `json:"jsDoc,omitempty"`
	Modifiers        []*Modifier       `json:"modifiers,omitempty"`
	Name             *Name             `json:"name,omitempty"`
	Parameters       []*Parameter      `json:"parameters,omitempty"`
	QuestionToken    *QuestionToken    `json:"questionToken,omitempty"`
	Type             *Type             `json:"type,omitempty"`
}

type Modifier struct {
	Node
}

type ModuleSpecifier struct {
	Node
	HasExtendedUnicodeEscape bool   `json:"hasExtendedUnicodeEscape"`
	Text                     string `json:"text"`
}

type Name struct {
	Node
	Elements                 []*Element  `json:"elements,omitempty"`
	EscapedText              string      `json:"escapedText"`
	Expression               *Expression `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool        `json:"hasExtendedUnicodeEscape"`
	OriginalKeywordKind      int         `json:"originalKeywordKind"`
	Text                     string      `json:"text"`
}

type NamedBindings struct {
	Node
	Elements []*Element `json:"elements,omitempty"`
	Name     *Name      `json:"name,omitempty"`
}

type Node struct {
	End                int       `json:"end"`
	Flags              int       `json:"flags"`
	Kind               kind.Kind `json:"kind"`
	ModifierFlagsCache int       `json:"modifierFlagsCache"`
	Pos                int       `json:"pos"`
	TransformFlags     int       `json:"transformFlags"`
}

type ObjectType struct {
	Node
	TypeName *TypeName `json:"typeName,omitempty"`
}

type OpeningElement struct {
	Node
	Attributes *Attributes `json:"attributes,omitempty"`
	TagName    *TagName    `json:"tagName,omitempty"`
}

type OpeningFragment struct {
	Node
}

type Operand struct {
	Node
	ArgumentExpression  *ArgumentExpression `json:"argumentExpression,omitempty"`
	Arguments           []*Argument         `json:"arguments,omitempty"`
	EscapedText         string              `json:"escapedText"`
	Expression          *Expression         `json:"expression,omitempty"`
	Name                *Name               `json:"name,omitempty"`
	NumericLiteralFlags int                 `json:"numericLiteralFlags"`
	Operand             *Operand            `json:"operand,omitempty"`
	Operator            int                 `json:"operator"`
	Text                string              `json:"text"`
}

type OperatorToken struct {
	Node
}

type Parameter struct {
	Node
	DotDotDotToken *DotDotDotToken `json:"dotDotDotToken,omitempty"`
	Initializer    *Initializer    `json:"initializer,omitempty"`
	Name           *Name           `json:"name,omitempty"`
	QuestionToken  *QuestionToken  `json:"questionToken,omitempty"`
	Type           *Type           `json:"type,omitempty"`
}

type ParameterName struct {
	Node
	EscapedText string `json:"escapedText"`
}

type Pragmas struct {
	Node
}

type Property struct {
	Node
	Body        *Body        `json:"body,omitempty"`
	Expression  *Expression  `json:"expression,omitempty"`
	Initializer *Initializer `json:"initializer,omitempty"`
	Name        *Name        `json:"name,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
}

type PropertyName struct {
	Node
	EscapedText              string `json:"escapedText"`
	HasExtendedUnicodeEscape bool   `json:"hasExtendedUnicodeEscape"`
	Text                     string `json:"text"`
}

type QuestionDotToken struct {
	Node
}

type QuestionToken struct {
	Node
}

type ReferencedFile struct {
	Node
}

type Right struct {
	Node
	ArgumentExpression       *ArgumentExpression     `json:"argumentExpression,omitempty"`
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	ColonToken               *ColonToken             `json:"colonToken,omitempty"`
	Condition                *Condition              `json:"condition,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Head                     *Head                   `json:"head,omitempty"`
	Left                     *Left                   `json:"left,omitempty"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	Operand                  *Operand                `json:"operand,omitempty"`
	Operator                 int                     `json:"operator"`
	OperatorToken            *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                     `json:"originalKeywordKind"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	QuestionDotToken         *QuestionDotToken       `json:"questionDotToken,omitempty"`
	QuestionToken            *QuestionToken          `json:"questionToken,omitempty"`
	Right                    *Right                  `json:"right,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan         `json:"templateSpans,omitempty"`
	Text                     string                  `json:"text"`
	Type                     *Type                   `json:"type,omitempty"`
	TypeArguments            []*TypeArgument         `json:"typeArguments,omitempty"`
	WhenFalse                *WhenFalse              `json:"whenFalse,omitempty"`
	WhenTrue                 *WhenTrue               `json:"whenTrue,omitempty"`
}

type SourceFile struct {
	Node
	AmdDependencies         []*AmdDependency          `json:"amdDependencies,omitempty"`
	BindDiagnostics         []*BindDiagnostic         `json:"bindDiagnostics,omitempty"`
	EndOfFileToken          *EndOfFileToken           `json:"endOfFileToken,omitempty"`
	ExternalModuleIndicator *ExternalModuleIndicator  `json:"externalModuleIndicator,omitempty"`
	FileName                string                    `json:"fileName"`
	HasNoDefaultLib         bool                      `json:"hasNoDefaultLib"`
	IdentifierCount         int                       `json:"identifierCount"`
	Identifiers             *Identifiers              `json:"identifiers,omitempty"`
	IsDeclarationFile       bool                      `json:"isDeclarationFile"`
	LanguageVariant         int                       `json:"languageVariant"`
	LanguageVersion         int                       `json:"languageVersion"`
	LibReferenceDirectives  []*LibReferenceDirective  `json:"libReferenceDirectives,omitempty"`
	NodeCount               int                       `json:"nodeCount"`
	ParseDiagnostics        any                       `json:"parseDiagnostics,omitempty"`
	Pragmas                 *Pragmas                  `json:"pragmas,omitempty"`
	ReferencedFiles         []*ReferencedFile         `json:"referencedFiles,omitempty"`
	ScriptKind              int                       `json:"scriptKind"`
	Statements              []*Statement              `json:"statements,omitempty"`
	Text                    string                    `json:"text"`
	TypeReferenceDirectives []*TypeReferenceDirective `json:"typeReferenceDirectives,omitempty"`
}

type Statement struct {
	Node
	Body            *Body             `json:"body,omitempty"`
	CaseBlock       *CaseBlock        `json:"caseBlock,omitempty"`
	CatchClause     *CatchClause      `json:"catchClause,omitempty"`
	Condition       *Condition        `json:"condition,omitempty"`
	DeclarationList *DeclarationList  `json:"declarationList,omitempty"`
	ElseStatement   *ElseStatement    `json:"elseStatement,omitempty"`
	Expression      *Expression       `json:"expression,omitempty"`
	HeritageClauses []*HeritageClause `json:"heritageClauses,omitempty"`
	ImportClause    *ImportClause     `json:"importClause,omitempty"`
	Incrementor     *Incrementor      `json:"incrementor,omitempty"`
	Initializer     *Initializer      `json:"initializer,omitempty"`
	Members         []*Member         `json:"members,omitempty"`
	Modifiers       []*Modifier       `json:"modifiers,omitempty"`
	ModuleSpecifier *ModuleSpecifier  `json:"moduleSpecifier,omitempty"`
	MultiLine       bool              `json:"multiLine"`
	Name            *Name             `json:"name,omitempty"`
	Parameters      []*Parameter      `json:"parameters,omitempty"`
	Statement       *Statement        `json:"statement,omitempty"`
	Statements      []*Statement      `json:"statements,omitempty"`
	ThenStatement   *ThenStatement    `json:"thenStatement,omitempty"`
	TryBlock        *TryBlock         `json:"tryBlock,omitempty"`
	Type            *Type             `json:"type,omitempty"`
}

type Tag struct {
	Node
	Comment string   `json:"comment"`
	TagName *TagName `json:"tagName,omitempty"`
}

type TagName struct {
	Node
	EscapedText string      `json:"escapedText"`
	Expression  *Expression `json:"expression,omitempty"`
	Name        *Name       `json:"name,omitempty"`
}

type TemplateSpan struct {
	Node
	Expression *Expression `json:"expression,omitempty"`
	Literal    *Literal    `json:"literal,omitempty"`
}

type ThenStatement struct {
	Node
	MultiLine  bool         `json:"multiLine"`
	Statements []*Statement `json:"statements,omitempty"`
}

type TryBlock struct {
	Node
	MultiLine  bool         `json:"multiLine"`
	Statements []*Statement `json:"statements,omitempty"`
}

type Type struct {
	Node
	ElementType   *ElementType    `json:"elementType,omitempty"`
	ExprName      *ExprName       `json:"exprName,omitempty"`
	Expression    *Expression     `json:"expression,omitempty"`
	Literal       *Literal        `json:"literal,omitempty"`
	Members       []*Member       `json:"members,omitempty"`
	Operator      int             `json:"operator"`
	ParameterName *ParameterName  `json:"parameterName,omitempty"`
	Parameters    []*Parameter    `json:"parameters,omitempty"`
	Type          *Type           `json:"type,omitempty"`
	TypeArguments []*TypeArgument `json:"typeArguments,omitempty"`
	TypeName      *TypeName       `json:"typeName,omitempty"`
	TypeParameter *TypeParameter  `json:"typeParameter,omitempty"`
	Types         []*Type         `json:"types,omitempty"`
}

type TypeArgument struct {
	Node
	ExprName      *ExprName       `json:"exprName,omitempty"`
	Literal       *Literal        `json:"literal,omitempty"`
	Members       []*Member       `json:"members,omitempty"`
	TypeArguments []*TypeArgument `json:"typeArguments,omitempty"`
	TypeName      *TypeName       `json:"typeName,omitempty"`
	Types         []*Type         `json:"types,omitempty"`
}

type TypeName struct {
	Node
	EscapedText string `json:"escapedText"`
	Left        *Left  `json:"left,omitempty"`
	Right       *Right `json:"right,omitempty"`
}

type TypeParameter struct {
	Node
	Constraint *Constraint `json:"constraint,omitempty"`
	Name       *Name       `json:"name,omitempty"`
}

type TypeReferenceDirective struct {
	Node
}

type VariableDeclaration struct {
	Node
	Name *Name `json:"name,omitempty"`
}

type WhenFalse struct {
	Node
	ArgumentExpression       *ArgumentExpression     `json:"argumentExpression,omitempty"`
	Arguments                []*Argument             `json:"arguments,omitempty"`
	Attributes               *Attributes             `json:"attributes,omitempty"`
	Body                     *Body                   `json:"body,omitempty"`
	Children                 []*Children             `json:"children,omitempty"`
	ClosingElement           *ClosingElement         `json:"closingElement,omitempty"`
	Elements                 []*Element              `json:"elements,omitempty"`
	EqualsGreaterThanToken   *EqualsGreaterThanToken `json:"equalsGreaterThanToken,omitempty"`
	EscapedText              string                  `json:"escapedText"`
	Expression               *Expression             `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                    `json:"hasExtendedUnicodeEscape"`
	Head                     *Head                   `json:"head,omitempty"`
	Left                     *Left                   `json:"left,omitempty"`
	MultiLine                bool                    `json:"multiLine"`
	Name                     *Name                   `json:"name,omitempty"`
	NumericLiteralFlags      int                     `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement         `json:"openingElement,omitempty"`
	Operand                  *Operand                `json:"operand,omitempty"`
	Operator                 int                     `json:"operator"`
	OperatorToken            *OperatorToken          `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                     `json:"originalKeywordKind"`
	Parameters               []*Parameter            `json:"parameters,omitempty"`
	Properties               []*Property             `json:"properties,omitempty"`
	Right                    *Right                  `json:"right,omitempty"`
	TagName                  *TagName                `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan         `json:"templateSpans,omitempty"`
	Text                     string                  `json:"text"`
	TypeArguments            []*TypeArgument         `json:"typeArguments,omitempty"`
}

type WhenTrue struct {
	Node
	ArgumentExpression       *ArgumentExpression `json:"argumentExpression,omitempty"`
	Arguments                []*Argument         `json:"arguments,omitempty"`
	Attributes               *Attributes         `json:"attributes,omitempty"`
	Children                 []*Children         `json:"children,omitempty"`
	ClosingElement           *ClosingElement     `json:"closingElement,omitempty"`
	Elements                 []*Element          `json:"elements,omitempty"`
	EscapedText              string              `json:"escapedText"`
	Expression               *Expression         `json:"expression,omitempty"`
	HasExtendedUnicodeEscape bool                `json:"hasExtendedUnicodeEscape"`
	Head                     *Head               `json:"head,omitempty"`
	Left                     *Left               `json:"left,omitempty"`
	MultiLine                bool                `json:"multiLine"`
	Name                     *Name               `json:"name,omitempty"`
	NumericLiteralFlags      int                 `json:"numericLiteralFlags"`
	OpeningElement           *OpeningElement     `json:"openingElement,omitempty"`
	Operand                  *Operand            `json:"operand,omitempty"`
	Operator                 int                 `json:"operator"`
	OperatorToken            *OperatorToken      `json:"operatorToken,omitempty"`
	OriginalKeywordKind      int                 `json:"originalKeywordKind"`
	Properties               []*Property         `json:"properties,omitempty"`
	Right                    *Right              `json:"right,omitempty"`
	TagName                  *TagName            `json:"tagName,omitempty"`
	TemplateSpans            []*TemplateSpan     `json:"templateSpans,omitempty"`
	Text                     string              `json:"text"`
	TypeArguments            []*TypeArgument     `json:"typeArguments,omitempty"`
}
