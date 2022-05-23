package calendarmonth

// from file "react-core/src/components/CalendarMonth/CalendarMonth.tsx"

type CalendarFormat struct {
	// MonthFormat - How to format months in Select. Optional.
	MonthFormat func(date any // Date ) any // React.ReactNode 
	// WeekdayFormat - How to format week days in header. Optional.
	WeekdayFormat func(date any // Date ) any // React.ReactNode 
	// LongWeekdayFormat - How to format days in header for screen readers. Optional.
	LongWeekdayFormat func(date any // Date ) any // React.ReactNode 
	// DayFormat - How to format days in buttons in table cells. Optional.
	DayFormat func(date any // Date ) any // React.ReactNode 
	// Locale - If using the default formatters which locale to use. Undefined defaults to current locale. See
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl#Locale_identification_and_negotiation.
	// Optional.
	Locale string
	// WeekStart - Day of week that starts the week. 0 is Sunday, 6 is Saturday. Optional.
	WeekStart any /* "0" | "1" | "2" | "3" | "4" | "5" | "6" | any // Weekday  */
	// RangeStart - Which date to start range styles from. Optional.
	RangeStart any // Date 
	// PrevMonthAriaLabel - Aria-label for the previous month button. Optional.
	PrevMonthAriaLabel string
	// NextMonthAriaLabel - Aria-label for the next month button. Optional.
	NextMonthAriaLabel string
	// YearInputAriaLabel - Aria-label for the year input. Optional.
	YearInputAriaLabel string
	// CellAriaLabel - Aria-label for the date cells. Optional.
	CellAriaLabel func(date any // Date ) string
}

type CalendarProps struct {
	// Date - Month/year to base other dates around. Optional.
	Date any // Date 
	// OnChange - Callback when date is selected. Optional.
	OnChange func(date any // Date )
	// Validators - Functions that returns if a date is valid and selectable. Optional.
	Validators []( func(date any // Date ) bool )
	// ClassName - Classname to add to outer div. Optional.
	ClassName string
	// OnSelectToggle - Optional.
	OnSelectToggle func(open bool)
}
