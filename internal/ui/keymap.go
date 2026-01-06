package ui

import tea "github.com/charmbracelet/bubbletea"

const (
	KeyUp           = "up"
	KeyDown         = "down"
	KeyLeft         = "left"
	KeyRight        = "right"
	KeyJ            = "j"
	KeyK            = "k"
	KeyG            = "g"
	KeyShiftG       = "G"
	KeyEnter        = "enter"
	KeySpace        = " "
	KeySlash        = "/"
	KeyEsc          = "esc"
	KeyQ            = "q"
	KeyQuestion     = "?"
	KeyY            = "y"
	KeyC            = "c"
	KeyD            = "d"
	KeyX            = "x"
	KeyP            = "p"
	KeyShiftN       = "N"
	KeyShiftD       = "D"
	KeyShiftS       = "S"
	KeyN            = "n"
	KeyM            = "m"
	KeyS            = "s"
	KeyCtrlV        = "ctrl+v"
	KeyCtrlC        = "ctrl+c"
	KeyCtrlL        = "ctrl+l"
	KeyBackspace    = "backspace"
	KeyDelete       = "delete"
	KeyHome         = "home"
	KeyEnd          = "end"
	KeyTab          = "tab"
	KeyShiftTab     = "shift+tab"
	KeyU_TR         = "ü"
	KeyG_TR         = "ğ"
	KeyBracketRight = "]"
	KeyBracketLeft  = "["
	Key1            = "1"
	Key2            = "2"
	Key3            = "3"
	Key4            = "4"
	KeyPgUp         = "pgup"
	KeyPgDn         = "pgdown"
	KeyCtrlR        = "ctrl+r"
	KeyZ            = "z"
	KeyO            = "o"
	KeyShiftT       = "T"
	KeyShiftW       = "W"
	KeyU            = "u"
	KeyShiftU       = "U"
)

func IsKey(msg tea.KeyMsg, keys ...string) bool {
	for _, k := range keys {
		if msg.String() == k {
			return true
		}
	}
	return false
}

type HelpSection struct {
	Title string
	Items []HelpItem
}

type HelpItem struct {
	Key  string
	Desc string
}

func GetHelpSections() []HelpSection {
	return []HelpSection{
		{
			Title: "Navigation",
			Items: []HelpItem{
				{"j/k", "Move down/up"},
				{"g/G", "Jump to top/bottom"},
				{"Enter", "Toggle detail view"},
				{"z", "Maximize/minimize detail"},
			},
		},
		{
			Title: "Filter",
			Items: []HelpItem{
				{"/", "Start filter"},
				{"Tab", "Cycle level filter"},
				{"Ctrl+R", "Clear filter"},
				{"ESC", "Exit filter mode"},
			},
		},
		{
			Title: "Selection",
			Items: []HelpItem{
				{"s", "Toggle selection"},
				{"S", "Select all/clear"},
				{"c", "Copy selected"},
				{"d", "Delete selected"},
			},
		},
		{
			Title: "Actions",
			Items: []HelpItem{
				{"y", "Copy visible logs + notes"},
				{"c", "Copy current line"},
				{"d", "Delete current"},
				{"u/U", "Undo/redo delete"},
				{"x", "Clear all"},
				{"p", "Paste from clipboard"},
				{"o", "Open file"},
			},
		},
		{
			Title: "Notes",
			Items: []HelpItem{
				{"N", "Write/edit note"},
				{"D", "Delete note"},
				{"n", "Show/hide note"},
				{"m", "Show/hide all notes"},
				{"]/[", "Next/prev noted line"},
			},
		},
		{
			Title: "Workspace",
			Items: []HelpItem{
				{"T", "New workspace"},
				{"Tab", "Next workspace"},
				{"Shift+Tab", "Prev workspace"},
				{"W", "Close workspace"},
			},
		},
		{
			Title: "Signal",
			Items: []HelpItem{
				{"1", "Error frequency"},
				{"2", "First/last seen"},
				{"3", "Burst detector"},
				{"4", "Error diversity"},
			},
		},
		{
			Title: "Tools",
			Items: []HelpItem{
				{"Ctrl+L", "HTTP status lookup"},
				{"?", "Toggle help"},
				{"q", "Quit"},
			},
		},
	}
}

func FooterHints(mode int) string {
	const (
		modeList   = 0
		modeFilter = 1
		modeDetail = 2
		modeHelp   = 3
		modeNotes  = 4
		modeLookup = 5
		modeSignal = 6
	)

	switch mode {
	case modeList:
		return "j/k:nav  /:filter  Enter:detail  ?:help  q:quit"
	case modeFilter:
		return "Type to filter  ESC:cancel  Enter:apply"
	case modeDetail:
		return "ESC:back  j/k:nav  c:copy  ?:help"
	case modeHelp:
		return "ESC/?:close"
	case modeNotes:
		return "Type notes  ESC:close  Ctrl+C:copy notes"
	case modeLookup:
		return "Type code  j/k:nav  Enter:copy  ESC:close"
	case modeSignal:
		return "c:copy  ESC:close"
	default:
		return ""
	}
}
