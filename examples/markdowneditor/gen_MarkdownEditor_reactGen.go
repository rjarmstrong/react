// Code generated by reactGen. DO NOT EDIT.

package markdowneditor

import "myitcv.io/react"

type MarkdownEditorElem struct {
	react.Element
}

func buildMarkdownEditor(cd react.ComponentDef) react.Component {
	return MarkdownEditorDef{ComponentDef: cd}
}

func buildMarkdownEditorElem(children ...react.Element) *MarkdownEditorElem {
	return &MarkdownEditorElem{
		Element: react.CreateElement(buildMarkdownEditor, nil, children...),
	}
}

// SetState is an auto-generated proxy proxy to update the state for the
// MarkdownEditor component.  SetState does not immediately mutate m.State()
// but creates a pending state transition.
func (m MarkdownEditorDef) SetState(state MarkdownEditorState) {
	m.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the MarkdownEditor component
func (m MarkdownEditorDef) State() MarkdownEditorState {
	return m.ComponentDef.State().(MarkdownEditorState)
}

// IsState is an auto-generated definition so that MarkdownEditorState implements
// the myitcv.io/react.State interface.
func (m MarkdownEditorState) IsState() {}

var _ react.State = MarkdownEditorState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (m MarkdownEditorDef) GetInitialStateIntf() react.State {
	return m.GetInitialState()
}

func (m MarkdownEditorState) EqualsIntf(val react.State) bool {
	return m == val.(MarkdownEditorState)
}
