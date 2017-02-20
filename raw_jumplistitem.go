package electron

import "github.com/gopherjs/gopherjs/js"

// JumpListItem a Structure
type JumpListItem struct {
	*js.Object
	// One of the following:
	Type JumpListItemJumpListItemType `js:"type"`
	// Path of the file to open, should only be set if is .
	Path string `js:"path"`
	// Path of the program to execute, usually you should specify which opens the current program. Should only be set if is .
	Program string `js:"program"`
	// The command line arguments when is executed. Should only be set if is .
	Args string `js:"args"`
	// The text to be displayed for the item in the Jump List. Should only be set if is .
	Title string `js:"title"`
	// Description of the task (displayed in a tooltip). Should only be set if is .
	Description string `js:"description"`
	// The absolute path to an icon to be displayed in a Jump List, which can be an arbitrary resource file that contains an icon (e.g. , , ). You can usually specify to show the program icon.
	IconPath string `js:"iconPath"`
	// The index of the icon in the resource file. If a resource file contains multiple icons this value can be used to specify the zero-based index of the icon that should be displayed for this task. If a resource file contains only one icon, this property should be set to zero.
	IconIndex float64 `js:"iconIndex"`
}

type JumpListItemJumpListItemType string

// consts
const (
	JumpListItemJumpListItemTypeTask      JumpListItemJumpListItemType = "task"
	JumpListItemJumpListItemTypeSeparator JumpListItemJumpListItemType = "separator"
	JumpListItemJumpListItemTypeFile      JumpListItemJumpListItemType = "file"
)
