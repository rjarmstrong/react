// Code generated by reactGen. DO NOT EDIT.

package react

// H4Props defines the properties for the <h4> element
type H4Props struct {
	AriaSet
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (h *H4Props) assign(v *_H4Props) {

	if h.AriaSet != nil {
		for dk, dv := range h.AriaSet {
			v.o.Set("aria-"+dk, dv)
		}
	}

	v.ClassName = h.ClassName

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

	if h.DataSet != nil {
		for dk, dv := range h.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	if h.OnChange != nil {
		v.o.Set("onChange", h.OnChange.OnChange)
	}

	if h.OnClick != nil {
		v.o.Set("onClick", h.OnClick.OnClick)
	}

	if h.Ref != nil {
		v.o.Set("ref", h.Ref.Ref)
	}

	v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = h.Style.hack()

}
