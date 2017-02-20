package electron

import "github.com/gopherjs/gopherjs/js"

// Task a Structure
type Task struct {
	*js.Object
	// Path of the program to execute, usually you should specify which opens the current program.
	Program string `js:"program"`
	// The command line arguments when is executed.
	Arguments string `js:"arguments"`
	// The string to be displayed in a JumpList.
	Title string `js:"title"`
	// Description of this task.
	Description string `js:"description"`
	// The absolute path to an icon to be displayed in a JumpList, which can be an arbitrary resource file that contains an icon. You can usually specify to show the icon of the program.
	IconPath string `js:"iconPath"`
	// The icon index in the icon file. If an icon file consists of two or more icons, set this value to identify the icon. If an icon file consists of one icon, this value is 0.
	IconIndex float64 `js:"iconIndex"`
}
