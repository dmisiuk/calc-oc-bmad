# Frontend Architecture

## Component Architecture

**Component Organization:**
```
terminal/
├── components/
│   ├── calculator.go      # Main calculator grid component
│   ├── display.go         # Result display component
│   ├── button.go          # Calculator button component
│   ├── history.go         # History panel component
│   ├── keyboard.go        # Keyboard input handler
│   └── theme.go           # Theme management component
├── widgets/
│   ├── grid.go            # Calculator grid layout
│   ├── button_grid.go     # Number and operation buttons
│   ├── display_panel.go   # Main display area
│   └── history_panel.go   # History sidebar
├── events/
│   ├── mouse.go           # Mouse event handling
│   ├── keyboard.go        # Keyboard event handling
│   └── navigation.go      # Focus navigation
├── themes/
│   ├── default.go         # Default color theme
│   ├── dark.go            # Dark theme
│   └── light.go           # Light theme
└── utils/
    ├── focus.go           # Focus management
    ├── layout.go          # Layout utilities
    └── state.go           # UI state management
```

**Component Template:**
```go
package components

import (
    "github.com/calculator/internal/models"
    "github.com/calculator/internal/config"
    "github.com/rivo/tview"
)

type CalculatorComponent struct {
    app         *tview.Application
    grid        *tview.Grid
    display     *tview.TextView
    buttonGrid  *tview.Grid
    historyPanel *tview.List
    config      *config.Configuration
    theme       *Theme
    currentExpr string
    focusIndex  int
}

func NewCalculatorComponent(app *tview.Application, cfg *config.Configuration) *CalculatorComponent {
    return &CalculatorComponent{
        app:        app,
        config:     cfg,
        theme:      GetTheme(cfg.Theme),
        currentExpr: "",
        focusIndex: 0,
    }
}

func (cc *CalculatorComponent) Initialize() error {
    // Initialize TUI components and layout
    cc.createDisplay()
    cc.createButtonGrid()
    cc.createHistoryPanel()
    cc.setupLayout()
    cc.setupEventHandlers()
    return nil
}

func (cc *CalculatorComponent) UpdateDisplay(text string) {
    cc.display.SetText(text)
}

func (cc *CalculatorComponent) AddToExpression(char string) {
    cc.currentExpr += char
    cc.UpdateDisplay(cc.currentExpr)
}

func (cc *CalculatorComponent) ClearExpression() {
    cc.currentExpr = ""
    cc.UpdateDisplay("0")
}
```

## State Management Architecture

**State Structure:**
```go
type UIState struct {
    CurrentExpression string
    DisplayText      string
    LastResult       float64
    History          []models.Calculation
    Config           *config.Configuration
    IsRunning        bool
    FocusPosition    Position
    SelectedButton   string
    ShowHistory      bool
    ErrorMessage     string
}

type Position struct {
    Row int
    Col int
}

type ButtonState struct {
    Label    string
    Value    string
    Focused  bool
    Enabled  bool
    Position Position
}
```

**State Management Patterns:**
- TUI-managed state through tview components
- Reactive updates to UI components
- Event-driven state synchronization
- Focus management for keyboard navigation
- Component-local state for complex widgets

## Event Handling Architecture

**Event Organization:**
```
events/
├── mouse_events.go     # Mouse click and movement handling
├── keyboard_events.go  # Keyboard input and shortcut handling
├── focus_events.go     # Focus management and navigation
├── ui_events.go        # General UI event handling
└── shortcut_events.go  # Keyboard shortcut definitions
```

**Event Handling Pattern:**
```go
type EventHandler struct {
    calculator *components.CalculatorComponent
    parser     *parser.InputParser
    engine     *calculation.CalculationEngine
    history    *history.HistoryManager
}

func (eh *EventHandler) HandleButtonClick(button string) error {
    switch button {
    case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
        eh.calculator.AddToExpression(button)
    case "+", "-", "*", "/":
        eh.calculator.AddToExpression(" " + button + " ")
    case "=":
        return eh.processCalculation()
    case "C":
        eh.calculator.ClearExpression()
    case "H":
        eh.calculator.ToggleHistory()
    }
    return nil
}

func (eh *EventHandler) HandleKeyEvent(key *tview.EventKey) *tview.EventKey {
    switch key.Key {
    case tview.KeyEnter:
        eh.processCalculation()
    case tview.KeyEscape:
        eh.calculator.ClearExpression()
    case tview.KeyCtrlH:
        eh.calculator.ToggleHistory()
    default:
        if key.Rune != 0 {
            eh.calculator.AddToExpression(string(key.Rune))
        }
    }
    return key
}
```

## Frontend Services Layer

**TUI Service Setup:**
```go
package services

type TUIService struct {
    app        *tview.Application
    calculator *components.CalculatorComponent
    events     *events.EventHandler
    config     *config.Configuration
}

func NewTUIService(cfg *config.Configuration) *TUIService {
    app := tview.NewApplication()
    calculator := components.NewCalculatorComponent(app, cfg)
    events := events.NewEventHandler(calculator)
    
    return &TUIService{
        app:        app,
        calculator: calculator,
        events:     events,
        config:     cfg,
    }
}

func (ts *TUIService) Initialize() error {
    if err := ts.calculator.Initialize(); err != nil {
        return err
    }
    
    ts.app.SetRoot(ts.calculator.GetGrid(), true)
    return nil
}

func (ts *TUIService) Run() error {
    return ts.app.Run()
}
```

**Service Example:**
```go
type FocusService struct {
    calculator *components.CalculatorComponent
    focusMap   map[string]*tview.Primitive
    currentPos int
}

func (fs *FocusService) MoveFocus(direction string) {
    switch direction {
    case "up":
        fs.currentPos = max(0, fs.currentPos-1)
    case "down":
        fs.currentPos = min(len(fs.focusMap)-1, fs.currentPos+1)
    case "left":
        fs.currentPos = max(0, fs.currentPos-1)
    case "right":
        fs.currentPos = min(len(fs.focusMap)-1, fs.currentPos+1)
    }
    
    // Update focus visually
    fs.updateFocusVisual()
}
```
